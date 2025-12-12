package controller

import (
	"github.com/CHenrique-Oliv/estudo-go/src/config/rest_err"
	"github.com/gin-gonic/gin"
)

func FindUserById(c *gin.Context) {

	err := rest_err.NewBadRequestError("Erro de teste")
	c.JSON(err.Code, err)
}

func FindUserByEmail(c *gin.Context) {

}
