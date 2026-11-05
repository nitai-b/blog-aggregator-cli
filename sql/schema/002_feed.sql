-- +goose up
CREATE TABLE feeds (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL,
  name TEXT NOT NULL,
  url TEXT NOT NULL UNIQUE,
  created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),

  CONSTRAINT fk_feeds_user FOREIGN KEY (user_id)
    REFERENCES users(id) ON DELETE CASCADE
);

-- +goose down
DROP TABLE feeds;
