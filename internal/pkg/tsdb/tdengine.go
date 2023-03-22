package tsdb

import (
	"database/sql"
	"flush-tsdb/internal/pkg/configs"
	"fmt"
	"sync"

	_ "github.com/taosdata/driver-go/v3/taosSql"
)

var tdEngineOnce sync.Once
var tdEngineClient *TDEngine

type TDEngine struct {
	TxsClient    *sql.DB
	TxsGasClient *sql.DB
	AmountClient *sql.DB
}

func NewTDEngine() *TDEngine {
	tdEngineOnce.Do(func() {
		tdEngineClient = &TDEngine{}
	})
	return tdEngineClient
}

func (t *TDEngine) Init(cfg configs.TSDB) error {
	txsClient, err := sql.Open("taosSql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, "avata_txs"))
	if err != nil {
		return fmt.Errorf("connect TDengine db %s, err: %w", "avata_txs", err)
	}
	t.TxsClient = txsClient

	txsGasClient, err := sql.Open("taosSql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, "avata_txs_gas"))
	if err != nil {
		return fmt.Errorf("connect TDengine db %s, err: %w", "avata_txs_gas", err)
	}
	t.TxsGasClient = txsGasClient

	amountClient, err := sql.Open("taosSql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, "avata_amounts"))
	if err != nil {
		return fmt.Errorf("connect TDengine db %s, err: %w", "avata_amounts", err)
	}
	t.AmountClient = amountClient

	return nil
}
