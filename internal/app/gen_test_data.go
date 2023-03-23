package app

import (
	"runtime"

	"github.com/jony-lee/go-progress-bar"
	"github.com/sirupsen/logrus"

	"flush-tsdb/internal/app/model"
	"flush-tsdb/internal/pkg/tsdb"
)

func GenerateTestData(n int, log *logrus.Entry) {
	if err := tsdb.Init(tsdb.DSN); err != nil {
		log.WithError(err).Errorln("init TDEngine connection err")
		return
	}

	ms := []model.Generator{
		model.TxGasData{},
		model.TxData{},
		model.AmountData{},
	}
	for i := range ms {
		log.Infof("start generate test data for %s", ms[i].STableName())
		gen(n, ms[i], log)
		log.Infof("generate test data for %s finishlly", ms[i].STableName())
	}
	log.Infof("all generate tasks compeleted !")
}
func gen(n int, g model.Generator, log *logrus.Entry) {
	var err error
	var data model.TSData

	bar := progress.New(int64(n), ProcessBarOptions...)

	for j := 0; j < n; j++ {
		data = g.GenerateTestData()
		_, err = tsdb.DBs[g.STableName()].Exec(data.InsertSQL())
		if err != nil {
			log.WithError(err).Errorf("insert err, sql: %s", data.InsertSQL())
			return
		}

		if j%1e4 == 0 {
			runtime.GC()
		}
		bar.Done(1)
	}

	bar.Finish()
}
