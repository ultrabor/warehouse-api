package v1

import "github.com/gin-gonic/gin"

func (h *ProductHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/products", h.CreateProduct)
	r.GET("/products/:id", h.GetProduct)
	r.PUT("/products/:id", h.UpdateProduct)
	r.DELETE("/products/:id", h.DeleteProduct)
	r.GET("/products", h.ListProducts)
}
