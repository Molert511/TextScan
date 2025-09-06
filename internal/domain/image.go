package domain

import "time"

type Image struct {
	Timestamp time.Time
	FilePath  string
	Size      uint64
	Text      ScannedText
}
