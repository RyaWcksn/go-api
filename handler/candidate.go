package handler

/* type candidateHandler struct {
	candidateService candidate.ServiceCandidate
}

func NewCandidateHandler(service candidate.ServiceCandidate) *candidateHandler {
	return &candidateHandler{service}
}

func (h candidateHandler) CandidatePostHandler(c *gin.Context) {
	var model candidate.Candidate

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	candidate, err := h.candidateService.Create(model)
    fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": string(err.Error()),
		})
        return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":  candidate,
		"error": nil,
	})
}

func (h candidateHandler) CandidateGetAllHandler(c *gin.Context) {
	candidates, err := h.candidateService.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data":  nil,
			"error": err.Error(),
		})
        return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":  candidates,
		"error": nil,
	})
} */
