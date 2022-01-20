-- +goose Up
-- +goose StatementBegin
CREATE TABLE team (
  team_id int(11) NOT NULL AUTO_INCREMENT,
  team_name varchar(255) NOT NULL,
  date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (team_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE team;
-- +goose StatementEnd
