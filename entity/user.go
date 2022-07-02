package entity

import (
	"errors"
	"html"
	"strings"

	"github.com/Hsmnasiri/Torob-sample-core/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Password  string    `gorm:"size:255;not null;" json:"password"`
	Email     string    `gorm:"size:255;not null;unique'" json:"email"`
	Name      string    `json:"name"`
	Role      string    `gorm:"default:18;not null" json:"role"`
	Favorites []Product `gorm:"many2many:user_favorites;"`
	Recent    []Product `gorm:"many2many:user_recent;"`
}

func GetUsers() ([]User, error) {
	users := []User{}
	DB.Find(&users)
	return users, nil
}

func GetUserByID(uid uint) (User, error) {

	var u User

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("user not found")
	}

	u.PrepareGive()

	return u, nil

}
func GetUserByUsername(usr string) (User, error) {

	var u User

	if err := DB.First(&u, usr).Error; err != nil {
		return u, errors.New("user not found")
	}

	u.PrepareGive()

	return u, nil

}

func (u *User) PrepareGive() {
	u.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {

	var err error

	u := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(u.ID, u.Role)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}
