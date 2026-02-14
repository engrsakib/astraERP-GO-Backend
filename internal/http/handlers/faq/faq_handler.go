package faq

import (
	"fmt"
	"net/http"
	"strings"

	dto "github.com/engrsakib/erp-system/internal/dto/faq"
	faqService "github.com/engrsakib/erp-system/internal/services/faq"
	"github.com/engrsakib/erp-system/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type FaqHandler struct {
	Service *faqService.FaqService
}

func NewFaqHandler(service *faqService.FaqService) *FaqHandler {
	return &FaqHandler{Service: service}
}

// CreateFaq godoc
// @Summary Create a new FAQ Topic
// @Description Create a new FAQ headline and photo
// @Tags Faq
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateFaqRequest true "FAQ Data"
// @Success 201 {object} utils.APIResponse{data=dto.FaqResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /faqs [post]
func (h *FaqHandler) CreateFaq(c *gin.Context) {
	var req dto.CreateFaqRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request payload. Please ensure the 'headline' is provided and the JSON format is correct.", err)
		return
	}

	claimsInterface, exists := c.Get("claims")
	if !exists {
		utils.SendError(c, http.StatusUnauthorized, "Unauthorized access. Token claims are missing from the request context.", nil)
		return
	}

	claims, ok := claimsInterface.(jwt.MapClaims)
	if !ok {
		if castedMap, ok := claimsInterface.(map[string]interface{}); ok {
			claims = castedMap
		} else {
			utils.SendError(c, http.StatusInternalServerError, "Internal Server Error. Failed to process token claims format.", nil)
			return
		}
	}

	idInterface, ok := claims["id"]
	if !ok {
		utils.SendError(c, http.StatusUnauthorized, "Unauthorized access. User ID is missing within the token claims.", nil)
		return
	}

	var userID int64
	switch v := idInterface.(type) {
	case float64:
		userID = int64(v)
	case int:
		userID = int64(v)
	case int64:
		userID = v
	default:
		utils.SendError(c, http.StatusInternalServerError, "Internal Server Error. User ID format in token is invalid.", nil)
		return
	}

	res, err := h.Service.CreateFaq(userID, req)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create the FAQ topic in the database. Please try again.", err)
		return
	}
	utils.SendResponse(c, http.StatusCreated, "FAQ topic created successfully.", res, nil)
}

// CreateAnswer godoc
// @Summary Add an Answer to an FAQ
// @Description Add a new question and answer to an existing FAQ topic
// @Tags Faq
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateAnswerRequest true "Answer Data"
// @Success 201 {object} utils.APIResponse{data=dto.FaqAnswerResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /faqs/answers [post]
func (h *FaqHandler) CreateAnswer(c *gin.Context) {
	var req dto.CreateAnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request payload. Please ensure FAQ ID, Question, and Answer are provided correctly.", err)
		return
	}

	claimsInterface, exists := c.Get("claims")
	if !exists {
		utils.SendError(c, http.StatusUnauthorized, "Unauthorized access. Token claims are missing from the request context.", nil)
		return
	}

	claims, ok := claimsInterface.(jwt.MapClaims)
	if !ok {
		if castedMap, ok := claimsInterface.(map[string]interface{}); ok {
			claims = castedMap
		} else {
			utils.SendError(c, http.StatusInternalServerError, "Internal Server Error. Failed to process token claims format.", nil)
			return
		}
	}

	idInterface, ok := claims["id"]
	if !ok {
		utils.SendError(c, http.StatusUnauthorized, "Unauthorized access. User ID is missing within the token claims.", nil)
		return
	}

	var userID int64
	switch v := idInterface.(type) {
	case float64:
		userID = int64(v)
	case int:
		userID = int64(v)
	case int64:
		userID = v
	default:
		utils.SendError(c, http.StatusInternalServerError, "Internal Server Error. User ID format in token is invalid.", nil)
		return
	}

	res, err := h.Service.CreateAnswer(userID, req)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to save the answer. Please verify the FAQ ID and try again.", err)
		return
	}
	utils.SendResponse(c, http.StatusCreated, "FAQ answer added successfully.", res, nil)
}

// GetFaqs godoc
// @Summary Get All FAQs
// @Description Get a paginated list of FAQs (Headlines only)
// @Tags Faq
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param search query string false "Search by headline"
// @Success 200 {object} utils.APIResponse{data=[]dto.FaqResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /faqs [get]
func (h *FaqHandler) GetFaqs(c *gin.Context) {
	var query dto.PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid query parameters provided for pagination or search.", err)
		return
	}

	if query.Page <= 0 {
		query.Page = 1
	}
	if query.Limit <= 0 {
		query.Limit = 10
	}

	res, meta, err := h.Service.GetFaqs(query)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve the list of FAQs from the database.", err)
		return
	}
	utils.SendResponse(c, http.StatusOK, "FAQ list retrieved successfully.", res, meta)
}

// GetFaq godoc
// @Summary Get Single FAQ Details
// @Description Get a specific FAQ and all its answers by ID
// @Tags Faq
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "FAQ ID"
// @Success 200 {object} utils.APIResponse{data=dto.FaqResponse}
// @Failure 404 {object} utils.APIResponse
// @Router /faqs/{id} [get]
func (h *FaqHandler) GetFaq(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Service.GetFaqByID(id)
	if err != nil {
		utils.SendError(c, http.StatusNotFound, fmt.Sprintf("FAQ with ID %s was not found.", id), err)
		return
	}
	utils.SendResponse(c, http.StatusOK, "FAQ details retrieved successfully.", res, nil)
}

// UpdateFaq godoc
// @Summary Update FAQ
// @Description Update FAQ headline or photo
// @Tags Faq
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "FAQ ID"
// @Param request body dto.UpdateFaqRequest true "Update Data"
// @Success 200 {object} utils.APIResponse{data=dto.FaqResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /faqs/{id} [put]
func (h *FaqHandler) UpdateFaq(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateFaqRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request body. Please check the update data format.", err)
		return
	}

	res, err := h.Service.UpdateFaq(id, req)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "record not found") {
			utils.SendError(c, http.StatusNotFound, fmt.Sprintf("Unable to update: FAQ with ID %s not found.", id), err)
		} else {
			utils.SendError(c, http.StatusInternalServerError, "An internal error occurred while trying to update the FAQ.", err)
		}
		return
	}
	utils.SendResponse(c, http.StatusOK, "FAQ updated successfully.", res, nil)
}

// DeleteFaq godoc
// @Summary Delete FAQ
// @Description Delete an FAQ and all associated answers
// @Tags Faq
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "FAQ ID"
// @Success 200 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /faqs/{id} [delete]
func (h *FaqHandler) DeleteFaq(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.DeleteFaq(id); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "record not found") {
			utils.SendError(c, http.StatusNotFound, fmt.Sprintf("Unable to delete: FAQ with ID %s not found.", id), err)
		} else {
			utils.SendError(c, http.StatusInternalServerError, "An internal error occurred while trying to delete the FAQ.", err)
		}
		return
	}
	utils.SendResponse(c, http.StatusOK, "FAQ and associated answers deleted successfully.", nil, nil)
}