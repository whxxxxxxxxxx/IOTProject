package dto

import "IOTProject/internal/app/device/model"

type DeviceItem struct {
	DeviceId     string `json:"device_id"`
	Name         string `json:"name"`
	Model        string `json:"model"`
	SerialNumber string `json:"serial_number"`
	Location     struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
	NetworkInfo struct {
		IpAddress  string `json:"ip_address"`
		MacAddress string `json:"mac_address"`
	} `json:"network_info"`
	Security struct {
		EncryptionStatus     string `json:"encryption_status"`
		AuthenticationMethod string `json:"authentication_method"`
	} `json:"security"`
}

type DeviceItemPage struct {
	Total int            `json:"total"`
	List  []model.Device `json:"list"`
}

type DeviceItemPageReq struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type DeviceItemSearchReq struct {
	PageReq      DeviceItemPageReq `json:"page_req"`
	DeviceNo     string            `json:"device_no"`
	Name         string            `json:"name"`
	Model        string            `json:"model"`
	SerialNumber string            `json:"serial_number"`
}
