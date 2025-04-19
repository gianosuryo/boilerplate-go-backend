package config

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sarana-dafa-ai-service/helper"
	"sarana-dafa-ai-service/model/web"
	"sarana-dafa-ai-service/storage/env"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

// Fiber storage interface
func NewFiberStorage() (sessionStorage *redis.Client) {

	// Use storage from redis
	/////////////////////////
	fmt.Println("----------------------------")
	fmt.Println("Connect redis instance start")
	redisClient, err := GetRedisInstace()
	if err != nil {
		fmt.Println("Connect redis instance error: ", err)
		fmt.Println("----------------------------")
	} else {
		fmt.Println("Connect redis instance success")
		fmt.Println("----------------------------")
	}

	// fiberStorageRedis "github.com/gofiber/storage/redis/v3"
	// dbRedisSession, _ := strconv.Atoi(os.Getenv(env.REDIS_DB_SESSION))
	// redisPort, _ := strconv.Atoi(os.Getenv(env.REDIS_PORT))

	// Create a client with support for TLS
	// cer, _ := helper.LoadCertficateAndKeyFromFile("server-ca.pem")
	// // if err != nil {
	// // 	log.Println(err)
	// // 	return
	// // }
	// tlsCfg := &tls.Config{
	// 	MinVersion:         tls.VersionTLS12,
	// 	InsecureSkipVerify: true,
	// 	Certificates:       []tls.Certificate{*cer},
	// }

	// sessionStorage = fiberStorageRedis.New(fiberStorageRedis.Config{
	// 	Host: os.Getenv(env.REDIS_HOST),
	// 	Port: redisPort,
	// 	// TLSConfig: tlsCfg,
	// 	PoolSize: 10 * runtime.GOMAXPROCS(0),
	// 	Database: dbRedisSession,
	// })

	// Use database for storage
	///////////////////////////
	// dbPort, _ := strconv.Atoi(os.Getenv(env.DB_PORT))
	// sessionStorage = mysql_storage.New(mysql_storage.Config{
	// 	Host:       os.Getenv(env.DB_HOST),
	// 	Port:       dbPort,
	// 	Database:   os.Getenv(env.DB_NAME),
	// 	Username:   os.Getenv(env.DB_USER),
	// 	Password:   os.Getenv(env.DB_PASSWORD),
	// 	Table:      "golang_fiber_storage",
	// 	Reset:      false,
	// 	GCInterval: 10 * time.Second,
	// })

	// Use storage from  memory
	///////////////////////////
	// sessionStorage = memory.New()

	return redisClient
}

// Instance
func NewFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: CustomErrorHandler, // Override default error handler
		BodyLimit:    50 * 1024 * 1024,   // Max upload file 50 MB
	})

	return app
}
func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	return customErrorHandlerProcess(ctx, err, http.StatusInternalServerError)
}
func customErrorHandlerProcess(ctx *fiber.Ctx, err error, defaultCode int) error {
	// Status code defaults to 500
	code := defaultCode
	errString := err.Error()
	errResponse := []web.ErrorResponse{}

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		errString = e.Error()
	}

	// Create and place error response
	if os.Getenv(env.IS_DEBUG) == "1" || code == http.StatusUnauthorized {
		errResponse = append(errResponse, helper.CreateErrorResponse("internal", errString))
	}

	// Log the error
	helper.ErrorLog().WithFields(logrus.Fields{
		// Connection
		"createdAt": ctx.Context().ConnTime().UnixMilli(),
		"method":    ctx.Method(),
		"path":      ctx.Path(),
		"query":     ctx.Queries(),
		// Response status
		"status":     code, // int
		"statusText": http.StatusText(code),
	}).Error(errString)

	return helper.BuildErrorResponse(ctx, code, errResponse)
}

// Recover Middleware
func SetRecover(app *fiber.App) {
	app.Use(recover.New(recover.Config{
		EnableStackTrace: false,
	}))
}

func SetCORS(app *fiber.App) {
	app.Use(cors.New())
}

// JWT Middleware
func SetJwtMiddleware(app *fiber.App) {
	app.Use(jwtMiddlewareProcess)
}
func customJWTErrorHandler(ctx *fiber.Ctx, err error) error {
	return customErrorHandlerProcess(ctx, err, http.StatusUnauthorized)
}
func jwtMiddlewareProcess(c *fiber.Ctx) error {
	err := helper.ExtractJWTHeader(c)
	if err != nil {
		return customJWTErrorHandler(c, err)
	}
	// Continue request
	return c.Next()
}

// ACL Middleware
func SetAclMiddleware(app *fiber.App) {
	app.Use(Acl.RoutePermission())
}

// Access Middleware
func SetAccessLogger(app *fiber.App) {

	app.Use(logger.New(logger.Config{
		Output:        io.Discard,
		DisableColors: true,
		Format:        "${method}|${ip}|${status}|${latency}|${path}|${queryParams}",
		Done: func(c *fiber.Ctx, logString []byte) {
			logSplit := strings.Split(string(logString[:]), "|")
			statusInt, _ := strconv.Atoi(logSplit[2])

			// Extract user information
			tokenInfo := helper.GetTokenInfo(c)

			// Make log with log access structur:
			// /model/collection/log_access.go
			helper.AccessLog().WithFields(logrus.Fields{
				// Connection
				"createdAt": c.Context().ConnTime().UnixMilli(),
				"method":    logSplit[0],
				"path":      logSplit[4],
				"query":     logSplit[5],
				"ip":        logSplit[1],
				// Response status
				"status":     statusInt, // Int
				"statusText": http.StatusText(statusInt),
				// Latency number
				"latencyMs":   uint64(0),   // Uint64
				"latencyText": logSplit[3], //
				// User
				"userId":   tokenInfo.UserId,
				"userName": tokenInfo.UserName,
				"userRole": tokenInfo.UserRole,
			}).Info("in")
		},
	}))
}
