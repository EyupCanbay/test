package internal

import (
	"context"
	"fmt"
	"net/http"
	"tesodev-korpes/CustomerService/internal/types"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

type ServiceInterface interface {
	GetByID(ctx context.Context, id string) (*types.Customer, error)
	Create(ctx context.Context, customer *CustomerServiceRequest) (string, error)
	Update(ctx context.Context, id string, update interface{}) error
	Delete(ctx context.Context, id string) error
}

func NewHandler(e *echo.Echo, service *Service) {
	handler := &Handler{service: service}

	g := e.Group("/customer")
	g.GET("/:id", handler.GetByID)
	g.POST("/", handler.Create)
	g.PUT("/:id", handler.Update)
	g.PATCH("/:id", handler.PartialUpdate)
	g.DELETE("/:id", handler.Delete)
}

func (h *Handler) GetByID(c echo.Context) error {
	id := c.Param("id")
	customer, err := h.service.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	customerResponse := ToCustomerResponse(customer)
	return c.JSON(http.StatusOK, customerResponse)
}

func (h *Handler) Create(c echo.Context) error {
	var customerReq types.CustomerRequestModel
	if err := c.Bind(&customerReq); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Println(customerReq)
	req := CustomerServiceRequest{
		FirstName: customerReq.FirstName,
	}

	customerId, err := h.service.Create(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := map[string]string{
		"message":   "Succeeded",
		"createdId": customerId,
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *Handler) Update(c echo.Context) error {
	id := c.Param("id")
	var update interface{}
	if err := c.Bind(&update); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.service.Update(c.Request().Context(), id, update); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Customer updated successfully")
}

func (h *Handler) PartialUpdate(c echo.Context) error {
	id := c.Param("id")
	var update interface{}
	if err := c.Bind(&update); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.service.Update(c.Request().Context(), id, update); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Customer partially updated successfully")
}

func (h *Handler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Customer deleted successfully")
}
