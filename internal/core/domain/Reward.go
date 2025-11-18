package domain

import (
	"time"
	"gorm.io/gorm"
)

type $model struct {
	gorm.Model
	// Fields will be added in full version
}

func (m *$model) TableName() string { return "${model,,}s" }
