-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
      id INT NOT NULL AUTO_INCREMENT,
      age TINYINT UNSIGNED NOT NULL,
      auth_token_hash VARCHAR(64) NULL,
      email VARCHAR(100) NOT NULL,
      password VARCHAR(100) NOT NULL,
      name VARCHAR(100) NOT NULL,
      surname VARCHAR(100) NOT NULL,
      sex VARCHAR(100) NOT NULL,
      city VARCHAR(100) NULL,
      interests JSON NULL,
      PRIMARY KEY ( id )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
