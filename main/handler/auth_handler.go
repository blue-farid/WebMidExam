package handler

import (
	"github.com/blue-farid/WebMidExam/model"
	r "github.com/blue-farid/WebMidExam/repository"
	"github.com/blue-farid/WebMidExam/request"
	s "github.com/blue-farid/WebMidExam/security"
	"github.com/labstack/echo/v4"
	"net/http"
)

func LoginHandler(c echo.Context) error {
	si := &request.LoginRequest{}

	if err := c.Bind(si); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	u, err := r.FindUser(si.Username)

	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{"message": "can't find the user!"})
	}

	if !s.CheckPasswordHash(si.Password, u.Password) {
		return c.JSON(http.StatusBadRequest,
			map[string]string{"message": "username or password is incorrect!"})
	}

	var t string
	t, err = s.GenerateJWT(u.ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": "can't generate the token!"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": t})

}

func SignupHandler(c echo.Context) error {
	si := &request.SignupRequest{}
	if err := c.Bind(si); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	u := &model.User{}

	u.Password = si.Password
	u.Username = si.Username

	if err := u.CreationValidate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	h, e := s.HashPassword(u.Password)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": e.Error()})
	}

	u.Password = h

	res, uErr := r.SaveUser(u)

	if uErr != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": "can't create the user!"})
	}

	return c.JSON(http.StatusOK, res)
}
