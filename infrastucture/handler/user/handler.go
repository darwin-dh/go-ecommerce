package user

import (
	"net/http"

	user "ecommer.com/ecommerce/domain/User"
	"ecommer.com/ecommerce/model"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase user.UserCase
}

func newHandler(uc user.UserCase) handler {
	return handler{useCase: uc}
}

func (h handler) Create(c echo.Context) error {
	m := new(model.User)
	if err := c.Bind(&m); err != nil {
		if err := h.useCase.Create(m); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
	}
	return c.JSON(http.StatusCreated, m)

}

func (h handler) GetAll(c echo.Context) error {
	users, err := h.useCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, users)

}
