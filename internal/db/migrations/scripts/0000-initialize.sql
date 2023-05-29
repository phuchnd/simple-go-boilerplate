-- +migrate Up

CREATE TYPE product_template_type AS ENUM (
    'variant',
    'none_variant',
    'bundle',
);

CREATE TABLE IF NOT EXISTS product_templates (

);

-- +migrate Down

DROP TABLE IF EXISTS product_templates;

DROP TYPE IF EXISTS product_template_type;