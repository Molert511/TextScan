package repository

import (
	"fmt"
	"os"

	"github.com/Molert511/TextScan/internal/domain"
	"github.com/otiai10/gosseract/v2"
)

type ocrRepository struct{}

func NewOCRRepository() domain.OCRRepository {
	return &ocrRepository{}
}

func (r *ocrRepository) ExtractText(ocrDomain domain.OCRDomain) (domain.OCRDomain, error) {
	client := gosseract.NewClient()
	defer client.Close()

	_, err := os.Stat(ocrDomain.ImagePath)
	if os.IsNotExist(err) {
		return ocrDomain, fmt.Errorf("file not found: %s", ocrDomain.ImagePath)
	}

	err = client.SetImage(ocrDomain.ImagePath)
	if err != nil {
		return ocrDomain, err
	}

	text, err := client.Text()
	if err != nil {
		return ocrDomain, err
	}

	ocrDomain.Text = text

	return ocrDomain, nil
}
