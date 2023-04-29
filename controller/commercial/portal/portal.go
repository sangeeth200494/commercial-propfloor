package portal

import (
	//"commercial-propfloor/controller/commercial/manage_commercial"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// type Prop struct {
// 	Title     string
// 	MainImage string
// 	Done      bool
// }

// type PropFloor struct {
// 	PageTitle string
// 	Props     []Prop
// }

// func Sangeeth() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		a := "sangeeth"
// 		fmt.Println(a)

// 		return
// 	}
// }

func Cities() gin.HandlerFunc {
	return func(c *gin.Context) {
		godotenv.Load()
		//manage_commercial.SelectCityindatabase()
		Response, _ := http.Get(os.Getenv("PORT_URL"))

		ResponseData, _ := ioutil.ReadAll(Response.Body)

		fmt.Fprintf(os.Stdout, "%s", ResponseData)

	}
}
