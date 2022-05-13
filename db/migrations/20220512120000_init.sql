-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS verifications
(
  id         uuid                  DEFAULT gen_random_uuid() PRIMARY KEY,
  method     varchar(100) NOT NULL DEFAULT '',
  status     varchar(100) NOT NULL DEFAULT '',
  deeplink   text         NOT NULL DEFAULT '',
  created_at timestamptz  NOT NULL DEFAULT now(),
  updated_at timestamptz  NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS verifications;
-- +goose StatementEnd
