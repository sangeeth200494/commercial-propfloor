package manage_commercial

import (
	"commercial-propfloor/database"
	"commercial-propfloor/models"
	"context"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/go-playground/validator/v10"

	"reflect"

	"github.com/joho/godotenv"

	//"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
)

func SaveBasicBuildingDetails() gin.HandlerFunc {
	return func(c *gin.Context) {

		name := c.PostForm("name")
		location := c.PostForm("location")
		availablity_for := c.PostForm("availability_for")
		completion_status := c.PostForm("completion_status")
		furnishing_status := c.PostForm("furnishing_status")
		no_of_floor := c.PostForm("no_of_floor")
		parking := c.PostForm("parking")
		oc := c.PostForm("oc")
		lift := c.PostForm("lift")
		overlooking := c.PostForm("overlooking")

		InsertBasicBuildingDetailInDB(name, location, availablity_for, completion_status, furnishing_status, no_of_floor, parking, oc, lift, overlooking)

	}
}

func InsertBasicBuildingDetailInDB(name string, location string, availablity_for string, completion_status string,
	furnishing_status string, no_of_floor string, parking string, oc string, lift string, overlooking string) (id int64) {
	godotenv.Load()
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	collection := client.Database("building_details").Collection("commertial_building")
	doc := models.Property{Name: name, Location: location, AvailabilityFor: availablity_for, CompletionStatus: completion_status,
		FurnishingStatus: furnishing_status, Floors: no_of_floor, Parking: parking, Oc: oc, Lift: lift, OverLooking: overlooking}

	pattern := regexp.MustCompile("^[A-Za-z]*$")
	pattern2 := regexp.MustCompile("^[1-9][0-9]*$")

	validate := validator.New()
	err := validate.Struct(doc)
	//var  string
	//fmt.Scanln(&doc)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
			fmt.Println("ankit")
		}

	} else if pattern.MatchString(doc.Name) && pattern.MatchString(doc.Location) && pattern.MatchString(doc.AvailabilityFor) && pattern.MatchString(doc.CompletionStatus) && pattern.MatchString(doc.FurnishingStatus) && pattern.MatchString(doc.Parking) && pattern.MatchString(doc.Oc) && pattern.MatchString(doc.Lift) && pattern.MatchString(doc.OverLooking) && pattern2.MatchString(doc.Floors) {
		res, errr := collection.InsertOne(context.Background(), doc)
		if errr != nil {
			log.Fatal(errr)
		}
		idd := res.InsertedID
		fmt.Println("datatype", reflect.TypeOf(idd))
		fmt.Println("inserted-cityid  ", idd)
	} else {
		fmt.Println("invalid input")
	}

	fmt.Println(os.Getenv("APP_NAME"))
	return
}
