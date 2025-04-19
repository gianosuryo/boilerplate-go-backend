package config

import (
	"errors"
	"net/http"
	"sarana-dafa-ai-service/helper"

	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

var Acl *casbin.Middleware

func InitCasbin() {
	Acl = casbin.New(casbin.Config{
		ModelFilePath: "./config/acl/rbac_model.conf",
		PolicyAdapter: fileadapter.NewAdapter("./config/acl/rbac_policy.csv"),
		Unauthorized: func(c *fiber.Ctx) error {
			return customErrorHandlerProcess(c, errors.New(http.StatusText(http.StatusUnauthorized)), http.StatusUnauthorized)
		},
		Forbidden: func(c *fiber.Ctx) error {
			return customErrorHandlerProcess(c, errors.New(http.StatusText(http.StatusForbidden)), http.StatusForbidden)
		},

		Lookup: func(c *fiber.Ctx) string {
			tokenInfo := helper.GetTokenInfo(c)
			return tokenInfo.UserRole
		},
	})
}
