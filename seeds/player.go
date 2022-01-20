package seeds

import (
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func (s Seed) PlayerSeed() {

	for i := 0; i < 50; i++ {
		stmt, _ := s.db.Prepare(`INSERT INTO player(player_name, player_number,team_id) VALUES (?,?,?)`)
		_, err := stmt.Exec(faker.Name(), rand.Intn(100), rand.Intn(20))
		if err != nil {
			panic(err)
		}
	}
}
