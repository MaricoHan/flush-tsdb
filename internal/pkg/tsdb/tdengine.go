package tsdb

import (
	"database/sql"
	"fmt"

	_ "github.com/taosdata/driver-go/v3/taosRestful"
)

var Metrics = []string{"amounts", "txs", "txs_gas"}

var DBs = map[string]*sql.DB{}
var DSN string

func Init(dsn string) error {
	for i := range Metrics {
		db, err := sql.Open("taosRestful", dsn+"avata_"+Metrics[i])
		if err != nil {
			return fmt.Errorf("connect TDengine db %s, err: %w", "avata_"+Metrics[i], err)
		}
		DBs[Metrics[i]] = db
	}
	return nil
}
