package usecase

import (
	"github.com/Molert511/TextScan/internal/domain"
)

type ocrUsecase struct {
	repo domain.OCRRepository
}

func NewOCRUsecase(repo domain.OCRRepository) domain.OCRUsecase {
	return &ocrUsecase{repo: repo}
}

func (u *ocrUsecase) ExtractText(ocrDomain domain.OCRDomain) (domain.OCRDomain, error) {
	return u.repo.ExtractText(ocrDomain)
}
