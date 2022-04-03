package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitCandidateRoute(db *gorm.DB, route *gin.Engine) {
    /* candidateRepo := candidate.NewCandidateRepository(db)
    candidateService := candidate.NewCandidateService(candidateRepo)
    candidateHandler := handler.NewCandidateHandler(candidateService)

    candidateRoute := route.Group("/candidate")
    {
        candidateRoute.GET("", candidateHandler.CandidateGetAllHandler)
        candidateRoute.POST("", candidateHandler.CandidatePostHandler)
    } */
}
