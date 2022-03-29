package handler

import (
	"net/http"

	"github.com/RyaWcksn/go-api/candidate"
	"github.com/gin-gonic/gin"
)

type candidateHandler struct {
    candidateService candidate.ServiceCandidate
}

func NewCandidateHandler(service candidate.ServiceCandidate) *candidateHandler{
    return &candidateHandler{service}
}

func (h candidateHandler) CandidatePostHandler(c *gin.Context) {
    var model candidate.Candidate

    if err := c.ShouldBindJSON(&model); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    candidate, err := h.candidateService.Create(model)
    if err != nil {
        panic(err)
    }
    c.JSON(http.StatusOK, gin.H{
        "data": candidate,
    })
}
