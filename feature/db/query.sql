-- name: GetFeaturesByOrganisation :many
SELECT f.*, of.enabled
FROM organisation_feature of
         JOIN feature f on of.feature_id = f.id
where of.organisation_id = $1;

-- name: IsEnabledByKeyAndOrganisation :one
SELECT of.enabled
FROM organisation_feature of
         JOIN feature f on of.feature_id = f.id
where of.organisation_id = $1
  and f.key = $2;

-- name: GetOrganisationsWithFeature :many
SELECT of.organisation_id
FROM organisation_feature of
         JOIN feature f on of.feature_id = f.id
where f.key = $1;

-- todo feature key as would be much better :)
-- name: UpdateOrganisationFeature :exec
INSERT INTO organisation_feature(organisation_id, feature_id, enabled)
VALUES ($1, (select id from feature where key = $2), $3)
ON CONFLICT (organisation_id, feature_id) DO UPDATE SET enabled = $3;