package controller

import (
	"codifin-challenge/domain/model"
	"codifin-challenge/domain/service"
	"codifin-challenge/domain/utils"
	"codifin-challenge/infrastructure/web/dto"
	"codifin-challenge/infrastructure/web/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ShoppingCartController struct {
	shoppingCartService service.ShoppingCartService
}

func NewShoppingCartController(service service.ShoppingCartService) *ShoppingCartController {
	return &ShoppingCartController{shoppingCartService: service}
}

// NewCart
// @Summary Create a new shopping cart
// @Description Creates a new shopping cart with the specified items
// @Tags Shopping Carts
// @ID new-cart
// @Accept json
// @Produce json
// @Param items body []dto.ItemData true "Items to add to the shopping cart"
// @Success 201 {object} dto.ShoppingCartDTO "Created shopping cart"
// @Failure 400 {object} responses.ErrorDTO "Invalid item data"
// @Failure 500 {object} responses.ErrorDTO "Failed to create shopping cart"
// @Router /carts [post]
func (ctrl *ShoppingCartController) NewCart(c *gin.Context) {
	var items []*dto.ItemData
	if err := c.BindJSON(&items); err != nil {
		responses.SendError(c, utils.ToUserError(http.StatusBadRequest, "Datos de carrito incorrectos", err))
		return
	}

	itemsCart := dto.ToItemsCart(0, items)
	newCart := &model.ShoppingCart{
		Items: itemsCart,
	}

	err := ctrl.shoppingCartService.CreateShoppingCart(newCart)
	if err != nil {
		responses.SendError(c, utils.GetCustomError(err))
		return
	}

	created, err := ctrl.shoppingCartService.FindCart(newCart.ID)
	if err != nil {
		responses.SendError(c, utils.GetCustomError(err))
		return
	}

	cartDTO := dto.ToShoppingCartDTO(created)
	responses.SendSuccess(c, http.StatusCreated, cartDTO)
}

// FindShoppingCart
// @Summary Get a shopping cart by ID
// @Description Retrieves a shopping cart by its ID
// @Tags Shopping Carts
// @ID find-shopping-cart
// @Accept json
// @Produce json
// @Param id path int true "Shopping cart ID"
// @Success 200 {object} dto.ShoppingCartDTO "Found shopping cart"
// @Failure 400 {object} responses.ErrorDTO "Invalid shopping cart ID"
// @Failure 500 {object} responses.ErrorDTO "Failed to retrieve shopping cart"
// @Router /cart/{id} [get]
func (ctrl *ShoppingCartController) FindShoppingCart(c *gin.Context) {
	cartID, _ := strconv.Atoi(c.Param("id"))

	cart, err := ctrl.shoppingCartService.FindCart(uint(cartID))
	if err != nil {
		responses.SendError(c, utils.GetCustomError(err))
		return
	}

	cartDTO := dto.ToShoppingCartDTO(cart)
	responses.SendSuccess(c, http.StatusOK, cartDTO)
}

// AddItem
// @Summary Add an item to a shopping cart
// @Description Adds an item to the specified shopping cart
// @Tags Shopping Carts
// @ID add-item
// @Accept json
// @Produce json
// @Param id path int true "Shopping cart ID"
// @Param item body dto.ItemData true "Item data"
// @Success 200 {object} dto.ShoppingCartDTO "Updated shopping cart"
// @Failure 400 {object} responses.ErrorDTO "Invalid shopping cart ID or item data"
// @Failure 500 {object} responses.ErrorDTO "Failed to add item to shopping cart"
// @Router /cart/{id}/items [post]
func (ctrl *ShoppingCartController) AddItem(c *gin.Context) {
	cartID, _ := strconv.Atoi(c.Param("id"))

	var item dto.ItemData
	if err := c.BindJSON(&item); err != nil {
		responses.SendError(c, utils.ToUserError(http.StatusBadRequest, "Datos de producto incorrectos", err))
		return
	}

	itemsCart := dto.ToItemsCart(uint(cartID), []*dto.ItemData{&item})
	err := ctrl.shoppingCartService.AddItemsToShoppingCart(itemsCart)
	if err != nil {
		responses.SendError(c, utils.GetCustomError(err))
		return
	}

	cart, err := ctrl.shoppingCartService.FindCart(uint(cartID))
	if err != nil {
		responses.SendError(c, utils.GetCustomError(err))
		return
	}

	cartDTO := dto.ToShoppingCartDTO(cart)
	responses.SendSuccess(c, http.StatusOK, cartDTO)
}

// RemoveItems
// @Summary Remove items from a shopping cart
// @Description Removes the specified items from the shopping cart
// @Tags Shopping Carts
// @ID remove-items
// @Accept json
// @Produce json
// @Param id path int true "Shopping cart ID"
// @Param productIds body []int true "IDs of the products to remove"
// @Success 200 {object} dto.ShoppingCartDTO "Updated shopping cart"
// @Failure 400 {object} responses.ErrorDTO "Invalid shopping cart ID or item IDs"
// @Failure 500 {object} responses.ErrorDTO "Failed to remove items from shopping cart"
// @Router /cart/{id}/items [delete]
func (ctrl *ShoppingCartController) RemoveItems(c *gin.Context) {
	cartID, _ := strconv.Atoi(c.Param("id"))

	var productIds []uint
	if err := c.BindJSON(&productIds); err != nil {
		responses.SendError(c, utils.ToUserError(http.StatusBadRequest, "Id de producto incorrecto", err))
		return
	}

	err := ctrl.shoppingCartService.RemoveItemsFromShoppingCart(uint(cartID), productIds)
	if err != nil {
		responses.SendError(c, utils.GetCustomError(err))
		return
	}

	cart, err := ctrl.shoppingCartService.FindCart(uint(cartID))
	if err != nil {
		responses.SendError(c, utils.GetCustomError(err))
		return
	}

	cartDTO := dto.ToShoppingCartDTO(cart)
	responses.SendSuccess(c, http.StatusOK, cartDTO)
}
