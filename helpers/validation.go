package helpers

import (
	"regexp"

	"github.com/iSatVeerSingh/gossip/models"
)

func CreateUserValidation(user *models.User) (map[string]string, bool) {
	userError := make(map[string]string)

	if len(user.FirstName) < 1 {
		userError["firstname"] = "First name is required"
	}

	if len(user.LastName) < 1 {
		userError["lastname"] = "Last name is required"
	}

	if match, _ := regexp.MatchString("^[a-z0-9.]+@[a-z0-9]+.[a-z]+$", user.Email); !match {
		userError["email"] = "Invalid email"
	}
	if len(user.Username) < 4 {
		userError["username"] = "Username length must be atleast 4"
	}
	if len(user.Password) < 8 {
		userError["password"] = "Password length must be atleast 8"
	}

	if len(userError) != 0 {
		return userError, false
	}
	return userError, true
}
