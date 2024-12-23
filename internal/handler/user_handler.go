package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"moydom_api/internal/domain"
	"moydom_api/internal/service"
	"net/http"
	"os"
	"time"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// CreateUser godoc
// @Summary      Создание нового пользователя
// @Description  Создаёт нового пользователя с учётом введённых данных.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        input  body      domain.AuthInput  true  "Данные для создания пользователя"
// @Success      201    {object}  domain.User
// @Failure      400    {object} map[string]string "Ошибка"
// @Router       /auth/signup [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var input domain.AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user, err := h.userService.CreateUser(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": user})

	}
}

// Login godoc
// @Summary      Авторизация пользователя
// @Description  Позволяет пользователю авторизоваться, используя имя пользователя и пароль.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        input  body      domain.AuthInput  true  "Данные для авторизации"
// @Success      200    {object}   map[string]string "Token"
// @Failure      400    {object}   map[string]string "Ошибка"
// @Router      /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var input domain.AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var (
		userFound domain.User
		err       error
	)
	if userFound, err = h.userService.GetUserByUsername(input.Username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userFound.ID,
		// FIXME: set 10 seconds
		"exp": time.Now().Add(time.Hour * 10).Unix(),
	})
	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GetUserProfile godoc
// @Summary      Получить профиль пользователя
// @Description  Возвращает данные текущего авторизованного пользователя.
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200    {object}  domain.User
// @Failure      401    {object}  map[string]string "Ошибка"
// @Router       /user/profile [get]
func (h *UserHandler) GetUserProfile(c *gin.Context) {
	user, _ := c.Get("currentUser")
	c.JSON(http.StatusOK, gin.H{"user": user})
}
