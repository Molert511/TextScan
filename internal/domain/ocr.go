package domain

import (
	"time"
)

type OCRDomain struct {
	Timestamp time.Time
	ImagePath string
	Text      string
}

type OCRRepository interface {
	ExtractText(ocrDomain OCRDomain) (OCRDomain, error)
}

type OCRUsecase interface {
	ExtractText(ocrDomain OCRDomain) (OCRDomain, error)
}
