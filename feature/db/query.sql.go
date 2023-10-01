// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getFeaturesByOrganisation = `-- name: GetFeaturesByOrganisation :many
SELECT f.id, f.key, f.name, f.description, of.enabled
FROM organisation_feature of
         JOIN feature f on of.feature_id = f.id
where of.organisation_id = $1
`

type GetFeaturesByOrganisationRow struct {
	ID          int64
	Key         string
	Name        string
	Description pgtype.Text
	Enabled     bool
}

func (q *Queries) GetFeaturesByOrganisation(ctx context.Context, organisationID int64) ([]GetFeaturesByOrganisationRow, error) {
	rows, err := q.db.Query(ctx, getFeaturesByOrganisation, organisationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeaturesByOrganisationRow
	for rows.Next() {
		var i GetFeaturesByOrganisationRow
		if err := rows.Scan(
			&i.ID,
			&i.Key,
			&i.Name,
			&i.Description,
			&i.Enabled,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrganisationsWithFeature = `-- name: GetOrganisationsWithFeature :many
SELECT of.organisation_id
FROM organisation_feature of
         JOIN feature f on of.feature_id = f.id
where f.key = $1
`

func (q *Queries) GetOrganisationsWithFeature(ctx context.Context, key string) ([]int64, error) {
	rows, err := q.db.Query(ctx, getOrganisationsWithFeature, key)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var organisation_id int64
		if err := rows.Scan(&organisation_id); err != nil {
			return nil, err
		}
		items = append(items, organisation_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const isEnabledByKeyAndOrganisation = `-- name: IsEnabledByKeyAndOrganisation :one
SELECT of.enabled
FROM organisation_feature of
         JOIN feature f on of.feature_id = f.id
where of.organisation_id = $1
  and f.key = $2
`

type IsEnabledByKeyAndOrganisationParams struct {
	OrganisationID int64
	Key            string
}

func (q *Queries) IsEnabledByKeyAndOrganisation(ctx context.Context, arg IsEnabledByKeyAndOrganisationParams) (bool, error) {
	row := q.db.QueryRow(ctx, isEnabledByKeyAndOrganisation, arg.OrganisationID, arg.Key)
	var enabled bool
	err := row.Scan(&enabled)
	return enabled, err
}

const updateOrganisationFeature = `-- name: UpdateOrganisationFeature :exec
INSERT INTO organisation_feature(organisation_id, feature_id, enabled)
VALUES ($1, (select id from feature where key = $2), $3)
ON CONFLICT (organisation_id, feature_id) DO UPDATE SET enabled = $3
`

type UpdateOrganisationFeatureParams struct {
	OrganisationID int64
	Key            string
	Enabled        bool
}

// todo feature key as would be much better :)
func (q *Queries) UpdateOrganisationFeature(ctx context.Context, arg UpdateOrganisationFeatureParams) error {
	_, err := q.db.Exec(ctx, updateOrganisationFeature, arg.OrganisationID, arg.Key, arg.Enabled)
	return err
}