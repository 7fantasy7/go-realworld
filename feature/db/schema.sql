CREATE TABLE feature
(
    id          BIGSERIAL PRIMARY KEY,
    key         text NOT NULL,
    name        text NOT NULL,
    description text
);

CREATE TABLE organisation_feature
(
    organisation_id bigint not null,
    feature_id      bigint not null,
    enabled         bool   not null
);

CREATE UNIQUE INDEX idx_organisation_feature_org_feature
    ON organisation_feature (organisation_id, feature_id);