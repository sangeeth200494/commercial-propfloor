package manage_commercial

import (
	"commercial-propfloor/database"
	"commercial-propfloor/models"
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func AddCountryDetails() gin.HandlerFunc {
	return func(c *gin.Context) {

		//fmt.Println("ankit")
		country_name := c.PostForm("country_name")
		currency_name := c.PostForm("currency_name")
		currency_symbol := c.PostForm("currency_symbol")

		/*a := */
		InsertCountryDetailsInDB(country_name, currency_name, currency_symbol)
		//fmt.Println(a)
	}
}

func InsertCountryDetailsInDB(country_name string, currency_name string, currency_symbol string) (id int64) {
	godotenv.Load()
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("country")
	doc_c := models.Destination{CountryName: country_name, CurrencyName: currency_name, CurrencySymbol: currency_symbol}

	pattern_c := regexp.MustCompile("^[a-zA-Z]*$")

	validate := validator.New()
	err := validate.Struct(doc_c)
	//fmt.Println(err)
	// 	var a = 2
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}
	} else if pattern_c.MatchString(doc_c.CountryName) && pattern_c.MatchString(doc_c.CurrencyName) && pattern_c.MatchString(doc_c.CurrencySymbol) {
		CountryCount := CheckCountryExist(country_name)
		fmt.Println(CountryCount)
		if CountryCount == 0 {
			res, errr := collection.InsertOne(context.Background(), doc_c)
			//fmt.Println(res)
			if errr != nil {
				log.Fatal(errr)
			}
			idd_c := res.InsertedID
			fmt.Println("datatype", reflect.TypeOf(idd_c))
			fmt.Println("inserted-Destinationid  ", idd_c)
		} else {
			fmt.Println("Country Already Exist")
		}
	} else {
		fmt.Println("invalid input")
	}

	fmt.Println(os.Getenv("APP_NAME"))
	return

}

func CheckCountryExist(country_name_a string) (counts int64) {
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("country")
	//filter := bson.D{{"country_name" ,country_name_a}}

	res, err := collection.CountDocuments(context.Background(), bson.M{"country_name": country_name_a})
	if err != nil {
		fmt.Println(err)
	}

	return res
}

// creating state function
func AddStateDetails() gin.HandlerFunc {
	return func(c *gin.Context) {

		state_name := c.PostForm("state_name")
		state_language := c.PostForm("state_language")

		InsertStateDetailsInDB(state_name, state_language)
		//fmt.Println(a)
	}
}
func InsertStateDetailsInDB(state_name string, state_language string) (id int64) {
	godotenv.Load()
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("state")
	doc := models.State{StateName: state_name, StateLanguage: state_language}

	pattern := regexp.MustCompile("^[a-zA-Z]*$")

	validate := validator.New()
	err := validate.Struct(doc)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}

	} else if pattern.MatchString(doc.StateName) && pattern.MatchString(doc.StateLanguage) {
		StateCount := CheckStateExist(state_name)
		if StateCount == 0 {
			res, errr := collection.InsertOne(context.Background(), doc)
			if errr != nil {
				log.Fatal(errr)
			}
			idd := res.InsertedID
			fmt.Println("datatype", reflect.TypeOf(idd))
			fmt.Println("Inserted-State", idd)
		} else {
			fmt.Println("State Already Exist")
		}
	} else {
		fmt.Println("Invalid Input")
	}

	fmt.Println(os.Getenv("APP_NAME"))
	return

}

func CheckStateExist(state_name string) (counts int64) {
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("state")

	res, err := collection.CountDocuments(context.Background(), bson.M{"state_name": state_name})
	if err != nil {
		fmt.Println(err)
	}
	return res
}

// City function creation
func AddCityDetails() gin.HandlerFunc {
	return func(c *gin.Context) {

		//fmt.Println("ankit")
		city_name := c.PostForm("city_name")

		/*a := */
		InsertCityDetailsInDB(city_name)
		//fmt.Println(a)
	}
}

func InsertCityDetailsInDB(city_name string) (id int64) {
	godotenv.Load()
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("city")
	doc := models.City{CityName: city_name}

	pattern := regexp.MustCompile("^[a-zA-Z]*$")

	validate := validator.New()
	err := validate.Struct(doc)
	//fmt.Println(err)
	// 	var a = 2
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}
	} else if pattern.MatchString(doc.CityName) {
		CityCount := CheckCityExist(city_name)
		fmt.Println(CityCount)
		if CityCount == 0 {
			res, errr := collection.InsertOne(context.Background(), doc)
			//fmt.Println(res)
			if errr != nil {
				log.Fatal(errr)
			}
			idd_c := res.InsertedID
			fmt.Println("datatype", reflect.TypeOf(idd_c))
			fmt.Println("inserted-Destinationid  ", idd_c)
		} else {
			fmt.Println("Country Already Exist")
		}
	} else {
		fmt.Println("invalid input")
	}

	fmt.Println(os.Getenv("APP_NAME"))
	return

}

func CheckCityExist(country_name_a string) (counts int64) {
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("country")
	//filter := bson.D{{"country_name" ,country_name_a}}

	res, err := collection.CountDocuments(context.Background(), bson.M{"country_name": country_name_a})
	if err != nil {
		fmt.Println(err)
	}

	return res
}

// Locality function creation
func AddLocalityDetails() gin.HandlerFunc {
	return func(c *gin.Context) {

		locality_name := c.PostForm("locality_name")

		a := InsertLocalityDetailsInDB(locality_name)
		fmt.Println(a)
	}
}

func InsertLocalityDetailsInDB(locality_name string) (id int64) {
	godotenv.Load()
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("locality")
	doc := models.Locality{LocalityName: locality_name}

	pattern := regexp.MustCompile("^[a-zA-Z]*$")

	validate := validator.New()
	err := validate.Struct(doc)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}
	} else if pattern.MatchString(doc.LocalityName) {
		LocalityCount := CheckLocalityExist(locality_name)
		fmt.Println(LocalityCount)
		if LocalityCount == 0 {
			res, errr := collection.InsertOne(context.Background(), doc)
			if errr != nil {
				log.Fatal(errr)
			}
			idd := res.InsertedID
			fmt.Println("datatype", reflect.TypeOf(idd))
			fmt.Println("Inserted-Localityid", idd)
		} else {
			fmt.Println("Locality Name Already Exist")
		}
	} else {
		fmt.Println("Invalid Input")
	}
	fmt.Println(os.Getenv("APP_NAME"))
	return
}

func CheckLocalityExist(locality_name string) (counts int64) {
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("locality")

	res, err := collection.CountDocuments(context.Background(), bson.M{"locality_name": locality_name})
	if err != nil {
		fmt.Println(err)
	}
	return res
}
