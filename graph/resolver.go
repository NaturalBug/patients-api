package graph

import (
	"database/sql"

	"github.com/NaturalBug/patients-api/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	patients []*model.Patient
	orders   []*model.Order
	DB       *sql.DB
}
