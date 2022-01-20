-- +goose Up
-- +goose StatementBegin
CREATE TABLE player (
  player_id int(11) NOT NULL AUTO_INCREMENT,
  player_number int(11) NOT NULL,
  player_name varchar(255) NOT NULL,
  team_id int(11) NOT NULL,
  date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (player_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE player;
-- +goose StatementEnd
