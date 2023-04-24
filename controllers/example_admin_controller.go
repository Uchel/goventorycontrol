package controllers

import (
	"fmt"
	"log"
	"net/http"

	"go_inventory_ctrl/entity"
	"go_inventory_ctrl/usecase"

	// "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ExampleaAdminController struct {
	usecase usecase.ExampleAdminUsecase
}

func (c *ExampleaAdminController) FindAllExampleAdmin(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)

	res := c.usecase.FindAllExampleAdmin()
	ctx.JSON(http.StatusOK, gin.H{
		"data":  res,
		"email": email,
	})
}

func (c *ExampleaAdminController) FindByIdExampleAdmin(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)

	id := ctx.Param("id")
	fmt.Println(id)
	res := c.usecase.FindByIdExampleAdmin(id)
	ctx.JSON(http.StatusOK, gin.H{
		"data":  res,
		"email": email,
	})
}

func (c *ExampleaAdminController) RegisterExampleAdmin(ctx *gin.Context) {

	file, err := ctx.FormFile("photo")
	if err != nil {
		log.Println(err)
	}
	filename := file.Filename

	var newExampleAdmin entity.ExampleAdmin
	newExampleAdmin.Photo = filename

	if errFile := ctx.SaveUploadedFile(file, fmt.Sprintf("./images/%s", filename)); errFile != nil {
		log.Println(errFile)
	}

	if err := ctx.ShouldBind(&newExampleAdmin); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid input")
		fmt.Println(err)
		return
	}

	res := c.usecase.RegisterExampleAdmin(&newExampleAdmin)

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
		// "email":email,
	})
}

func (c *ExampleaAdminController) EditExampleAdmin(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)

	var exampleAdmin entity.ExampleAdmin

	if err := ctx.BindJSON(&exampleAdmin); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.EditExampleAdmin(&exampleAdmin)
	ctx.JSON(http.StatusOK, gin.H{
		"data":  res,
		"email": email,
	})
}

func (c *ExampleaAdminController) UnregExampleAdmin(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)

	id := ctx.Param("id")

	res := c.usecase.UnregExampleAdmin(id)
	ctx.JSON(http.StatusOK, gin.H{
		"data":  res,
		"email": email,
	})
}

func NewExampleAdminController(u usecase.ExampleAdminUsecase) *ExampleaAdminController {
	controller := ExampleaAdminController{
		usecase: u,
	}

	return &controller
}
