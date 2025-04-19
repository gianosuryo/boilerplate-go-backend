package helper

import (
	"net/http"
	"sarana-dafa-ai-service/model/web"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BuildSuccessResponse(c *fiber.Ctx, httpStatus int, data any) error {
	response := BuildJsonResponse(c, httpStatus, 1, data, nil, nil)
	c.Status(httpStatus)
	return c.JSON(response)
}
func BuildSuccessResponsePagination(c *fiber.Ctx, httpStatus int, data any, pag any) error {
	response := BuildJsonResponse(c, httpStatus, 1, data, nil, pag)
	c.Status(httpStatus)
	return c.JSON(response)
}
func BuildErrorResponse(c *fiber.Ctx, httpStatus int, err []web.ErrorResponse) error {
	fulfilled := uint8(0)
	if httpStatus == 404 {
		fulfilled = 1
	}
	response := BuildJsonResponse(c, httpStatus, fulfilled, nil, groupErrorResponse(err), nil)
	c.Status(httpStatus)
	return c.JSON(response)
}

func groupErrorResponse(err []web.ErrorResponse) []web.ErrorResponse {
	// Create map
	errorMap := map[string][]string{}
	for _, v := range err {
		errorMap[v.Title] = append(errorMap[v.Title], v.Messages...)
	}

	// Save to the new instance
	newErrs := []web.ErrorResponse{}
	for index, key := range errorMap {
		messages := []string{}
		messages = append(messages, key...)
		newErrs = append(newErrs, CreateErrorResponse(index, messages...))
	}
	return newErrs
}
func BuildJsonResponse(c *fiber.Ctx, httpStatus int, fulfilled uint8, data any, err []web.ErrorResponse, pag any) web.ResponseStandard {
	response := web.ResponseStandard{
		StatusCode: httpStatus,
		Message:    http.StatusText(httpStatus),
		Fulfilled:  fulfilled,
		Data:       data,
		Errors:     err,
		Pagination: pag,
	}
	return response
}
func CreateErrorResponse(title string, message ...string) web.ErrorResponse {
	errResp := web.ErrorResponse{
		Title:    title,
		Messages: message,
	}
	return errResp
}
func CreateValidationErrorResponse(validatorError error) (sliceErrorResponse []web.ErrorResponse) {
	var message string
	for _, err := range validatorError.(validator.ValidationErrors) {
		message = strings.Trim("is "+err.ActualTag()+" "+err.Param(), " ")
		sliceErrorResponse = append(sliceErrorResponse, CreateErrorResponse(err.Field(), message))
	}
	return sliceErrorResponse
}
