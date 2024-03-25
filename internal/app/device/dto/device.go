package dto

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
