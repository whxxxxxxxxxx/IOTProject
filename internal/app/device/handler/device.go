package handler

import (
	"IOTProject/internal/app/device/dao"
	"IOTProject/internal/app/device/dto"
	"IOTProject/internal/app/device/model"
	"IOTProject/internal/app/device/service"
	model2 "IOTProject/internal/model"
	"IOTProject/middleware/response"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"regexp"
	"strconv"
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

	service.UpdateDevicesList()
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

	service.UpdateDevicesList()
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
	//获取query参数
	pageNo := c.Query("page")
	//转化为数字类型
	pageNoInt, _ := strconv.Atoi(pageNo)
	pageSize := c.Query("pageSize")
	pageSizeInt, _ := strconv.Atoi(pageSize)
	if pageNo == "" || pageSize == "" {
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
	if pageNo != "" && pageSize != "" {
		var totalDevices dto.DeviceItemPage
		var devices []model.Device
		err := dao.Device.Model(&model.Device{}).
			Preload("Location").
			Preload("NetworkInfo").
			Preload("Security").
			WithContext(c.Request.Context()).
			Offset((pageNoInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&devices).Error
		if err != nil {
			response.ServiceErr(c, err)
			return
		}
		totalDevices.Total = len(devices)
		totalDevices.List = devices
		response.HTTPSuccess(c, totalDevices)

	}

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

func ListDevicesPage2(c *gin.Context) {
	var req dto.DeviceItemPageReq
	var totalDevices dto.DeviceItemPage2
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ServiceErr(c, err)
		return
	}
	if req.PageSize == 0 {
		response.ServiceErr(c, errors.New("PageSize can not be 0"))
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
	var count int64
	err = dao.Device.Model(&model.Device{}).WithContext(c.Request.Context()).Count(&count).Error
	if err != nil {
		response.ServiceErr(c, err)
		return
	}
	pageNum := int(count) / req.PageSize
	if int(count)%req.PageSize != 0 {
		pageNum++
	}
	totalDevices.TotalPage = pageNum
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

func SearchDevices2(c *gin.Context) {
	var req dto.DeviceItemSearchReq
	var totalDevices dto.DeviceItemPage2
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ServiceErr(c, err)
		return
	}

	var count int64
	err := dao.Device.Model(&model.Device{}).
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
		}).WithContext(c.Request.Context()).Count(&count).Error
	if err != nil {
		response.ServiceErr(c, err)
		return
	}

	var devices []model.Device
	err = dao.Device.Model(&model.Device{}).
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
	pageNum := int(count) / req.PageReq.PageSize
	if int(count)%req.PageReq.PageSize != 0 {
		pageNum++
	}
	totalDevices.TotalPage = pageNum
	totalDevices.Total = len(devices)
	totalDevices.List = devices
	response.HTTPSuccess(c, totalDevices)
}

func StatusData(c *gin.Context) {
	var status []struct {
		EncryptionStatus string
		Count            int
	}
	err := dao.Device.Model(&model.Security{}).
		Select("encryption_status , count(*) as count").
		Group("encryption_status").
		Scan(&status).Error
	if err != nil {
		response.ServiceErr(c, err)
		return
	}
	response.HTTPSuccess(c, status)
}

func LocationData(c *gin.Context) {
	var location []model.Location
	err := dao.Device.Model(&model.Location{}).
		WithContext(c.Request.Context()).
		Find(&location).Error
	if err != nil {
		response.ServiceErr(c, err)
		return
	}
	response.HTTPSuccess(c, location)
}

func ModelData(c *gin.Context) {
	var models []struct {
		ModelData string
		Count     int
	}
	err := dao.Device.Model(&model.Device{}).
		Select("model_data , count(*) as count").
		Group("model_data").
		Scan(&models).Error
	if err != nil {
		response.ServiceErr(c, err)
		return
	}
	response.HTTPSuccess(c, models)
}

func NameData(c *gin.Context) {
	var name []struct {
		Name  string
		Count int
	}
	err := dao.Device.Model(&model.Device{}).
		Select("name , count(*) as count").
		Group("name").
		Scan(&name).Error
	if err != nil {
		response.ServiceErr(c, err)
		return
	}
	re := regexp.MustCompile("[\u4e00-\u9fa5]+")
	merged := make(map[string]int)
	for _, nc := range name {
		// 提取中文部分
		chineseOnly := re.FindString(nc.Name)
		if chineseOnly != "" {
			// 如果已存在，累加数量
			merged[chineseOnly] += nc.Count
		}
	}
	response.HTTPSuccess(c, merged)
}
