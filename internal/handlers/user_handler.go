package handlers

import (
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/requests"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/responses"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/services"
)

type UserHandler struct {
	service services.UserServiceInterface
}

func NewUserHandler(service services.UserServiceInterface) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.Translate(c, "errors.internalServer", nil), err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(users))
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("User not found", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(user))
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userRequest requests.UserCreateRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Internal Server Error", err.Error()))
		return
	}

	var userResponse responses.UserResponse
	if err := h.service.CreateUser(&userRequest, &userResponse); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Internal Server Error", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(userResponse))
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var userRequest requests.UserUpdateRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Internal Server Error", err.Error()))
		return
	}

	var userResponse responses.UserResponse
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.UpdateUser(&userRequest, &userResponse, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Internal Server Error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(userResponse))
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Internal Server Error", err.Error()))
		return
	}
	c.JSON(http.StatusNoContent, utils.SuccessResponse(""))
}
