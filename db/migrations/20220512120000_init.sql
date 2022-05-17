-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS verifications
(
  id         uuid                  DEFAULT gen_random_uuid() PRIMARY KEY,
  method     varchar(100) NOT NULL DEFAULT '',
  status     varchar(100) NOT NULL DEFAULT '',
  deeplink   text         NOT NULL DEFAULT '',
  chat_id    bigint,
  created_at timestamptz  NOT NULL DEFAULT now(),
  updated_at timestamptz  NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx__verifications__chat_id__created_at ON verifications (chat_id, created_at DESC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS verifications;
-- +goose StatementEnd
