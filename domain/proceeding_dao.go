package domain

import (
	"gorm.io/gorm"
	"time"
)

type Proceeding struct {
	gorm.Model
	ImportedLetterNum string
	ExportedLetterNum string
	ImportedLetterId  string
	ExportedLetterId  string
	SrvNum            string
	DateTime          time.Time
}
