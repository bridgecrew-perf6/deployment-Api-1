package controllers

import (
	"deployment/models"
	"deployment/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controllers struct {
	repository repository.RepositoryBook
}

func NewControllersBook(r repository.RepositoryBook) *Controllers {
	return &Controllers{r}
}

func (r *Controllers) GetBooks(c echo.Context) error {
	books, err := r.repository.GetBooks()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"messages": "success get all books",
		"book":     books,
	})
}

func (r *Controllers) GetBookById(c echo.Context) error {
	// var user models.User
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)
	book, err := r.repository.GetBookById(intId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get book",
		"book":     book,
	})
}

func (r *Controllers) CreateBook(c echo.Context) error {
	var book models.Book
	c.Bind(&book)

	var err error
	book, err = r.repository.CreateBook(book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"messages": "success create book",
		"book":     book,
	})
}

func (r *Controllers) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)
	err := r.repository.DeleteBook(intId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"messages": "success deleted book",
	})
}

func (r *Controllers) UpdateBook(c echo.Context) error {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)

	book, err := r.repository.GetBookById(intId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var tmpBook models.Book
	c.Bind(&tmpBook)
	book.Title = tmpBook.Title
	book.Author = tmpBook.Author

	// var bookRes models.Book
	bookRes, err := r.repository.Updatebook(book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"messages": "success update book",
		"book":     bookRes,
	})
}
