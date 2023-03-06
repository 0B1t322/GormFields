package models

import (
	"github.com/0B1t322/GormFields/field"
	"github.com/google/uuid"
)

const ModelTableName = field.TableName("Model")

var (
	ModelIdField    = field.NewTableField(ModelTableName, `"Id"`)
	ModelTitleField = field.NewTableField(ModelTableName, `"Title"`)
)

type Model struct {
	ID    uuid.UUID `gorm:"column:Id;type:uuid;primaryKey"  json:"id"`
	Title string    `gorm:"column:Title;type:text;not null" json:"title"`
}

func (Model) TableName() string {
	return ModelTableName.String()
}
