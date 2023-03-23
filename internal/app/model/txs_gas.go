package model

import (
	"fmt"
	"math/rand"

	"github.com/golang-module/carbon"

	"flush-tsdb/internal/pkg/uuid"
)

const TxsGasSuperTable = "txs_gas"
const TxsGasDBName = "avata_txs_gas"

// TxGasData account指标
type TxGasData struct {
	ChainID   int
	UserID    uint64
	ProjectID uint64
	Version   uint8
	Account   string
	Timestamp carbon.Carbon
	Gas       uint64
	Business  uint64

	// 自此之前时序库里最新的值
	BaseGas      uint64
	BaseBusiness uint64
}

func (t TxGasData) DBName() string {
	return TxsGasDBName
}

func (t TxGasData) STableName() string {
	return TxsGasSuperTable
}

func (t TxGasData) TableName() string {
	// txs_gas_<project_id>_<account>
	if t.Version < 2 {
		return fmt.Sprintf("%s_%d_%s", t.STableName(), t.ProjectID, t.Account)
	}
	// txs_gas_<project_id>_<version>_<account>
	return fmt.Sprintf("%s_%d_%d_%s", t.STableName(), t.ProjectID, t.Version, t.Account)
}

func (t TxGasData) InsertSQL() string {
	return fmt.Sprintf(
		"INSERT INTO %s.%s USING %s TAGS(%d, %d, %d, '%s', %d) VALUES(%d,%d,%d)",
		t.DBName(),
		t.TableName(),
		t.STableName(),
		t.ChainID,
		t.UserID,
		t.ProjectID,
		t.Account,
		t.Version,
		GetLocal8UTC(t.Timestamp, 8).TimestampMilli(),
		t.BaseGas+t.Gas,
		t.BaseBusiness+t.Business,
	)
}

func (t TxGasData) GenerateTestData() TSData {
	return TxGasData{
		ChainID:      rand.Intn(1234121234),
		UserID:       uint64(rand.Intn(12341241234)),
		ProjectID:    uint64(rand.Intn(2342142143)),
		Version:      uint8(rand.Intn(200)),
		Account:      uuid.GetUuid("asfdasdf"),
		Timestamp:    carbon.Now(),
		Gas:          0,
		Business:     0,
		BaseGas:      0,
		BaseBusiness: 0,
	}
}

// GetLocal8UTC
// @Description: 将 utc 时间 s 转换为对应 offset 时区当地早上 8 点对应的 utc 时间
// @param s ：原 utc 时间
// @param offset ：时区
// @return carbon.Carbon
func GetLocal8UTC(s carbon.Carbon, offset int) carbon.Carbon {
	return s.AddHours(offset).StartOfDay().AddHours(8).SubHours(offset)
}
