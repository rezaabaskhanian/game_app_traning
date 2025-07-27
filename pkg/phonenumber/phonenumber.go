package phonenumber

import "strconv"

func IsValid(phonenumber string) bool {
	if len(phonenumber) != 11 {
		return false
	}

	if phonenumber[0:2] != "09" {
		return false
	}

	if _, err := strconv.Atoi(phonenumber[:2]); err != nil {
		return false
	}

	return true
}
