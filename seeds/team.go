package seeds

import (
	"github.com/bxcodec/faker/v3"
)

func (s Seed) TeamSeed() {

	for i := 0; i < 50; i++ {
		stmt, _ := s.db.Prepare(`INSERT INTO team(team_name) VALUES (?)`)
		_, err := stmt.Exec(faker.Name())
		if err != nil {
			panic(err)
		}
	}
}
