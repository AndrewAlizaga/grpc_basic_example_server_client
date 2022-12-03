package controller

import (
	"errors"
	"log"
	"net/http"

	accountclientv1 "github.com/AndrewAlizaga/-grpc_basic_example_grpc_client/pkg/v1/accounts"
	accountsvc1 "github.com/AndrewAlizaga/grpc_basic_example_proto/pkg/proto/v1/account"
	accountapiv1 "github.com/AndrewAlizaga/grpc_basic_example_proto/pkg/proto/v1/services/account"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context){

	var accountBody accountsvc1.Account
	if err := c.BindJSON(&accountBody); err != nil {
		log.Println("bad request on body")
		c.IndentedJSON(http.StatusBadRequest, errors.New("bad request on body"))
	}else {
		accountRequest := accountapiv1.SignUpRequest {
			Account: &accountBody,
		}

		res, err := accountclientv1.AccountSignUp(&accountRequest)

		if err != nil {
			log.Println("bad response from server")
			c.IndentedJSON(http.StatusForbidden, err.Error())
			return
		}

		if res.GetError() != ""{
			log.Println("bad response from server")
			c.IndentedJSON(http.StatusForbidden, res.GetError())
			return
		} 

		log.Println("successful signup")
		c.IndentedJSON(http.StatusAccepted, res)

	}

}


func Login(c *gin.Context){
	var accountRequest accountapiv1.LoginRequest

	if err := c.BindJSON(&accountRequest); err != nil {
		log.Println("bad request on body")
		c.IndentedJSON(http.StatusBadRequest, errors.New("bad request on body"))
		
	}else {
		

		res, err := accountclientv1.AccountLogin(&accountRequest)

		if err != nil {
			log.Println("bad response from server")
			c.IndentedJSON(http.StatusForbidden, err.Error())
			return
		}

		if res.GetError() != ""{
			log.Println("bad response from server")
			c.IndentedJSON(http.StatusForbidden, res)
			return
		} 

		log.Println("successful login")
		c.IndentedJSON(http.StatusAccepted, res)

	}

}