// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Feature struct {
	ID          int64
	Key         string
	Name        string
	Description pgtype.Text
}

type OrganisationFeature struct {
	OrganisationID int64
	FeatureID      int64
	Enabled        bool
}
