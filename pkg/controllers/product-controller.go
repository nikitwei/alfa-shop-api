package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nikitwei/alfa-shop-api/pkg/models"
	"github.com/nikitwei/alfa-shop-api/pkg/utils"
	"net/http"
	"strconv"
)

var NewProduct models.MsProduct

func AddProduct(w http.ResponseWriter, r *http.Request) {
	addProduct := &models.MsProduct{}
	utils.ParseBody(r, addProduct)
	p := models.AddProduct(addProduct)
	res, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	newProduct := models.GetProducts()
	res, _ := json.Marshal(newProduct)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProductID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["productId"]
	ID, err := strconv.ParseInt(productID, 0, 0)
	if err != nil {
		fmt.Println(vars)
		fmt.Print("Error while parsing")
	}
	productDetails, _ := models.GetProductID(ID)

	res, _ := json.Marshal(productDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	updateProduct := &models.MsProduct{}
	utils.ParseBody(r, updateProduct)

	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	productDetails, db := models.GetProductID(ID)

	productDetails.UpdatedBy = 1

	if updateProduct.Name != "" {
		productDetails.Name = updateProduct.Name
	}
	if updateProduct.Description != "" {
		productDetails.Description = updateProduct.Description
	}
	if updateProduct.Image != "" {
		productDetails.Image = updateProduct.Image
	}
	if updateProduct.PriceUnit != "" {
		productDetails.PriceUnit = updateProduct.PriceUnit
	}
	if updateProduct.Price != 0 {
		productDetails.Price = updateProduct.Price
	}
	db.Save(&productDetails)

	res, _ := json.Marshal(productDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func PatchProduct(w http.ResponseWriter, r *http.Request) {
	updateProduct := &models.MsProduct{}
	utils.ParseBody(r, updateProduct)

	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	productDetails, db := models.GetProductID(ID)

	productDetails.UpdatedBy = 1

	if updateProduct.Name != "" {
		productDetails.Name = updateProduct.Name
	}
	if updateProduct.Description != "" {
		productDetails.Description = updateProduct.Description
	}
	if updateProduct.Image != "" {
		productDetails.Image = updateProduct.Image
	}
	if updateProduct.PriceUnit != "" {
		productDetails.PriceUnit = updateProduct.PriceUnit
	}
	if updateProduct.Price != 0 {
		productDetails.Price = updateProduct.Price
	}
	db.Save(&productDetails)

	res, _ := json.Marshal(productDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["productId"]
	ID, err := strconv.ParseInt(productID, 0, 0)
	if err != nil {
		fmt.Print("Error while parsing")
	}
	product := models.DeleteProduct(ID)

	res, _ := json.Marshal(product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
