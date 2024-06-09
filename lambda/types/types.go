package types

import (
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

func NewUser(registerUser RegisterUser) (User, error) {
	// the higher you go, the more security you get
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	return User{
		Username:     registerUser.Username,
		PasswordHash: string(hashedPassword),
	}, nil
}

func ValidatePassword(hashedPassword string, planTextPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(planTextPassword))
	return err == nil
}
