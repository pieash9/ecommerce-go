package database

var ProductList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
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

	ProductList = append(ProductList, prd1, prd2, prd3)
}
