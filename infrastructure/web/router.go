package web

import (
	"codifin-challenge/infrastructure/web/responses"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func (s *Server) setRoutes() {
	s.router.Use(s.middlewares.AddCORS())

	s.router.GET("", func(c *gin.Context) {
		responses.SendSuccess(c, http.StatusOK, gin.H{"message": "El servicio está en línea"})
		return
	})

	s.router.GET("ping", func(c *gin.Context) {
		responses.SendSuccess(c, http.StatusOK, gin.H{"message": "pong"})
		return
	})

	s.addV1Routes()
}

func (s *Server) addV1Routes() {
	v1 := s.router.Group("v1")
	v1.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	products := v1.Group("products")
	products.GET("", s.controllers.productCtrl.FindProducts)
	products.POST("", s.controllers.productCtrl.NewProduct)

	product := v1.Group("product")
	product.GET(":id", s.controllers.productCtrl.FindProduct)
	product.PATCH(":id", s.controllers.productCtrl.UpdateProduct)
	product.DELETE(":id", s.controllers.productCtrl.RemoveProduct)

	carts := v1.Group("carts")
	carts.POST("", s.controllers.shoppingCartCtrl.NewCart)

	cart := v1.Group("cart")
	cart.POST(":id/items", s.controllers.shoppingCartCtrl.AddItem)
	cart.DELETE(":id/items", s.controllers.shoppingCartCtrl.RemoveItems)
	cart.GET(":id", s.controllers.shoppingCartCtrl.FindShoppingCart)

}
