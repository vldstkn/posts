-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Post(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    title VARCHAR(200) NOT NULL,
    content VARCHAR(1000) NOT NULL,
    author_id INT NOT NULL,
    CONSTRAINT fk_author FOREIGN KEY (author_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Post;
-- +goose StatementEnd
