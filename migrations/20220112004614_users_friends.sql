-- +goose Up
-- +goose StatementBegin
CREATE TABLE users_friends(
    source_id INT NOT NULL,
    target_id INT NOT NULL,
    PRIMARY KEY (source_id, target_id),
    CONSTRAINT fk_friends_source_id
      FOREIGN KEY (source_id)
          REFERENCES users(id)
          ON DELETE NO ACTION
          ON UPDATE NO ACTION,
    CONSTRAINT fk_friends_target_id
      FOREIGN KEY (target_id)
          REFERENCES users(id)
          ON DELETE NO ACTION
          ON UPDATE NO ACTION
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users_friends;
-- +goose StatementEnd
