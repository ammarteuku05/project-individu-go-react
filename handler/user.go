package handler

import (
	"net/http"
	"projectpenyewaanlapangan/entity"
	"projectpenyewaanlapangan/helper"
	"projectpenyewaanlapangan/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

//showUserHandler for handling show all user in db from route "/users"
func (h *userHandler) ShowUserHandler(c *gin.Context) {
	users, err := h.userService.GetAllUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			`message`: `error in internal server error`,
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUserHandler for handing if user / external create new user from route "/users"
func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var inputUser entity.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newUser, err := h.userService.SaveNewUser(inputUser)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create new User", 201, "status Created", newUser)
	c.JSON(201, response)
}

// get user by 1
// 1. get by id sesuai dengan paramter yg dikasih (repository)
// 2. service akan menampikan hasil user by id dengan format yang sudah ditentukan
// 3. handler kita tangkap id dengan c.Param kemudian kita kirim ke service, terus kita tangkap responsenya

func (h *userHandler) GetUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		responseError := helper.APIResponse("error bad request user ID", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get user by ID", 200, "success", user)
	c.JSON(200, response)
}
