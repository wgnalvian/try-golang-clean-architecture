package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"wgnalvian.com/test1/entity"
	"wgnalvian.com/test1/exception"
	"wgnalvian.com/test1/model"
	"wgnalvian.com/test1/service"
)

type ProductController struct {
	ProductService service.ProductServiceInterface
}

func (p *ProductController) GetAll(c *gin.Context) {

	var products []entity.Product = p.ProductService.GetAllProductFromJsonFile()
	var productResponses []model.GetProductResponse
	for _, product := range products {
		productResponse := model.GetProductResponse{
			Id:    product.Id,
			Name:  product.Name,
			Price: product.Price,
		}
		productResponses = append(productResponses, productResponse)
	}

	c.JSON(200, gin.H{"status": 200, "products": productResponses})
}

func (p *ProductController) GetProductById(c *gin.Context) {
	i, err := strconv.Atoi(c.Param("id"))
	exception.PanicIfNeeded(err)
	product := p.ProductService.GetProductByIdFromJsonFile(i)
	productResponse := model.GetProductResponse{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
	}
	c.JSON(200, gin.H{"status": 200, "product": productResponse})

}

func (p *ProductController) CreateProduct(c *gin.Context) {
	var productRequest model.CreateProductRequest
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		exception.PanicIfNeeded(err)
	}
	product := p.ProductService.WriteDataToJsonFile(productRequest)

	c.JSON(201, gin.H{"status": 201, "product": product})

}

func (p *ProductController) DeleteProduct(c *gin.Context) {
	i, err := strconv.Atoi(c.Param("id"))
	exception.PanicIfNeeded(err)
	p.ProductService.DeleteProductById(i)
	c.JSON(204, nil)
}

func (p *ProductController) UpdateProduct(c *gin.Context) {
	i, err := strconv.Atoi(c.Param("id"))
	exception.PanicIfNeeded(err)
	var productRequest map[string]interface{}
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		exception.PanicIfNeeded(err)
	}
	product := p.ProductService.UpdateProductJsonFile(i, productRequest)
	c.JSON(200, gin.H{"status": 200, "product": product})
}

func (p *ProductController) Route(r *gin.Engine) {
	r.GET("/products", p.GetAll)
	r.GET("/product/:id", p.GetProductById)
	r.POST("/product", p.CreateProduct)
	r.DELETE("/product/:id", p.DeleteProduct)
	r.PATCH("product/:id", p.UpdateProduct)

}
