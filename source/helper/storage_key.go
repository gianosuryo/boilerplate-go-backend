package helper

import "fmt"

func StorageKeyAcademyConfig(code string) string {
	return fmt.Sprintf("acconf_%s", code)
}
func StorageKeyPaymentUrlCache(code string) string {
	return fmt.Sprintf("payment_url_cache_%s", code)
}
func StorageKeyCallbackCache(code string) string {
	return fmt.Sprintf("payment_callback_cache_%s", code)
}
func StorageKeyResetPasswordToken(token string) string {
	return fmt.Sprintf("reset_password_token_%s", token)
}
