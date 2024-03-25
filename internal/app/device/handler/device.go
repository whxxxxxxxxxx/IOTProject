package handler

import (
	"IOTProject/internal/app/device/dao"
	"IOTProject/internal/app/device/dto"
	"IOTProject/internal/app/device/model"
	model2 "IOTProject/internal/model"
	"IOTProject/middleware/response"
	"github.com/gin-gonic/gin"
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
	deviceItem = model.Device{
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

	err = dao.Device.Model(&model.Device{}).
		WithContext(c.Request.Context()).
		Save(&deviceItem).Error
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
