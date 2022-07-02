package api

import (
	"net/http"
	"unicode"

	"github.com/Hsmnasiri/Torob-sample-core/entity"
	"github.com/Hsmnasiri/Torob-sample-core/utils"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

func PasswordCheck(pass string) bool {
	var (
		upp, low, num bool
		tot           uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || tot < 8 {
		return false
	}

	return true
}

func UsernameCheck(username string) bool {
	_, err := entity.GetUserByUsername(username)

	return err == nil
}

func CurrentUser(c *gin.Context) {

	user_id, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := entity.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := entity.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := entity.LoginCheck(u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func Register(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := entity.User{}
	isPasswordOk := PasswordCheck(input.Password)
	if !isPasswordOk {
		c.JSON(http.StatusBadRequest, gin.H{"message": "password is not safe"})
		return
	}

	isUsernameOk := UsernameCheck(input.Password)
	if !isUsernameOk {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user is already taken"})
		return
	}

	u.Username = input.Username
	u.Password = input.Password
	u.Email = input.Email
	u.Role = input.Role

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}
func GetUsers(c *gin.Context) {

	users, err := entity.GetUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "users find success", "users": users})

}
