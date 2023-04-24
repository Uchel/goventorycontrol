package router

import (
	"database/sql"
	"go_inventory_ctrl/controllers"
	"go_inventory_ctrl/repository"
	"go_inventory_ctrl/usecase"

	"github.com/gin-gonic/gin"
)

func ProducStRoutes(router *gin.Engine, db *sql.DB) {

	//Controller ProductSt ProductSt
	productStRepo := repository.NewProductStRepo(db)
	productStUsecase := usecase.NewProductStUsecase(productStRepo)
	productStctrl := controllers.NewProductStController(productStUsecase)

	//Router Product St
	productStRoutes := router.Group("/api/v1/product-st")
	productStRoutes.GET("", productStctrl.FindAllProductsSt)
	productStRoutes.GET("/:id", productStctrl.FindProductsStById)
	productStRoutes.POST("", productStctrl.RegisterProductSt)
	productStRoutes.PUT("", productStctrl.EditProductSt)
	productStRoutes.DELETE("/:id", productStctrl.UnregProductSt)

}
