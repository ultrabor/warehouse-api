package v1

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ultrabor/warehouse-api/internal/domain"
)

type ProductUseCase interface {
	Create(ctx context.Context, p *domain.Product) (int64, error)
	GetByID(ctx context.Context, id int64) (*domain.Product, error)
	Update(ctx context.Context, p *domain.Product) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]domain.Product, error)
}

type ProductHandler struct {
	useCase domain.ProductUseCase
}

func NewProductHandler(u domain.ProductUseCase) *ProductHandler {
	return &ProductHandler{useCase: u}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var p domain.Product

	if err := c.ShouldBindJSON(&p); err != nil {
		handleError(c, err)
		return
	}

	id, err := h.useCase.Create(c.Request.Context(), &p)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})

}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handleError(c, err)
		return
	}

	product, err := h.useCase.GetByID(c.Request.Context(), id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.useCase.GetAll(c.Request.Context())
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handleError(c, err)
		return
	}

	var p domain.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		handleError(c, err)
		return
	}

	p.ID = id

	err = h.useCase.Update(c.Request.Context(), &p)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": p.ID})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handleError(c, err)
		return
	}

	err = h.useCase.Delete(c.Request.Context(), id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
