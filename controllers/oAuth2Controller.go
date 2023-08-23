package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"golang.org/x/oauth2"
)

func GoogleOAuth2GetUserInfo(c *gin.Context) {
	// Check for Authorization header
	accessToken := c.GetHeader("Authorization")
	fmt.Println(accessToken)

	conf := &oauth2.Config{}

	// Create an http.Client using the refreshed access token
	client := conf.Client(context.TODO(), &oauth2.Token{
		AccessToken: accessToken,
	})
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	name := gjson.GetBytes(content, "name").String()
	email := gjson.GetBytes(content, "email").String()

	fmt.Printf("email: %v, name: %v \n", name, email)

	// Respond with the item
	c.JSON(http.StatusOK, gin.H{
		"Name":  name,
		"Email": email,
	})
}
