package service

import (
	"wgnalvian.com/test1/entity"
	"wgnalvian.com/test1/model"
)

type ProductServiceInterface interface {
	GetAllProductFromJsonFile() []entity.Product
	GetProductByIdFromJsonFile(id int) entity.Product
	WriteDataToJsonFile(product model.CreateProductRequest) entity.Product
	DeleteProductById(id int)
	UpdateProductJsonFile(id int, product map[string]interface{}) entity.Product
}
