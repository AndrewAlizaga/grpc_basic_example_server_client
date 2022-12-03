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
			c.IndentedJSON(http.StatusForbidden, err)
		}

		if res.GetError() != ""{
			log.Println("bad response from server")
			c.IndentedJSON(http.StatusForbidden, err)
		} 

		log.Println("successful signup")
		c.IndentedJSON(http.StatusAccepted, res)

	}

	return
}


func Login(c *gin.Context){
	
}