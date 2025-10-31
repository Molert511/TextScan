package http

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Molert511/TextScan/internal/domain"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc domain.OCRUsecase
}

func NewHandler(r *gin.Engine, uc domain.OCRUsecase) {
	newHandler := &Handler{uc: uc}

	r.POST("/extract-text", newHandler.extractText)
}

func (h *Handler) extractText(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image is required"})
		return
	}

	uploadDir := "./tmp_uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create upload directory"})
		return
	}

	newOCRDomain := domain.OCRDomain{
		ImagePath: filepath.Join(uploadDir, file.Filename),
		Timestamp: time.Now(),
	}

	if err := c.SaveUploadedFile(file, newOCRDomain.ImagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	result, err := h.uc.ExtractText(newOCRDomain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OCR failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"text": result.Text,
	})
}
