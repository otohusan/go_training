package products

import (
	"go-training/application/service/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService auth.AuthService
}

func NewProductHandler(productService auth.AuthService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.productService.ParseToken("")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	// 商品作成のロジック
	c.JSON(http.StatusOK, "成功")
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	// 商品削除のロジック
}
