package http

import (
	"assignment-dua-fga/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	OrderUsecase models.OrderUsecase
}

func (h OrderHandler) CreateOrder(context *gin.Context) {
	var order models.Order
	if err := context.ShouldBindJSON(&order); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err := h.OrderUsecase.Store(order)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "berhasil membuat order baru",
	})
	return
}

func (h OrderHandler) GetOrders(context *gin.Context) {
	orders, err := h.OrderUsecase.Fetch()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, orders)
}

func (h OrderHandler) UpdateById(context *gin.Context) {
	var order models.Order
	err := context.ShouldBindJSON(&order)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	idParam := context.Param("orderId")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	order.ID = id
	err = h.OrderUsecase.Update(id, order)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "berhasil update order",
	})
}

func (h OrderHandler) DeleteById(context *gin.Context) {
	idParam := context.Param("orderId")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	err = h.OrderUsecase.Delete(id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "berhasil menghapus order",
	})
}

func NewOrderHandler(r *gin.Engine, us models.OrderUsecase) {
	handler := &OrderHandler{OrderUsecase: us}

	r.POST("/orders", handler.CreateOrder)
	r.GET("/orders", handler.GetOrders)
	r.PUT("/orders/:orderId", handler.UpdateById)
	r.DELETE("/orders/:orderId", handler.DeleteById)

}
