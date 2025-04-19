package service

import (
	"context"
	"errors"
	"os"
	"sarana-dafa-ai-service/helper"
	"sarana-dafa-ai-service/model"
	"sarana-dafa-ai-service/model/claim"
	"sarana-dafa-ai-service/model/web"
	"sarana-dafa-ai-service/storage/env"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

/**
* Interface
**/
type BumameAuthService interface {
	// Public
	Login(ctx context.Context, request map[string]interface{}) (web.ResponseToken, error)
}

/**
* Object implementation creation
**/
type BumameAuthServiceImpl struct {
	DB *gorm.DB
}

func NewBumameAuthService(db *gorm.DB) BumameAuthService {
	return &BumameAuthServiceImpl{
		DB: db,
	}
}

func (service BumameAuthServiceImpl) Login(
	ctx context.Context, request map[string]interface{}) (
	rt web.ResponseToken, err error) {
	// Create transaction in this service
	tx := service.DB

	// Select user by email
	findUser := model.Admin{}
	tx = tx.Where("email = ?", request["email"]).Where("is_deleted = 0").Find(&findUser)

	// If not found the return directly
	if findUser.Id == 0 || findUser.Status == -1 {
		return web.ResponseToken{}, errors.New("email or password incorrect")
	}
	// Verify password
	// fmt.Println(helper.PasswordHash(request["password"].(string)))
	match := helper.PasswordVerify(findUser.Password, request["password"].(string))

	// If password not verfied then return email and password incorect
	if !match {
		return web.ResponseToken{}, errors.New("email or password incorrect")
	}
	rt.User = web.ResponseUser{
		Id:       findUser.Id,
		Name:     findUser.Name,
		Role:     findUser.Role,
		Email:    findUser.Email,
		UserName: findUser.UserName,
	}

	// Do JWT Generation
	jwtExpired, _ := strconv.Atoi(os.Getenv(env.JWT_EXPIRED))
	if jwtExpired == 0 {
		jwtExpired = 1
	}
	defaultClaim := claim.JwtCustomClaims{}
	defaultClaim.UserId = findUser.Id
	defaultClaim.UserName = findUser.Name
	defaultClaim.UserUserName = findUser.UserName
	defaultClaim.UserEmail = findUser.Email
	defaultClaim.UserRole = findUser.Role
	defaultClaim.RegisteredClaims = jwt.RegisteredClaims{
		IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * time.Duration(jwtExpired))},
	}

	// Add claim data with JWT Header
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &defaultClaim)

	// Add sign the to the JWT Token
	signedToken, err := rawToken.SignedString([]byte(os.Getenv(env.JWT_SECRET)))
	helper.PanicIfError(err)

	// Encrypt the token
	encryptedToken, err := helper.EncryptToken(signedToken)
	helper.PanicIfError(err)
	rt.Token = encryptedToken

	return rt, nil
}
