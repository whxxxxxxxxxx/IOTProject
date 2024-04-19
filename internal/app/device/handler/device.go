package handler

import (
	"IOTProject/internal/app/device/dao"
	"IOTProject/internal/app/device/dto"
	"IOTProject/internal/app/device/model"
	model2 "IOTProject/internal/model"
	"IOTProject/middleware/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateDevice(c *gin.Context) {
	var req dto.DeviceItem
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ServiceErr(c, err)
		return
	}
	deviceItem := model.Device{
		DeviceNo:     req.DeviceId,
		Name:         req.Name,
		ModelData:    req.Model,
		SerialNumber: req.SerialNumber,
		Location: model.Location{
			Latitude:  req.Location.Latitude,
			Longitude: req.Location.Longitude,
		},
		NetworkInfo: model.NetworkInfo{
			IPAddress:  req.NetworkInfo.IpAddress,
			MacAddress: req.NetworkInfo.MacAddress,
		},
		Security: model.Security{
			EncryptionStatus:     req.Security.EncryptionStatus,
			AuthenticationMethod: req.Security.AuthenticationMethod,
		},
	}

	err := dao.Device.Model(&model.Device{}).WithContext(c.Request.Context()).Create(&deviceItem).Error
	if err != nil {
		response.ServiceErr(c, err)
		return
	}

	response.HTTPSuccess(c, nil)

}

func DeleteDevice(c *gin.Context) {
	id := c.Param("id")

	err := dao.Device.Model(&model.Device{}).
		Select(clause.Associations).
		WithContext(c.Request.Context()).
		Delete(&model.Device{
			Base: model2.Base{
				ID: id,
			}}).Error
	if err != nil {
		response.ServiceErr(c, err)
		return
	}
	response.HTTPSuccess(c, nil)
}

func UpdateDevice(c *gin.Context) {
	id := c.Param("id")
	var req dto.DeviceItem
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ServiceErr(c, err)
		return
	}
	var deviceItem model.Device
	err := dao.Device.Model(&model.Device{}).
		WithContext(c.Request.Context()).
		Where("id = ?", id).First(&deviceItem).Error
	deviceItem.DeviceNo = req.DeviceId
	deviceItem.Name = req.Name
	deviceItem.ModelData = req.Model
	deviceItem.SerialNumber = req.SerialNumber
	deviceItem.Location = model.Location{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}
	deviceItem.NetworkInfo = model.NetworkInfo{
		IPAddress:  req.NetworkInfo.IpAddress,
		MacAddress: req.NetworkInfo.MacAddress,
	}
	deviceItem.Security = model.Security{
		EncryptionStatus:     req.Security.EncryptionStatus,
		AuthenticationMethod: req.Security.AuthenticationMethod,
	}

	err = dao.Device.Model(&deviceItem).
		Session(&gorm.Session{FullSaveAssociations: true}).
		WithContext(c.Request.Context()).
		Where("id = ?", id).
		Updates(&deviceItem).Error
	if err != nil {
		response.ServiceErr(c, err)
		return
	}

	response.HTTPSuccess(c, nil)
}

func GetDevice(c *gin.Context) {
	id := c.Param("id")
	var device model.Device
	err := dao.Device.Model(&model.Device{}).
		Preload("Location").
		Preload("NetworkInfo").
		Preload("Security").
		WithContext(c.Request.Context()).
		Where("id = ?", id).First(&device).Error
	if err != nil {
		response.ServiceErr(c, err)
		return

	}
	response.HTTPSuccess(c, device)
}

func ListDevices(c *gin.Context) {
	var devices []model.Device
	err := dao.Device.Model(&model.Device{}).
		Preload("Location").
		Preload("NetworkInfo").
		Preload("Security").
		WithContext(c.Request.Context()).
		Find(&devices).Error
	if err != nil {
		response.ServiceErr(c, err)
		return
	}
	response.HTTPSuccess(c, devices)
}

func ListDevicesPage(c *gin.Context) {
	var req dto.DeviceItemPageReq
	var totalDevices dto.DeviceItemPage
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ServiceErr(c, err)
		return
	}
	var devices []model.Device
	err := dao.Device.Model(&model.Device{}).
		Preload("Location").
		Preload("NetworkInfo").
		Preload("Security").
		WithContext(c.Request.Context()).
		Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Find(&devices).Error
	if err != nil {
		response.ServiceErr(c, err)
		return
	}
	totalDevices.Total = len(devices)
	totalDevices.List = devices
	response.HTTPSuccess(c, totalDevices)
}

func SearchDevices(c *gin.Context) {
	var req dto.DeviceItemSearchReq
	var totalDevices dto.DeviceItemPage
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ServiceErr(c, err)
		return
	}
	var devices []model.Device
	err := dao.Device.Model(&model.Device{}).
		Preload("Location").
		Preload("NetworkInfo").
		Preload("Security").
		Scopes(func(db *gorm.DB) *gorm.DB {
			if req.DeviceNo != "" {
				db = db.Where("device_no LIKE ?", "%"+req.DeviceNo+"%")
			}
			if req.Name != "" {
				db = db.Where("name LIKE ?", "%"+req.Name+"%")
			}
			if req.Model != "" {
				db = db.Where("model LIKE ?", "%"+req.Model+"%")
			}
			if req.SerialNumber != "" {
				db = db.Where("serial_number LIKE ?", "%"+req.SerialNumber+"%")
			}
			return db
		}).
		WithContext(c.Request.Context()).
		Offset((req.PageReq.Page - 1) * req.PageReq.PageSize).Limit(req.PageReq.PageSize).Find(&devices).Error
	if err != nil {
		response.ServiceErr(c, err)
		return
	}
	totalDevices.Total = len(devices)
	totalDevices.List = devices
	response.HTTPSuccess(c, totalDevices)
}
