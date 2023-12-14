package handler

import (
	"github.com/blue-farid/WebMidExam/model"
	r "github.com/blue-farid/WebMidExam/repository"
	"github.com/blue-farid/WebMidExam/request"
	s "github.com/blue-farid/WebMidExam/security"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetBasketHandler(c echo.Context) error {
	id, err := exportId(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid basket ID"})
	}

	basket, err := r.GetBasket(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Something went wrong!"})
	}

	if basket.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "The basket not found!"})
	}

	if !basket.CheckOwnership(s.ExtractUserIDFromToken(c)) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "The Basket Is not yours!"})

	}

	return c.JSON(http.StatusOK, basket)
}

func GetBasketsHandler(c echo.Context) error {
	baskets, err := r.GetAllBaskets(s.ExtractUserIDFromToken(c))

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": "can't return all baskets!"})
	}

	return c.JSON(http.StatusOK, baskets)
}

func UpdateBasketHandler(c echo.Context) error {
	var u = &request.UpdateBasket{}
	id, err := exportId(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid basket ID"})
	}

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := u.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	basket, err := r.GetBasket(id)

	if !basket.CheckOwnership(s.ExtractUserIDFromToken(c)) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "The Basket Is not yours!"})

	}

	b := &model.Basket{Data: u.Data, State: u.State, UserID: s.ExtractUserIDFromToken(c)}

	res, uErr := r.UpdateBasket(b, id)

	if uErr != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": "can't update this basket!"})
	}

	return c.JSON(http.StatusOK, res)
}

func CreateBasketHandler(c echo.Context) error {
	cr := &request.CreateBasket{}
	if err := c.Bind(cr); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := cr.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	b := &model.Basket{Data: cr.Data, State: cr.State, UserID: s.ExtractUserIDFromToken(c)}
	res, uErr := r.CreateBasket(b)

	if uErr != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": "can't create the basket!"})
	}

	return c.JSON(http.StatusOK, res)
}

func DeleteHandler(c echo.Context) error {

	u, _ := r.GetUser(s.ExtractUserIDFromToken(c))

	if !u.HasRole("DELETE") {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "you can't delete any basket!"})
	}

	id, err := exportId(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid basket ID"})
	}

	err = r.DeleteBasket(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Basket not found"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "the basket has been deleted!"})
}

func exportId(c echo.Context) (int, error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return -1, err
	}

	return id, nil
}
