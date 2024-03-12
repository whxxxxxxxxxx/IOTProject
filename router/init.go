package router

import "github.com/gin-gonic/gin"

func InitRouter() {
	// This is where you would define your routes
	r := gin.Default()

	err := r.Run("3000")
	if err != nil {
		panic(err)
	}

}
