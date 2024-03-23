package appInitialize

import "IOTProject/internal/app/device"

func init() {
	apps = append(apps, &device.Device{Name: "Device module"})
}
