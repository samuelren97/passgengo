package generator

import (
	"fmt"
	"passgengo/internal/hashing"
	"passgengo/internal/logging"
	"strconv"
	"strings"
)

func GetGeneratorFromWizard() (Generator, error) {
	builder := NewGeneratorBuilder()

	length, err := getUserIntInput(
		fmt.Sprintf("Please enter the password length (min: %d, max: %d): ", MinPassLen, MaxPassLen),
		func(i int) bool { return i >= MinPassLen && i <= MaxPassLen },
	)
	if err != nil {
		return nil, err
	}
	builder.Length(length)

	hexString, err := getYesNoInput("Do you wish to generate a hexadecimal string ?")
	if err != nil {
		return nil, err
	}
	if hexString {
		builder.HexString()
		return builder.Build(), nil
	}

	base64String, err := getYesNoInput("Do you wish to generate a standard base64 string ?")
	if err != nil {
		return nil, err
	}
	if base64String {
		builder.Base64String()
		return builder.Build(), nil
	}

	noUpper, err := getYesNoInput("Do you wish to disable upper case characters ?")
	if err != nil {
		return nil, err
	}
	if noUpper {
		builder.NoUpperChars()
	}

	noSpecial, err := getYesNoInput("Do you wish to disable special characters ?")
	if err != nil {
		return nil, err
	}
	if noSpecial {
		builder.NoSpecialChars()
	}

	hasHashing, err := getYesNoInput("Do you wish to hash your password ?")
	if err != nil {
		return nil, err
	}
	if hasHashing {
		hashingMethods := hashing.GetHashingMethodsString()
		hashingMethod, err := getUserIntInput(
			fmt.Sprintf("Choose a hashing method from the list: \n%s\n", hashingMethods),
			func(i int) bool { return i >= 0 && i < len(hashing.HashMethods) },
		)
		if err != nil {
			return nil, err
		}

		builder.HashingMethod(hashingMethod)
	}

	return builder.Build(), nil
}

func getYesNoInput(message string) (bool, error) {
	message = message + "(y/n): "
	logging.Input(message)

	var input string
	isNotValid := true
	for isNotValid {
		if _, err := fmt.Scan(&input); err != nil {
			return false, err
		}

		input = strings.ToLower(input)
		isNotValid = input != "y" && input != "n"

		if isNotValid {
			logging.Warning("The input is invalid")
			logging.Input(message)
		}
	}

	return input == "y", nil
}

func getUserIntInput(message string, condition func(int) bool) (int, error) {
	logging.Input(message)

	var input string
	isNotValid := true
	num := -1
	for isNotValid {
		var err error
		if _, err := fmt.Scan(&input); err != nil {
			return -1, err
		}

		num, err = strconv.Atoi(input)
		isNotValid = err != nil || !condition(num)

		if isNotValid {
			logging.Warning("The input is not valid, try again")
			logging.Input(message)
		}
	}
	return num, nil
}
