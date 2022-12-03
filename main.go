package main

import (
	"log"

	controllers "github.com/AndrewAlizaga/grpc_basic_example_server_client/internal/v1/controller"
	"github.com/gin-gonic/gin"
)

func main(){


}

func startGinServer(){
	port := "8081" 
	router := gin.Default()

	//Account Routes
	account := router.Group("/account")
	{

		//SignUp
		account.POST("/signup", controllers.SignUp)

		//Login
		account.POST("/login", controllers.Login)
	}

	router.Run(":"+port)
	log.Println("RUNNING GIN SERVER AT ", port)
}