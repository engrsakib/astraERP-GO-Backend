package faq

import (
	"net/http"

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


func (h *FaqHandler) CreateFaq(c *gin.Context) {
	var req dto.CreateFaqRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Data", err)
		return
	}

	userIDFloat, _ := c.Get("userID")
	userID := int64(userIDFloat.(float64))

	res, err := h.Service.CreateFaq(userID, req)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create FAQ", err)
		return
	}
	utils.SendResponse(c, http.StatusCreated, "FAQ Topic created successfully", res, nil)
}


func (h *FaqHandler) CreateAnswer(c *gin.Context) {
	var req dto.CreateAnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Data", err)
		return
	}

	userIDFloat, _ := c.Get("userID")
	userID := int64(userIDFloat.(float64))

	res, err := h.Service.CreateAnswer(userID, req)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to add answer", err)
		return
	}
	utils.SendResponse(c, http.StatusCreated, "Answer added successfully", res, nil)
}


func (h *FaqHandler) GetFaqs(c *gin.Context) {
	var query dto.PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Params", err)
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


func (h *FaqHandler) GetFaq(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Service.GetFaqByID(id)
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "FAQ Not Found", err)
		return
	}
	utils.SendResponse(c, http.StatusOK, "Success", res, nil)
}


func (h *FaqHandler) UpdateFaq(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateFaqRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Data", err)
		return
	}

	res, err := h.Service.UpdateFaq(id, req)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to update FAQ", err)
		return
	}
	utils.SendResponse(c, http.StatusOK, "FAQ Updated", res, nil)
}


func (h *FaqHandler) DeleteFaq(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.DeleteFaq(id); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to delete FAQ", err)
		return
	}
	utils.SendResponse(c, http.StatusOK, "FAQ Deleted", nil, nil)
}