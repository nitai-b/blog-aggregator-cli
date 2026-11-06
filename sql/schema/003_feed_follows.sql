-- +goose up
CREATE TABLE feed_follows (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL,
  feed_id UUID NOT NULL,
  created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),

  CONSTRAINT fk_feed_follows_user FOREIGN KEY (user_id)
    REFERENCES users(id) ON DELETE CASCADE,

  CONSTRAINT fk_feed_follows_feed FOREIGN KEY (feed_id)
    REFERENCES feeds(id) ON DELETE CASCADE,

  CONSTRAINT feed_follows_user_feed_unique UNIQUE (user_id, feed_id)
);

-- +goose down
DROP TABLE feed_follows;
