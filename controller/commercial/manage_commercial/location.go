package manage_commercial

import (
	"commercial-propfloor/database"
	"commercial-propfloor/models"
	"context"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func AddCountryDetails() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("ankit")
		country_name := c.PostForm("country_name")
		currency_name := c.PostForm("currency_name")
		currency_symbol := c.PostForm("currency_symbol")

		a := InsertCountryDetailsInDB(country_name, currency_name, currency_symbol)
		fmt.Println(a)
	}
}

func InsertCountryDetailsInDB(country_name string, currency_name string, currency_symbol string) (id int64) {
	godotenv.Load()
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	collection := client.Database("india").Collection("country")
	doc_c := models.Destination{CountryName: country_name, CurrencyName: currency_name, CurrencySymbol: currency_symbol}

	//pattern_c := regexp.MustCompile("^[A-Za-z]$")

	validate := validator.New()
	err := validate.Struct(doc_c)
	fmt.Println(err)
	var a = 2
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}

	} else if a == 2 {
		res, errr := collection.InsertOne(context.Background(), doc_c)
		if errr != nil {
			log.Fatal(errr)
		}
		idd_c := res.InsertedID
		fmt.Println("datatype", reflect.TypeOf(idd_c))
		fmt.Println("inserted-cityid", idd_c)
	} else {
		fmt.Println("Yadav")
	}
	fmt.Println(os.Getenv("APP_NAME"))
	return

}
