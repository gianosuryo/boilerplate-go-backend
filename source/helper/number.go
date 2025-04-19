package helper

import (
	"strconv"

	"github.com/google/uuid"
)

func Uint64ToStr(number uint64) string {
	return strconv.FormatUint(number, 10)
}

func SetNumberToCommaSeparated(number uint64) string {
	stringedNumber := strconv.Itoa(int(number))
	commaSeparatedStringedNumber := ""
	for i := len(stringedNumber) - 1; i >= 0; i-- {
		commaSeparatedStringedNumber = string(stringedNumber[i]) + commaSeparatedStringedNumber
		if (len(stringedNumber)-i)%3 == 0 && i != 0 {
			commaSeparatedStringedNumber = "." + commaSeparatedStringedNumber
		}
	}

	return commaSeparatedStringedNumber
}

func ContainsInt(slice []int, item int) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func ContainsString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func ContainsInterface(slice []interface{}, item interface{}) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func RemoveDuplicateInt(slice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range slice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func RandomString(length int) string {
	return uuid.New().String()
}
