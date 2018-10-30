package http_echo

import (
	"context"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/mochadwi/go-article/models"

	articleUcase "github.com/mochadwi/go-article/article"
	"github.com/labstack/echo"

	"gopkg.in/go-playground/validator.v9"
	"time"
)

type HttpArticleHandler struct {
	AUsecase articleUcase.ArticleUsecase
}

func (a *HttpArticleHandler) GetAll(c echo.Context) error {

	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)
	cursor := c.QueryParam("cursor")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	listAr, nextCursor, err := a.AUsecase.GetAll(ctx, cursor, int64(num))

	//reqId := uuid.NewV4().String()
	var response = &models.BaseResponse{
		RequestID: "",
		Now:       time.Now().Unix(),
	}

	if err != nil {
		response.Code = getStatusCode(err)
		response.Message = err.Error()
		response.Data = listAr
		return c.JSON(getStatusCode(err), response)
	}

	if len(*listAr) > 0 {
		response.Message = models.DATA_AVAILABLE_SUCCESS
	} else {
		response.Message = models.DATA_EMPTY_SUCCESS
	}

	response.Code = http.StatusOK
	response.Data = listAr

	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(response.Code, response)
}

func (a *HttpArticleHandler) GetByTitle(c echo.Context) error {

	title := c.Param("title")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	art, err := a.AUsecase.GetByTitle(ctx, title)

	if err != nil {
		return c.JSON(getStatusCode(err), models.BaseResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, art)
}

func (a *HttpArticleHandler) GetByID(c echo.Context) error {

	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	art, err := a.AUsecase.GetByID(ctx, id)

	if err != nil {
		return c.JSON(getStatusCode(err), models.BaseResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, art)
}

func isRequestValid(m *models.Article) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *HttpArticleHandler) Create(c echo.Context) error {
	var article models.Article
	err := c.Bind(&article)

	var response = &models.BaseResponse{
		RequestID: "",
		Now:       time.Now().Unix(),
	}

	if err != nil {
		response.Code = http.StatusUnprocessableEntity
		response.Message = string(err.Error())
		response.Data = article
		return c.JSON(response.Code, response)
	}

	if ok, err := isRequestValid(&article); !ok {
		response.Code = http.StatusBadRequest
		response.Message = string(err.Error())
		response.Data = article
		return c.JSON(response.Code, response)
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ar, err := a.AUsecase.Create(ctx, &article)

	if err != nil {
		response.Code = getStatusCode(err)
		response.Message = string(err.Error())
		response.Data = article
		return c.JSON(response.Code, response)
	}

	response.Code = http.StatusCreated
	response.Message = models.DATA_CREATED_SUCCESS
	response.Data = ar
	return c.JSON(response.Code, response)
}

func (a *HttpArticleHandler) Update(c echo.Context) error {
	var idA, err = strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNoContent, err.Error())
	}

	var id = int64(idA)
	var article = &models.Article{
		ID: id,
	}

	err = c.Bind(&article)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(article); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ar, err := a.AUsecase.Update(ctx, article)

	if err != nil {
		return c.JSON(getStatusCode(err), models.BaseResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, ar)
}

func (a *HttpArticleHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	_, err = a.AUsecase.Delete(ctx, id)

	if err != nil {

		return c.JSON(getStatusCode(err), models.BaseResponse{Message: err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case models.INTERNAL_SERVER_ERROR:

		return http.StatusInternalServerError
	case models.NOT_FOUND_ERROR:
		return http.StatusNotFound
	case models.CONFLIT_ERROR:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func NewArticleHttpHandler(e *echo.Echo, us articleUcase.ArticleUsecase) {
	handler := &HttpArticleHandler{
		AUsecase: us,
	}

	e.GET("/article", handler.GetAll)
	e.POST("/article", handler.Create)
	e.GET("/article/:title", handler.GetByTitle)
	e.PUT("/article/:id", handler.Update)
}
