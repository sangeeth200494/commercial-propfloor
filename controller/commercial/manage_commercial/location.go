package manage_commercial

import (
	"commercial-propfloor/controller"
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
		country_name := c.PostForm("country_name") // taking input from user

		currency_name := c.PostForm("currency_name")
		currency_symbol := c.PostForm("currency_symbol")

		/*a := */
		InsertCountryDetailsInDB(country_name, currency_name, currency_symbol)
		//fmt.Println(a)
	}
}

func InsertCountryDetailsInDB(country_name string, currency_name string, currency_symbol string) (id int64) {
	godotenv.Load()
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))                                            // connecting mongodb
	fmt.Println(os.Getenv("APP_NAME"))                                                                             //  printing the value of app_name
	collection := client.Database("india").Collection("country")                                                   //function inserts the doc document into the "country" collection of the "india" database using
	doc := models.Country{CountryName: country_name, CurrencyName: currency_name, CurrencySymbol: currency_symbol} // coonecting/calling struct calls Destinations.

	pattern := regexp.MustCompile("^[a-zA-Z]*$") // regexp is a package for writing regular expression  Regular expressions are commonly used to search for specific patterns in text, to validate user input

	validate := validator.New() // function creates a new validator ,  which is used to validate structs and fields based on tags.
	err := validate.Struct(doc) // The validate.Struct(doc) function validates the doc struct using the tags defined on its fields. It returns an error if any of the fields fail validation.
	//controller.ErrorLogger.Println("Something is Missing", err)
	//controller.ReadFileError()
	// 	var a = 2
	//fmt.Println("error is **************", err)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) { // ValidationErrors is an array of FieldError's for use in custom error messages post validation.
			controller.ErrorLogger.Println("Validation Error", err.Field(), err.Tag())
			controller.ReadFileError()
		}
	} else if pattern.MatchString(doc.CountryName) && pattern.MatchString(doc.CurrencyName) && pattern.MatchString(doc.CurrencySymbol) {
		CountryCount := CheckCountryExist(country_name) // function calls the CheckCountryExist() function to check if the country already exists in the database.
		fmt.Println(CountryCount)
		if CountryCount == 0 {
			res, errr := collection.InsertOne(context.Background(), doc)
			//fmt.Println(res)
			if errr != nil {
				controller.ErrorLogger.Fatal("invalid input", errr) // function is called to log the error and terminate the program
				controller.ReadFileError()
			}
			idd := res.InsertedID
			fmt.Println("datatype", reflect.TypeOf(idd)) //check the data type
			fmt.Println("country inserted", idd)
			controller.InfoLogger.Println("inserted-Destinationid") // for displaying information about the inserted document to the user.
			controller.ReadFileInfo()
		} else {
			controller.WarningLogger.Println("Country Already Exist")
			controller.ReadFileWarning()
		}
	} else {
		controller.WarningLogger.Println("invalid input")
		controller.ReadFileWarning()
	}

	controller.InfoLogger.Println(os.Getenv("APP_NAME"))
	controller.ReadFileInfo()
	return

}

func CheckCountryExist(country_name_a string) (counts int64) {
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("country")
	//filter := bson.D{{"country_name" ,country_name_a}}
	res, err := collection.CountDocuments(context.Background(), bson.M{"country_name": country_name_a}) // binary-encoded serialization format that is used to store and exchange documents in MongoDB
	if err != nil {
		controller.ErrorLogger.Println("Country Doesn't Exist", err)
		controller.ReadFileError()

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
			controller.ErrorLogger.Fatal("Validation Error", err.Field(), err.Tag())
			controller.ReadFileError()
		}

	} else if pattern.MatchString(doc.StateName) && pattern.MatchString(doc.StateLanguage) {
		StateCount := CheckStateExist(state_name)
		if StateCount == 0 {
			res, errr := collection.InsertOne(context.Background(), doc)
			if errr != nil {
				controller.ErrorLogger.Fatal("Invalid Input", errr)
				controller.ReadFileError()
			}
			idd := res.InsertedID
			fmt.Println("datatype", reflect.TypeOf(idd))
			controller.InfoLogger.Println("Inserted-State")
			controller.ReadFileInfo()
		} else {
			controller.WarningLogger.Println("State Already Exist")
			controller.ReadFileWarning()
		}
	} else {
		controller.WarningLogger.Println("Invalid Input")
		controller.ReadFileWarning()
	}

	controller.InfoLogger.Println(os.Getenv("APP_NAME"))
	controller.ReadFileInfo()
	return

}

