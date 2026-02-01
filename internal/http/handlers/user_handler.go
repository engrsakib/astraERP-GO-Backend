package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "github.com/engrsakib/erp-system/internal/models"
)

type UserHandler struct {
    DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
    return &UserHandler{DB: db}
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

// GetUsers godoc
// @Summary      List users
// @Tags         users
// @Produce      json
// @Success      200 {array} models.User
// @Router       /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
    var users []models.User
    if err := h.DB.Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
        return
    }
    c.JSON(http.StatusOK, users)
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
