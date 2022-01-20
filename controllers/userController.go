package controller

import (
	"context"
	"golang-restaurent-management/database"
	model "golang-restaurent-management/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 != nil {
			page = 1
		}

		startIndex := (page-1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{{}}}}
		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"total_cound", 1},
				{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
			}}}

			result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
				matchStage, projectStage})
				defer cancel()
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
				}

				var allUsers []bson.M
				if err = result.All(ctx, &allUsers); err != nil {
					log.Fatal(err)
				}
				c.JSON(http.StatusOK, allUsers[0])
 	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		userId := c.Param("user_id")

		var user model.User

		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
		}
		c.JSON(http.StatusOK, user)
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		//convert the JSON data coming from postman to something that golang understands

		//validate the data based on user struct

		//You'll check if the email has already been used by another user

		//hash password

		//You'll also check if the phone number has already been used by another user

		//create some extra details for the user object - created_at, updated_at, ID

		//generate token and refresh token (generate all tokens and functions from helper)

		//if all ok, then you insert this new user into the users collection

		//return status ok and the result back
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		//convert the login data from postman which is in JSON to going to golang readable format

		//find a user with that email and see if that user even exists

		//then you will verify the password

		//if all goes well, then you'll generate tokens

		//update tokens - token and refresh token

		//return statuspok
	}
}

func HashPassword(password string) string {

}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {

}