func CheckStateExist(state_name string) (counts int64) {
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("state")

	res, err := collection.CountDocuments(context.Background(), bson.M{"state_name": state_name}) //
	if err != nil {
		controller.ErrorLogger.Println("Country Doesn't Exist", err)
		controller.ReadFileError()
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
			controller.ErrorLogger.Fatal("Validation Error", err.Field(), err.Tag())
			controller.ReadFileError()
		}
	} else if pattern.MatchString(doc.CityName) {
		CityCount := CheckCityExist(city_name)
		fmt.Println(CityCount)
		if CityCount == 0 {
			res, errr := collection.InsertOne(context.Background(), doc)
			//fmt.Println(res)
			if errr != nil {
				controller.ErrorLogger.Fatal("Invalid Input", errr)
				controller.ReadFileError()
			}
			idd_c := res.InsertedID
			fmt.Println("datatype", reflect.TypeOf(idd_c))
			controller.InfoLogger.Println("inserted-Destinationid")
			controller.ReadFileInfo()
		} else {
			controller.WarningLogger.Println("City Already Exist")
			controller.ReadFileWarning()
		}
	} else {
		controller.WarningLogger.Println("invalid input")
		controller.ReadFileWarning()
	}

	controller.InfoLogger.Println(os.Getenv("APP_NAME"))
	controller.ReadFileInfo()
	return

}

// creating city Get function
func SelectCityindatabase() gin.HandlerFunc {
	return func(c *gin.Context) {
		godotenv.Load()
		client, ctx, cancel, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
		fmt.Println(os.Getenv("APP_NAME"))
		collection := client.Database("MONGODB_NAME_LOCATION").Collection("city")
		//doc := models.City{CityName: city_name}
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			controller.ErrorLogger.Println("City Doesn't Exist", err)
			controller.ReadFileError()
		}
		var city []bson.M
		if err = cursor.All(ctx, &city); err != nil {
			log.Println(err)
		}
		fmt.Println()
		//fmt.Println("datatype", reflect.TypeOf(building))
		//fmt.Println("cities\n", building)
		//fmt.Println("****************************************************")
		// for i := range city {
		// 	fmt.Println(city[i]["city_name"])
		// }

		//fmt.Println("selected city-----\n", building[0]["city_name"])
		//fmt.Println("****************************************************")
		database.Mongoclose(client, ctx, cancel)
		fmt.Println(city)
		c.IndentedJSON(200, city)
	}

}

func CheckCityExist(country_name_a string) (counts int64) {
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("country")
	//filter := bson.D{{"country_name" ,country_name_a}}

	res, err := collection.CountDocuments(context.Background(), bson.M{"country_name": country_name_a})
	if err != nil {
		controller.ErrorLogger.Println("Country Doesn't Exist", err)
		controller.ReadFileError()

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
			controller.ErrorLogger.Fatal("Validation Error", err.Field(), err.Tag())
			controller.ReadFileError()
		}
	} else if pattern.MatchString(doc.LocalityName) {
		LocalityCount := CheckLocalityExist(locality_name)
		fmt.Println(LocalityCount)
		if LocalityCount == 0 {
			res, errr := collection.InsertOne(context.Background(), doc)
			if errr != nil {
				controller.ErrorLogger.Fatal("Invalid Input", errr)
				controller.ReadFileError()
			}
			idd := res.InsertedID
			fmt.Println("datatype", reflect.TypeOf(idd))
			controller.InfoLogger.Println("Inserted-Localityid")
			controller.ReadFileInfo()
		} else {
			controller.WarningLogger.Println("Locality Name Already Exist")
			controller.ReadFileWarning()
		}
	} else {
		controller.WarningLogger.Println("Invalid Input")
		controller.ReadFileWarning()
	}
	controller.InfoLogger.Println(os.Getenv("APP_NAME"))
	controller.ReadFileInfo()
	return
}

func CheckLocalityExist(locality_name string) (counts int64) {
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("india").Collection("locality")

	res, err := collection.CountDocuments(context.Background(), bson.M{"locality_name": locality_name})
	if err != nil {
		controller.ErrorLogger.Println("Country Doesn't Exist", err)
		controller.ReadFileError()
	}
	return res
}
