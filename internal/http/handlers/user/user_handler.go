package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/engrsakib/erp-system/internal/models"
	"github.com/engrsakib/erp-system/internal/utils"
    dto "github.com/engrsakib/erp-system/internal/dto/user"
    userService "github.com/engrsakib/erp-system/internal/services/user"
)

type UserHandler struct {
    DB *gorm.DB
    UserService *userService.UserService
}

func NewUserHandler(db *gorm.DB, userService *userService.UserService) *UserHandler {
    return &UserHandler{DB: db, UserService: userService}
}

// CreateUser godoc
// @Summary      Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body  body      models.User  true  "User payload"
// @Success      201   {object}  models.User
// @Router       /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := models.User{
        Name:  input.Name,
        Email: input.Email,
    }

    if err := h.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, user)
}


// GetUser godoc
// @Summary      Get User by ID with Permissions
// @Description  Get detailed user info including permissions
// @Tags         User
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  utils.APIResponse{data=dto.UserResponse}
// @Failure      404  {object}  utils.APIResponse
// @Router       /api/v1/users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	
	id := c.Param("id")

	userResponse, err := h.UserService.GetUser(id)
	
	if err != nil {
		
		utils.SendError(c, http.StatusNotFound, "User not found", err)
		return
	}


	utils.SendResponse(c, http.StatusOK, "User retrieved successfully", userResponse, nil)
}

// UpdateUser godoc
// @Summary      Update user info (Mobile cannot be updated)
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id    path      int                    true  "User ID"
// @Param        body  body      dto.UpdateUserRequest  true  "User payload"
// @Success      200   {object}  utils.APIResponse{data=dto.UserResponse}
// @Router       /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var req dto.UpdateUserRequest

    
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.SendError(c, http.StatusBadRequest, "Invalid request body", err)
        return
    }

    
    updatedUser, err := h.UserService.UpdateUser(id, req)
    if err != nil {
        utils.SendError(c, http.StatusInternalServerError, "Failed to update user", err)
        return
    }

    
    utils.SendResponse(c, http.StatusOK, "User updated successfully", updatedUser, nil)
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Deletes a user and cascades delete to permissions
// @Tags         User
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  utils.APIResponse
// @Router       /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
    id := c.Param("id")


    if err := h.UserService.DeleteUser(id); err != nil {
        utils.SendError(c, http.StatusInternalServerError, "Failed to delete user", err)
        return
    }

    
    utils.SendResponse(c, http.StatusOK, "User deleted successfully", nil, nil)
}


// GetUsers godoc
// @Summary      Get All Users
// @Description  Get users list with pagination & search
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        page   query     int     false  "Page No"
// @Param        limit  query     int     false  "Limit"
// @Param        search query     string  false  "Search Keyword"
// @Success      200  {object}  utils.APIResponse
// @Router       /api/v1/users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	var query dto.PaginationQuery

	
	if err := c.ShouldBindQuery(&query); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid query parameters", err)
		return
	}

	
	if query.Page <= 0 { query.Page = 1 }
	if query.Limit <= 0 { query.Limit = 10 }

	
	users, meta, err := h.UserService.GetUsers(query)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to fetch users", err)
		return
	}

	
	utils.SendResponse(c, http.StatusOK, "Users retrieved successfully", users, meta)
}