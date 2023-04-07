package routes

import (
	"commercial-propfloor/controller/commercial/manage_commercial"

	"github.com/gin-gonic/gin"
)

func PrivateRoutes(privateRoutes *gin.Engine) {
	privateRoutes.POST("/Manage/Commercialportal/SaveBasicBuildingDetails", manage_commercial.SaveBasicBuildingDetails())
	privateRoutes.POST("/Manage/Commercialportal/AddCountryDetails", manage_commercial.AddCountryDetails())
}

func PublicRoutes(publicRoutes *gin.Engine) {

}
