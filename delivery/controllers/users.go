package controllers

import (
	"deployment/config"
	"deployment/delivery/custmiddlewares"
	"deployment/models"
	"deployment/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type StructControllerUser struct {
	repository repository.NewRepositoryUser
}

func InstenceControllerUser(repository repository.NewRepositoryUser) *StructControllerUser {
	return &StructControllerUser{repository}
}

//GET ALL USERS
func (s StructControllerUser) GetUsersController(c echo.Context) error {
	// var users []models.User

	users, err := s.repository.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

//GET USER
func (s StructControllerUser) GetUserController(c echo.Context) error {
	// var user models.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := s.repository.GetUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user",
		"users":    user,
	})
}

//CREATE USER
func (s StructControllerUser) CreateUserController(c echo.Context) error {
	tmpUser := models.User{}
	c.Bind(&tmpUser)

	user, err := s.repository.CreateUser(tmpUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

//DELETE USER
func (s StructControllerUser) DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// var user models.User
	err = s.repository.DeleteUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"messages": "success deleted",
	})

}

//UPDATE USER
func (s StructControllerUser) UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := s.repository.GetUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var tmpUser models.User
	c.Bind((&tmpUser))
	user.Name = tmpUser.Name
	user.Email = tmpUser.Email
	user.Password = tmpUser.Password

	userRes, err := s.repository.UpdateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"messages": "success update",
		"user":     userRes,
	})

}

// LOGIN
func (s StructControllerUser) Login(c echo.Context) error {
	tmpLogin := models.User{}
	c.Bind(&tmpLogin)
	res, err := s.repository.Login(tmpLogin.Email, tmpLogin.Password)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Data tidak ditemukan")
	}
	// fmt.Println(res)
	res.Token, _ = custmiddlewares.CreateToken(int(res.ID), config.JWT_SECRET)
	// if res.Token == "" {
	// 	res.Token, _ = custMiddlewares.CreateToken(int(res.ID), configs.JWT_SECRET)
	// }
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil login",
		"data":    res.Token,
	})
}
