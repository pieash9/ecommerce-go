package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")

	pId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == pId {
			util.SendData(w, product, http.StatusOK)
			return
		}
	}

	util.SendData(w, "NO data found with id "+strconv.Itoa(pId), http.StatusNotFound)
}
