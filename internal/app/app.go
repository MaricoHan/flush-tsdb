package app

import (
	"database/sql"
	"flush-tsdb/internal/pkg/configs"
	"flush-tsdb/internal/pkg/tsdb"
	"fmt"
	"github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

func Server(log *logrus.Entry) {

	if err := tsdb.NewTDEngine().Init(configs.NewConfigs().TSDB); err != nil {
		log.WithError(err).Errorln("init TDEngine connection err")
		return
	}

	dbs := map[string]*sql.DB{
		"txs":     tsdb.NewTDEngine().TxsClient,
		"txs_gas": tsdb.NewTDEngine().TxsGasClient,
		"amounts": tsdb.NewTDEngine().AmountClient,
	}
	for m, db := range dbs {
		if err := flush(m, db); err != nil {
			log.WithError(err).Errorf("flush metric %s err: %s", m, err)
			return
		}
		log.Infof("flush metric %s successfully.", m)
	}
	log.Infof("all flush tasks compeleted!")
}

func flush(m string, db *sql.DB) (err error) {
	_, err = db.Exec("alter table ? add tag version tinyint unsigned;", m)
	if err != nil && !strings.Contains(err.Error(), "Tag already exists") {
		return fmt.Errorf("add tag err: %w", err)
	}

	rows, err := db.Query("show tables;")
	if err != nil {
		return fmt.Errorf("query from tsdb err: %w", err)
	}

	if rows != nil {
		defer func(rows *sql.Rows) {
			if rErr := rows.Close(); rErr != nil {
				err = fmt.Errorf("rows close err: %w", rErr)
			}
		}(rows)

		for rows.Next() {
			var tb string
			if err = rows.Scan(&tb); err != nil {
				return fmt.Errorf("scan err: %w", err)
			}

			// project_id 为 0 的不刷 version
			if match, _ := regexp.MatchString("amounts_\\d+_0_\\d+_\\d+", tb); match {
				continue
			}

			_, err = db.Exec("alter table ?.? set tag version = 1", "avata_"+m, tb)
			if err != nil {
				return fmt.Errorf("set tag err: %w, table: %s", err, tb)
			}
		}
	}

	return nil
}
