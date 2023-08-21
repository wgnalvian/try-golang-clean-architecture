package service

import (
	"encoding/json"
	"os"
	"path/filepath"

	"wgnalvian.com/test1/entity"
	"wgnalvian.com/test1/exception"
	"wgnalvian.com/test1/model"
)

func NewProductService() ProductServiceInterface {
	return &productService{}
}

type productService struct {
}

func (p *productService) GetAllProductFromJsonFile() []entity.Product {

	return readJsonFileProducts()
}

func (p *productService) GetProductByIdFromJsonFile(id int) entity.Product {
	var products []entity.Product = readJsonFileProducts()
	var targetProduct entity.Product
	for _, product := range products {
		if product.Id == id {
			targetProduct = product
		}
	}
	return targetProduct
}

func (p *productService) DeleteProductById(id int) {

	products := readJsonFileProducts()
	var productsUpdate []entity.Product

	for _, item := range products {
		// exception.PanicIfNeeded(err)
		if id != item.Id {
			productsUpdate = append(productsUpdate, item)
		}
	}
	writeDataToJsonWithParam(productsUpdate)

}

func (p *productService) WriteDataToJsonFile(product model.CreateProductRequest) entity.Product {
	products := readJsonFileProducts()
	productEntity := entity.Product{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
	}

	products = append(products, productEntity)
	jsonData, err := json.MarshalIndent(products, "", "  ")
	exception.PanicIfNeeded(err)
	filePath := getCurrentPath()
	err = os.WriteFile(filePath, jsonData, 0644)
	return productEntity
}

func (p *productService) UpdateProductJsonFile(id int, product map[string]interface{}) entity.Product {
	var productsUpdate []entity.Product
	products := readJsonFileProducts()
	productTarget := p.GetProductByIdFromJsonFile(id)
	if _, ok := product["name"].(string); ok {

		productTarget.Name = product["name"].(string)
	}
	if _, ok := product["price"].(float64); ok {

		productTarget.Price = int(product["price"].(float64))
	}
	for _, i := range products {
		if i.Id != id {
			productsUpdate = append(productsUpdate, i)
		}
	}
	productsUpdate = append(productsUpdate, productTarget)
	writeDataToJsonWithParam(productsUpdate)
	return productTarget

}

func writeDataToJsonWithParam(products []entity.Product) {
	jsonData, err := json.MarshalIndent(products, "", "  ")
	exception.PanicIfNeeded(err)
	filePath := getCurrentPath()
	err = os.WriteFile(filePath, jsonData, 0644)

}

func readJsonFileProducts() []entity.Product {
	filePath := getCurrentPath()
	file, err := os.Open(filePath)
	exception.PanicIfNeeded(err)

	defer file.Close()

	var products []entity.Product

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&products)
	exception.PanicIfNeeded(err)

	return products
}

func getCurrentPath() string {
	currentDir, err := os.Getwd()
	exception.PanicIfNeeded(err)

	filePath := filepath.Join(currentDir, "data", "product.json")
	return filePath
}
