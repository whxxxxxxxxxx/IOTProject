package appInitialize

import "IOTProject/internal/app/data"

func init() {
	apps = append(apps, &data.Data{Name: "Data module"})
}
