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
		availability_for := c.PostForm("availability_for")
		completion_status := c.PostForm("completion_status")
		furnishing_status := c.PostForm("furnishing_status")
		no_of_floor := c.PostForm("no_of_floor")
		parking := c.PostForm("parking")
		lift := c.PostForm("lift")
		oc := c.PostForm("oc")
		overlooking := c.PostForm("overlooking")
		InsertBasicBuildingDetailInDB(name, location, availability_for, completion_status, furnishing_status, no_of_floor, parking, lift, oc, overlooking)
	}
}
func InsertBasicBuildingDetailInDB(name string, location string, availability_for string, completion_status string, furnishing_status string, no_of_floor string, parking string, lift string, oc string, overlooking string) (id int64) {
	godotenv.Load()
	client, _, _, _ := database.Mongoconnect(os.Getenv("MONGODB_HOST"))
	fmt.Println(os.Getenv("APP_NAME"))
	collection := client.Database("building_details").Collection("commertial_building")
	doc := models.Property{Name: name, Location: location, AvailabilityFor: availability_for, CompletionStatus: completion_status, FurnishingStatus: furnishing_status, Floors: no_of_floor, Parking: parking, Lift: lift, Oc: oc, OverLooking: overlooking}
	pattern := regexp.MustCompile("^[a-zA-Z]*$")
	pattern2 := regexp.MustCompile("^[1-9][0-9]*$")
	validate := validator.New()
	err := validate.Struct(doc)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}
	} else if pattern2.MatchString(doc.Floors) && pattern.MatchString(doc.Name) && pattern.MatchString(doc.AvailabilityFor) && pattern.MatchString(doc.Location) && pattern.MatchString(doc.CompletionStatus) && pattern.MatchString(doc.FurnishingStatus) && pattern.MatchString(doc.Parking) && pattern.MatchString(doc.Lift) && pattern.MatchString(doc.Oc) && pattern.MatchString(doc.OverLooking) {
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
