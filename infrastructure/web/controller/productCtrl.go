package controller

import (
	"codifin-challenge/domain/service"
	"codifin-challenge/domain/utils"
	"codifin-challenge/infrastructure/web/dto"
	"codifin-challenge/infrastructure/web/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

// FindProducts
// @Summary Get a paginated and filtered list of products
// @Description Retrieves a paginated and filtered list of products based on search criteria
// @Tags Products
// @ID find-products
// @Accept json
// @Produce json
// @Param page query int true "Page number" minimum(1) "The page number for pagination"
// @Param pageSize query int true "Page size" minimum(1) "The number of products per page"
// @Param searchTerm query string false "Search term to filter products by code or name"
// @Param orderBy query string false "Field to order results by. Can be 'name', 'price', or 'code'. Default is 'price'."
// @Param ascending query bool false "Whether to order results in ascending or descending order. Default is true."
// @Success 200 {object} dto.ProductsListResp "Paginated and filtered list of products"
// @Failure 400 {object} responses.ErrorDTO "Invalid page, pageSize, or search parameters"
// @Failure 500 {object} responses.ErrorDTO "Failed to retrieve products"
// @Router /products [get]
func (ctrl *ProductController) FindProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	searchTerm := c.Query("searchTerm")
	orderBy := c.DefaultQuery("orderBy", "price")
	ascending, _ := strconv.ParseBool(c.DefaultQuery("ascending", "true"))

	products, total, err := ctrl.productService.ProductsList(page, pageSize, searchTerm, orderBy, ascending)
	if err != nil {
		responses.SendError(c, utils.GetCustomError(err))
		return
	}

	resp := dto.ProductsListResp{
		Total:    total,
		Products: dto.ToProductsDTO(products),
	}

	responses.SendSuccess(c, http.StatusOK, resp)
}

// FindProduct
// @Summary Get a product by ID
// @Description Retrieves a product by its ID
// @Tags Products
// @ID find-product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} dto.ProductDTO "Product found"
// @Failure 400 {object} responses.ErrorDTO "Invalid product ID"
// @Failure 404 {object} responses.ErrorDTO "Product does not exist"
// @Failure 500 {object} responses.ErrorDTO "Failed to retrieve product"
// @Router /product/{id} [get]
func (ctrl *ProductController) FindProduct(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Param("id"))

	product, err := ctrl.productService.ProductByID(uint(productID))
	if err != nil {
		responses.SendError(c, utils.GetCustomError(err))
		return
	}

	productDTO := dto.ToProductDTO(product)
	responses.SendSuccess(c, http.StatusOK, productDTO)
}

// NewProduct
// @Summary Create a new product
// @Description Creates a new product with the provided data
// @Tags Products
// @ID new-product
// @Accept json
// @Produce json
// @Param data body dto.ProductData true "Product data"
// @Success 201 {object} dto.ProductDTO "Created product"
// @Failure 400 {object} responses.ErrorDTO "Invalid product data"
// @Failure 500 {object} responses.ErrorDTO "Failed to create product"
// @Router /products [post]
func (ctrl *ProductController) NewProduct(c *gin.Context) {
	var productData dto.ProductData

	if err := c.BindJSON(&productData); err != nil {
		responses.SendError(c, utils.ToUserError(http.StatusBadRequest, "Datos de producto incorrectos", err))
		return
	}

	newProduct := productData.ToProduct()

	if err := ctrl.productService.CreateProduct(newProduct); err != nil {
		responses.SendError(c, utils.GetCustomError(err))
		return
	}

	productDTO := dto.ToProductDTO(newProduct)
	responses.SendSuccess(c, http.StatusCreated, productDTO)
}

// UpdateProduct
// @Summary Update a product
// @Description Updates a product with the provided updates
// @Tags Products
// @ID update-product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param updates body dto.ProductData true "Product updates"
// @Success 200 {object} responses.SuccessDTO "Product updated successfully"
// @Failure 400 {object} responses.ErrorDTO "Invalid product ID or updates"
// @Failure 500 {object} responses.ErrorDTO "Failed to update product"
// @Router /product/{id} [patch]
func (ctrl *ProductController) UpdateProduct(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Param("id"))

	var updates map[string]interface{}
	if err := c.BindJSON(&updates); err != nil {
		responses.SendError(c, utils.ToUserError(http.StatusBadRequest, "Datos de actualizacion de producto incorrectos", err))
		return
	}

	if err := ctrl.productService.UpdateProduct(uint(productID), updates); err != nil {
		responses.SendError(c, utils.GetCustomError(err))
		return
	}

	resp := responses.SuccessDTO{Message: "Producto actulizado correctamente"}
	responses.SendSuccess(c, http.StatusOK, resp)
}

// RemoveProduct
// @Summary Delete a product
// @Description Deletes a product by its ID
// @Tags Products
// @ID remove-product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} responses.SuccessDTO "Product deleted successfully"
// @Failure 400 {object} responses.ErrorDTO "Invalid product ID"
// @Failure 500 {object} responses.ErrorDTO "Failed to delete product"
// @Router /product/{id} [delete]
func (ctrl *ProductController) RemoveProduct(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Param("id"))

	if err := ctrl.productService.DeleteProduct(uint(productID)); err != nil {
		responses.SendError(c, utils.GetCustomError(err))
		return
	}

	resp := responses.SuccessDTO{Message: "Producto eliminado correctamente"}
	responses.SendSuccess(c, http.StatusOK, resp)
}
