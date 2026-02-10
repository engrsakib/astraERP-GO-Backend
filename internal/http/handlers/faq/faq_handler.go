package faq

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	dto "github.com/engrsakib/erp-system/internal/dto/faq"
	faqService "github.com/engrsakib/erp-system/internal/services/faq"
	"github.com/engrsakib/erp-system/internal/utils"
)

// 1. Base Handler Struct
type FaqHandler struct {
	Service *faqService.FaqService
}

// 2. Constructor
func NewFaqHandler(service *faqService.FaqService) *FaqHandler {
	return &FaqHandler{Service: service}
}

// ---------------------- Methods ----------------------

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
// @Failure 500 {object} utils.APIResponse
// @Router /faqs [post]
func (h *FaqHandler) CreateFaq(c *gin.Context) {
	var req dto.CreateFaqRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Data Provided", err)
		return
	}

	// Middleware থেকে UserID বের করা
	userIDFloat, exists := c.Get("userID")
	if !exists {
		utils.SendError(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	userID := int64(userIDFloat.(float64))

	res, err := h.Service.CreateFaq(userID, req)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create FAQ topic", err)
		return
	}
	utils.SendResponse(c, http.StatusCreated, "FAQ Topic created successfully", res, nil)
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
// @Failure 500 {object} utils.APIResponse
// @Router /faqs/answers [post]
func (h *FaqHandler) CreateAnswer(c *gin.Context) {
	var req dto.CreateAnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Data Provided", err)
		return
	}

	userIDFloat, exists := c.Get("userID")
	if !exists {
		utils.SendError(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	userID := int64(userIDFloat.(float64))

	res, err := h.Service.CreateAnswer(userID, req)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to add answer", err)
		return
	}
	utils.SendResponse(c, http.StatusCreated, "Answer added successfully", res, nil)
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
		utils.SendError(c, http.StatusBadRequest, "Invalid Query Parameters", err)
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
		utils.SendError(c, http.StatusInternalServerError, "Failed to fetch FAQs", err)
		return
	}
	utils.SendResponse(c, http.StatusOK, "Success", res, meta)
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
		
		utils.SendError(c, http.StatusNotFound, "FAQ Not Found", err)
		return
	}
	utils.SendResponse(c, http.StatusOK, "Success", res, nil)
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
		utils.SendError(c, http.StatusBadRequest, "Invalid Update Data", err)
		return
	}

	res, err := h.Service.UpdateFaq(id, req)
	if err != nil {
		
		if strings.Contains(strings.ToLower(err.Error()), "record not found") {
			utils.SendError(c, http.StatusNotFound, "FAQ Not Found to Update", err)
		} else {
			utils.SendError(c, http.StatusInternalServerError, "Failed to update FAQ", err)
		}
		return
	}
	utils.SendResponse(c, http.StatusOK, "FAQ Updated Successfully", res, nil)
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
			utils.SendError(c, http.StatusNotFound, "FAQ Not Found to Delete", err)
		} else {
			utils.SendError(c, http.StatusInternalServerError, "Failed to delete FAQ", err)
		}
		return
	}
	utils.SendResponse(c, http.StatusOK, "FAQ Deleted Successfully", nil, nil)
}