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
// @Summary      Get user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.User
// @Router       /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
    id := c.Param("id")

    var user models.User
    if err := h.DB.First(&user, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
        return
    }

    c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary      Update user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "User ID"
// @Param        body  body      models.User true  "User payload"
// @Success      200   {object}  models.User
// @Router       /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
    id := c.Param("id")

    var user models.User
    if err := h.DB.First(&user, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
        return
    }

    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user.Name = input.Name
    user.Email = input.Email

    if err := h.DB.Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
        return
    }

    c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary      Delete user
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      204  {string}  string "No Content"
// @Router       /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
    id := c.Param("id")

    if err := h.DB.Delete(&models.User{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
        return
    }

    c.Status(http.StatusNoContent)
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

	// ১. কুয়েরি বাইন্ডিং (page, limit, search)
	if err := c.ShouldBindQuery(&query); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid query parameters", err)
		return
	}

	// ২. ডিফল্ট ভ্যালু চেক
	if query.Page <= 0 { query.Page = 1 }
	if query.Limit <= 0 { query.Limit = 10 }

	// ৩. সার্ভিস কল
	users, meta, err := h.UserService.GetUsers(query)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to fetch users", err)
		return
	}

	// ৪. স্ট্যান্ডার্ড রেসপন্স
	utils.SendResponse(c, http.StatusOK, "Users retrieved successfully", users, meta)
}