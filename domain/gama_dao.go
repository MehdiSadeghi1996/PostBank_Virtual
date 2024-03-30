package domain

import (
	"gorm.io/gorm"
	"time"
)

type Gama struct {
	gorm.Model
	SrvNum      string
	Description string
	DateTime    time.Time
}
