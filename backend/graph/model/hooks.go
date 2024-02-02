package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (q *Question) BeforeCreate(tx *gorm.DB) (err error) {
	q.ID = uuid.New().String()
	return 
}