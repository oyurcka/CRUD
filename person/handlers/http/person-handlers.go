package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/oyurcka/CRUD/model/app"
	"github.com/oyurcka/CRUD/model/person"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

type PersonHandler struct {
	PersonLogic person.Logic
}

type ResponseError struct {
	Message string `json:"message"`
}

func NewPersonHandler(e *echo.Echo, pl person.Logic) {
	handler := &PersonHandler{
		PersonLogic: pl,
	}
	e.GET("/persons", handler.Get)
	e.GET("/persons/:id", handler.GetByID)
	e.POST("/persons", handler.Store)
	e.PUT("/persons", handler.Update)
	e.DELETE("/persons/:id", handler.Delete)
}

func (ph *PersonHandler) Get(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listPerson, err := ph.PersonLogic.Get(ctx)

	if err != nil {
		logrus.Error(err)
		return c.JSON(ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listPerson)
}

func (ph *PersonHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusNotFound, err)
	}

	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	per, err := ph.PersonLogic.GetByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return c.JSON(ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, per)
}

func isRequestValid(per *app.Person) (bool, error) {
	validate := validator.New()
	err := validate.Struct(per)
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	return true, nil
}

func (ph *PersonHandler) Store(c echo.Context) error {
	var person app.Person
	err := c.Bind(&person)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&person); !ok {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = ph.PersonLogic.Store(ctx, &person)

	if err != nil {
		logrus.Error(err)
		return c.JSON(ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, person)
}

func (ph *PersonHandler) Update(c echo.Context) error {
	var person app.Person
	err := c.Bind(&person)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&person); !ok {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = ph.PersonLogic.Update(ctx, &person)

	if err != nil {
		logrus.Error(err)
		return c.JSON(ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, person)
}

func (ph *PersonHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusNotFound, err)
	}
	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = ph.PersonLogic.Delete(ctx, id)
	if err != nil {
		logrus.Error(err)
		return c.JSON(ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
