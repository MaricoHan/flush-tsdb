package tsdb

import (
	"database/sql"
	"fmt"

	_ "github.com/taosdata/driver-go/v3/taosSql"
)

var metrics = []string{"txs", "txs_gas", "amounts"}

var DBs = map[string]*sql.DB{}

var DSN string

func Init(dsn string) error {
	for i := range metrics {
		db, err := sql.Open("taosSql", dsn+"avata_"+metrics[i])
		if err != nil {
			return fmt.Errorf("connect TDengine db %s, err: %w", "avata_"+metrics[i], err)
		}
		DBs[metrics[i]] = db
	}
	return nil
}
