-- +goose Up
CREATE TABLE feeds (
	name TEXT NOT NULL,
	url TEXT NOT NULL UNIQUE,
	user_id uuid NOT NULL,
	FOREIGN KEY(user_id)
	REFERENCES users(id)
	ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;