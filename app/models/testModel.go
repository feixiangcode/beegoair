package models

import (
    "beegoair/core/model"
    "database/sql"
)
type TestModel struct {
    model.Model
}

func (this *TestModel) Init(db *sql.DB) (*TestModel) {
    this.InitModel(this.TableName(), db)
    return this
}

func (this *TestModel) TableName() string {
    return "test"
}

func NewTestModel(db *sql.DB) (*TestModel) {
    model := new(TestModel)
    model.Init(db)
    return model
}
