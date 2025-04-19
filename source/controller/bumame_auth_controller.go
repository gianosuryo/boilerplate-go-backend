package controller

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"sarana-dafa-ai-service/helper"
	"sarana-dafa-ai-service/model/web"
	"sarana-dafa-ai-service/service"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

/**
* Interface design
**/
type BumameAuthController interface {
	// Public
	Login(*fiber.Ctx) error
	ReadToken(*fiber.Ctx) error
}

type RequestLogin struct {
	Email    string `json:"email" form:"email" validate:"required" example:"admin@bumame.com"`
	Password string `json:"password" form:"password" validate:"required" example:"123456"`
}

/**
* Object implementation creation
* And inject depedencies
**/
type BumameAuthControllerImpl struct {
	BumameAuthService service.BumameAuthService
	Validate          *validator.Validate
}

func NewBumameAuthController(service service.BumameAuthService, valid *validator.Validate) BumameAuthController {
	return &BumameAuthControllerImpl{
		BumameAuthService: service,
		Validate:          valid,
	}
}

// @Summary		Login
// @Description	Login
// @Tags		auth
// @Accept		json
// @Produce		json
// @Param		request	body	RequestLogin	true	"Request"
// @Success		200		{object}	web.ResponseStandard
// @Failure		400		{object}	web.ResponseStandard
// @Failure		500		{object}	web.ResponseStandard
// @Router		/auth/login [post]
func (cont *BumameAuthControllerImpl) Login(c *fiber.Ctx) error {
	// Default error slice
	var errs []web.ErrorResponse

	request := RequestLogin{}
	c.BodyParser(&request)

	// Get the basic auth header data
	var basicAuthValue string

	// Get header Authorization
	headerName := "Authorization"
	authScheme := "Basic"

	// Get value from selected field
	authValue := c.Get(headerName)
	l := len(authScheme)
	if len(authValue) > l+1 && strings.EqualFold(authValue[:l], authScheme) {
		// Get Bearer value
		basicAuthValue = strings.TrimSpace(authValue[l:])
	}
	decodedByte, _ := base64.StdEncoding.DecodeString(basicAuthValue)
	decodedString := string(decodedByte)
	// Explode the string
	explodedString := strings.Split(decodedString, ":")
	if len(explodedString) == 2 {
		request.Email = explodedString[0]
		request.Password = explodedString[1]
	}

	// Validate the form data
	err := cont.Validate.Struct(request)
	if err != nil {
		errs = helper.CreateValidationErrorResponse(err)
		return helper.BuildErrorResponse(c, http.StatusBadRequest, errs)
	}

	// Last parameter is for login user flag
	responseToken, err := cont.BumameAuthService.Login(c.Context(), map[string]interface{}{
		"email":    request.Email,
		"password": request.Password,
	})
	if err != nil {
		errs = append(errs, helper.CreateErrorResponse("user", err.Error()))
	}

	if len(errs) > 0 {
		return helper.BuildErrorResponse(c, http.StatusBadRequest, errs)
	}

	return helper.BuildSuccessResponse(c, http.StatusOK, responseToken)
}

// @Summary		Read Token
// @Description	Read Token
// @Tags		auth
// @Accept		json
// @Produce		json
// @Success		200		{object}	web.ResponseStandard
// @Failure		400		{object}	web.ResponseStandard
// @Failure		500		{object}	web.ResponseStandard
// @Router		/auth/read-token [get]
func (cont *BumameAuthControllerImpl) ReadToken(c *fiber.Ctx) error {
	fmt.Println("ReadToken")
	tokenInfo := helper.GetTokenInfo(c)

	return helper.BuildSuccessResponse(c, http.StatusOK, tokenInfo)
}
