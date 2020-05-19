package util

import "regexp"

func ValidateUserRegister(name string, email string, pwd string, pwd2 string) (errMsg map[string]string ) {
	errMsg = make(map[string]string)
	if name == "" {
		errMsg["name"] = "Name field is required"
	}
	emailErr := validateEmail(email)
	if emailErr != "" {
		errMsg["email"] = emailErr
	}
	if pwd == "" {
		errMsg["password"] = "Password field is required"
	}
	if pwd2 == "" {
		errMsg["password2"] = "Confirm password field is required"
	} else if pwd != "" && pwd2 != pwd  {
		errMsg["password2"] = "Confirm password is not the same as password"
	}
	return
}

func ValidateUserLogin(email string, pwd string) (errMsg map[string]string) {
	errMsg = make(map[string]string)
	emailErr := validateEmail(email)
	if emailErr != "" {
		errMsg["email"] = emailErr
	}
	if pwd == "" {
		errMsg["password"] = "Password field is required"
	}
	return
}

func validateEmail(email string) string {
	if email == "" {
		return "Email field is required"
	} else {
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !re.MatchString(email) {
			return "Email is invalid"
		}
	}
	return ""
}