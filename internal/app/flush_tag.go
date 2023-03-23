package app

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	"github.com/jony-lee/go-progress-bar"
	"github.com/sirupsen/logrus"

	"flush-tsdb/internal/pkg/tsdb"
)

func FlushTag(log *logrus.Entry) {
	if err := tsdb.Init(tsdb.DSN); err != nil {
		log.WithError(err).Errorln("init TDEngine connection err")
		return
	}

	for _, m := range tsdb.Metrics {
		log.Infof("start flush tag for %s ", m)
		if err := flush(m, tsdb.DBs[m]); err != nil {
			log.WithError(err).Errorf("flush metric %s err: %s", m, err)
			return
		}
		log.Infof("flush tag for %s successfully", m)
	}
	log.Infof("all flush tasks compeleted !")
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

	var tbs []string
	var tb string
	if rows != nil {
		defer func(rows *sql.Rows) {

		}(rows)

		for rows.Next() {
			if err = rows.Scan(&tb); err != nil {
				return fmt.Errorf("scan err: %w", err)
			}

			// project_id 为 0 的不刷 version
			if match, _ := regexp.MatchString("amounts_\\d+_0_\\d+_\\d+", tb); match {
				continue
			}

			tbs = append(tbs, tb)
		}
	}

	if rErr := rows.Close(); rErr != nil {
		err = fmt.Errorf("rows close err: %w", rErr)
	}

	if len(tbs) == 0 {
		return nil
	}

	bar := progress.New(int64(len(tbs)), ProcessBarOptions...)
	for i := range tbs {
		_, err = db.Exec("alter table ?.? set tag version = 1", "avata_"+m, tbs[i])
		if err != nil {
			return fmt.Errorf("set tag err: %w, table: %s", err, tbs[i])
		}
		bar.Done(1)
	}
	bar.Finish()

	return nil
}
