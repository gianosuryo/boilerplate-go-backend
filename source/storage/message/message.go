package message

import "fmt"

const (
	Success        = "success"
	AlreadyExist   = "already exist"
	AlreadyUsed    = "already used"
	AlreadyExpired = "already expired"
	CannotEmpty    = "cannot empty"
	NotExist       = "not exist"
	NotEnough      = "not enough"
	DuplicateEntry = "duplicate entry"
	NotAllow       = "not allow"
	NotAvailable   = "not available"

	// With value error
	mustGreater       = "must greater than"
	invalidValue      = "invalid value"
	maxImageDimention = "maximum image dimention"
)

func MustGreater(field string) string {
	return mustGreater + ":" + field
}
func InvalidValue(value string) string {
	return invalidValue + ":" + value
}
func MaximumImageDimention(maxHeight int, maxWidth int) string {
	return fmt.Sprintf("%s:%dx%d", maxImageDimention, maxWidth, maxHeight)
}
