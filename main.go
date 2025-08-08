package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	fmt.Fprintln(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am pieash, I am a software engineer.")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {

	handleCors(w)
	handlePreflightReq(w, r)

	if r.Method != http.MethodGet {
		http.Error(w, "Plz give me e get req", 400)
		return
	}

	sendData(w, productList, http.StatusOK)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightReq(w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "Plz give me a post req", 400)
		return
	}

	var newProduct Product
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println("Error decoding product:", err)
		http.Error(w, "Invalid product data", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	sendData(w, newProduct, http.StatusCreated)
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Client-Id")
	w.Header().Set("Content-Type", "Application/json")
}

func handlePreflightReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /hello", http.HandlerFunc(helloHandler)) // mux mean router

	mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc("/products", getProducts)

	mux.HandleFunc("/create-products", createProduct)

	fmt.Println("Server is running on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "This is a fresh orange",
		Price:       1.99,
		ImgUrl:      "https://static.vecteezy.com/system/resources/thumbnails/028/180/885/small_2x/orange-fruit-isolated-on-white-background-whole-orange-citrus-fruit-ai-generative-photo.jpg",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "This is a fresh apple",
		Price:       2.49,
		ImgUrl:      "https://media.istockphoto.com/id/184276818/photo/red-apple.jpg?s=612x612&w=0&k=20&c=NvO-bLsG0DJ_7Ii8SSVoKLurzjmV0Qi4eGfn6nW3l5w=",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "This is a fresh banana",
		Price:       1.29,
		ImgUrl:      "https://www.shutterstock.com/image-photo/bunch-bananas-isolated-on-white-600nw-1722111529.jpg",
	}

	productList = append(productList, prd1, prd2, prd3)
}
