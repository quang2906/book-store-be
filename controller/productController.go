package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/quang2906/book_store_be/model"
	repo "github.com/quang2906/book_store_be/repository"
)

const limit int = 6

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products := repo.GetAllProducts()
	responseWithJSON(w, http.StatusOK, products)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid product id"})
		return
	}
	products := repo.GetAllProducts()
	for _, product := range products {
		if int(product.Id) == id {
			responseWithJSON(w, http.StatusOK, product)
			return
		}
	}
	responseWithJSON(w, http.StatusNotFound, map[string]string{"message": "product not found"})
}

func GetProductByCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid product id"})
		return
	}
	products := repo.GetAllProducts()
	for _, product := range products {
		if int(product.Id) == id {
			responseWithJSON(w, http.StatusOK, product)
			return
		}
	}
	responseWithJSON(w, http.StatusNotFound, map[string]string{"message": "product not found"})
}

func CreateProduct(writer http.ResponseWriter, request *http.Request) {
	var newProduct *model.Product
	if err := json.NewDecoder(request.Body).Decode(&newProduct); err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	repo.CreateNewProduct(newProduct)
	responseWithJSON(writer, http.StatusCreated, newProduct)
}

func UpdateProductById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid category id"})
		return
	}

	var updateProducts *model.Product
	if err := json.NewDecoder(request.Body).Decode(&updateProducts); err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	updateProducts.Id = int64(id)

	products := repo.GetAllProducts()
	for _, product := range products {
		if product.Id == int64(id) {
			repo.UpdateProductById(int64(id), updateProducts)
			responseWithJSON(writer, http.StatusOK, updateProducts)
			return
		}
	}

	responseWithJSON(writer, http.StatusNotFound, map[string]string{"message": "Product not found"})
}

func DeleteProductById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid Product id"})
		return
	}
	err = repo.DeleteProductById(int64(id))
	if err != nil {
		responseWithJSON(writer, http.StatusNotFound, map[string]string{"message": "Product not found"})
		return
	}
	responseWithJSON(writer, http.StatusOK, map[string]string{"message": "Product was deleted"})
}

func SearchProduct(writer http.ResponseWriter, request *http.Request) {

	// lay param
	nameProduct := request.URL.Query().Get("name")
	page, _ := strconv.Atoi(request.URL.Query().Get("page"))

	//check loai bo ki tu dac biet
	nameProduct = strings.Replace(nameProduct, "%", "", -1)
	nameProduct = strings.Replace(nameProduct, "-", "", -1)

	totalProduct := repo.TotalProduct(nameProduct)
	pageMax := pageMax(totalProduct)

	// tinh pagemax de ko cho goi hon so page co
	if page > pageMax {
		page = pageMax
	}

	pageOffset := pageRequest(page)

	products := repo.SearchProductRepo(nameProduct, pageOffset)

	res := model.ResponseProduct{
		TotalPage:    pageMax,
		TotalProduct: totalProduct,
		PageIndex:    page,
		Products:     products,
	}
	responseWithJSON(writer, http.StatusOK, res)
}
func SortProduct(writer http.ResponseWriter, request *http.Request) {
	// lay param
	sortProduct := request.URL.Query().Get("sort")
	page, _ := strconv.Atoi(request.URL.Query().Get("page"))
	// lay count
	totalProduct := repo.TotalProduct("")
	pageMax := pageMax(totalProduct)
	// tinh pagemax de ko cho goi hon so page co
	if page > pageMax {
		page = pageMax
	}
	pageOffset := pageRequest(page)
	fmt.Println(sortProduct)
	products := repo.SortProductRepo(sortProduct, pageOffset)
	if products == nil {
		responseWithJSON(writer, http.StatusNotFound, map[string]string{"messeage": "Sort not found"})
		return
	}

	res := model.ResponseProduct{
		TotalPage:    pageMax,
		TotalProduct: totalProduct,
		PageIndex:    page,
		Products:     products,
	}
	responseWithJSON(writer, http.StatusOK, res)
}

func pageRequest(page int) int {
	offset := (page - 1) * limit
	return offset
}
func pageMax(totalProduct int) int {
	var pageMax int
	if totalProduct%limit == 0 {
		pageMax = totalProduct / limit
	} else {
		pageMax = totalProduct/limit + 1
	}
	return pageMax
}
