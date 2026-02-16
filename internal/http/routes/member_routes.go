package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "github.com/engrsakib/erp-system/internal/http/middlewares"
    memberHandlers "github.com/engrsakib/erp-system/internal/http/handlers/group_member"
    memberRepo "github.com/engrsakib/erp-system/internal/repository/group"
    memberService "github.com/engrsakib/erp-system/internal/services/groupe/group_member"
)

func RegisterMemberRoutes(rg *gin.RouterGroup, db *gorm.DB) {

    // Repository
    repo := memberRepo.NewMemberRepository(db)

    // Services
    createService := memberService.NewCreateMemberService(repo)
    getAllService := memberService.NewGetAllMemberService(repo)
    getByIDService := memberService.NewGetMemberByIDService(repo)
    updateService := memberService.NewUpdateMemberService(repo)
    deleteService := memberService.NewDeleteMemberService(repo)

    // Handlers
    createHandler := memberHandlers.CreateMemberHandler(createService)
    getAllHandler := memberHandlers.GetAllMembersHandler(getAllService)
    getByIDHandler := memberHandlers.GetMemberByIDHandler(getByIDService)
    updateHandler := memberHandlers.UpdateMemberHandler(updateService)
    deleteHandler := memberHandlers.DeleteMemberHandler(deleteService)

    memberRoutes := rg.Group("/group-members")
    memberRoutes.Use(middlewares.JWTAuth())
    {
        // Create
        memberRoutes.POST("", middlewares.CheckPermission(db, "member.create"), createHandler)

        // Read
        memberRoutes.GET("", middlewares.CheckPermission(db, "member.view"), getAllHandler)
        memberRoutes.GET("/:id", middlewares.CheckPermission(db, "member.view"), getByIDHandler)

        // Update
        memberRoutes.PUT("/:id", middlewares.CheckPermission(db, "member.edit"), updateHandler)

        // Delete (Soft Delete)
        memberRoutes.DELETE("/:id", middlewares.CheckPermission(db, "member.delete"), deleteHandler)
    }
}
