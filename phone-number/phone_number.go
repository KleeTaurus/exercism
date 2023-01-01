package phonenumber

import "errors"

var (
	ErrWrongLength        = errors.New("Wrong length")
	ErrInvalidDigit       = errors.New("Invalid digit")
	ErrInvalidCountryCode = errors.New("Invalid country code")
)

func Number(phoneNumber string) (string, error) {
	input := []byte(phoneNumber)
	phoneNums := []byte{}
	for _, d := range input {
		if d >= 48 && d <= 57 {
			phoneNums = append(phoneNums, d)
		}
	}
	if len(phoneNums) != 10 && len(phoneNums) != 11 {
		return "", ErrWrongLength
	}

	if len(phoneNums) == 11 {
		// Remove country code
		if phoneNums[0] != 49 {
			return "", ErrInvalidCountryCode
		}
		phoneNums = phoneNums[1:]
	}

	if phoneNums[0] < 50 || phoneNums[3] < 50 {
		return "", ErrInvalidDigit
	}

	return string(phoneNums), nil
}

func AreaCode(phoneNumber string) (string, error) {
	phoneNums, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return phoneNums[:3], nil
}

func Format(phoneNumber string) (string, error) {
	phoneNums, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return "(" + phoneNums[:3] + ") " + phoneNums[3:6] + "-" + phoneNums[6:], nil
}
