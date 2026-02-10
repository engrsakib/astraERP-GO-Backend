package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	
	"github.com/engrsakib/erp-system/internal/http/middlewares"
	faqHandlers "github.com/engrsakib/erp-system/internal/http/handlers/faq"
	faqRepo "github.com/engrsakib/erp-system/internal/repository/faq"
	faqService "github.com/engrsakib/erp-system/internal/services/faq"
)

func RegisterFaqRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	
	repo := faqRepo.NewFaqRepository(db)
	service := faqService.NewFaqService(repo)
	handler := faqHandlers.NewFaqHandler(service)


	faqGroup := rg.Group("/faqs")
	faqGroup.Use(middlewares.JWTAuth()) 
	{
		
		faqGroup.POST("", middlewares.CheckPermission(db, "faq.create"), handler.CreateFaq)
		
	
		faqGroup.POST("/answers", middlewares.CheckPermission(db, "faq.create"), handler.CreateAnswer)
		
		
		// ------------------ Read Routes ------------------
		
	
		faqGroup.GET("", middlewares.CheckPermission(db, "faq.view"), handler.GetFaqs)
		
		
		faqGroup.GET("/:id", middlewares.CheckPermission(db, "faq.view"), handler.GetFaq)
		
		
		// ------------------ Update & Delete Routes ------------------
		
		faqGroup.PUT("/:id", middlewares.CheckPermission(db, "faq.edit"), handler.UpdateFaq)
		
		faqGroup.DELETE("/:id", middlewares.CheckPermission(db, "faq.delete"), handler.DeleteFaq)
	}
}