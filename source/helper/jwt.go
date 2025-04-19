package helper

import (
	"errors"
	"fmt"
	"os"
	"sarana-dafa-ai-service/model/claim"
	"sarana-dafa-ai-service/storage/env"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetTokenInfo(c *fiber.Ctx) (t claim.TokenInfo) {
	t, isOk := c.Locals("token_info").(claim.TokenInfo)
	if !isOk {
		return t
	}
	return t
}
func ExtractJWTHeader(c *fiber.Ctx) error {
	// Default token value
	var tokenValue string

	// Get header Authorization
	headerName := "Authorization"
	authScheme := "Bearer"

	// Get value from selected field
	authValue := c.Get(headerName)
	l := len(authScheme)
	if len(authValue) > l+1 && strings.EqualFold(authValue[:l], authScheme) {
		// Get Bearer value
		tokenValue = strings.TrimSpace(authValue[l:])
	} else {
		return errors.New("token is malformed: token contains an invalid number of segments")
	}

	if tokenValue == "null" {
		return errors.New("token not found")
	}

	// Decrypt bearer value
	decryptedToken, err := DecryptToken(tokenValue)
	if err != nil {
		return err
	}

	// Preparing JWT Token
	var token *jwt.Token

	token, err = jwt.Parse(decryptedToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		hmacSampleSecret := []byte(os.Getenv(env.JWT_SECRET))
		return hmacSampleSecret, nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		t := convertClaimsToTokenInfo(claims)
		c.Locals("token_info", t)
	} else {
		return err
	}

	// Continue request
	return nil
}
func convertClaimsToTokenInfo(claims jwt.MapClaims) (t claim.TokenInfo) {
	// Set UserId
	strUserId := fmt.Sprint(claims["user_id"])
	userId, err := strconv.ParseUint(strUserId, 10, 64)
	if err == nil {
		t.UserId = userId
	}

	// Set String Value
	t.UserName = fmt.Sprint(claims["user_name"])
	t.UserEmail = fmt.Sprint(claims["user_email"])
	t.UserRole = fmt.Sprint(claims["user_role"])

	return t
}
