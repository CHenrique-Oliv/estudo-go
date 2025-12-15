package controller

import (
	"fmt"
	"log"

	"github.com/CHenrique-Oliv/estudo-go/src/config/validation"
	"github.com/CHenrique-Oliv/estudo-go/src/controller/model/request"
	"github.com/CHenrique-Oliv/estudo-go/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("There are some incorrect filds, error=%s", err.Error())
		errRest := validation.ValidateUserErro(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:    "teste",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	log.Print(response)
}
