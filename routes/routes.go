package routes

import (
	"commercial-propfloor/controller/commercial/manage_commercial"

	"github.com/gin-gonic/gin"
)

func PrivateRoutes(privateRoutes *gin.Engine) {
	privateRoutes.POST("/Manage/Commercialportal/SaveBasicBuildingDetails", manage_commercial.SaveBasicBuildingDetails())
	privateRoutes.POST("/Manage/Commercialportal/AddCountryDetails", manage_commercial.AddCountryDetails())
	privateRoutes.POST("/Manage/Commercialportal/AddStateDetails", manage_commercial.AddStateDetails())
	privateRoutes.POST("/Manage/Commercialportal/AddCityDetails", manage_commercial.AddCityDetails())
	privateRoutes.POST("/Manage/Commercialportal/AddLocalityDetails", manage_commercial.AddLocalityDetails())
}

func PublicRoutes(publicRoutes *gin.Engine) {

}
