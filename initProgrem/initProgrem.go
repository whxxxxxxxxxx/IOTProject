package initProgrem

import (
	"IOTProject/model"
	"IOTProject/router"
)

func Init() {
	model.InitModel()
	router.InitRouter()

}
