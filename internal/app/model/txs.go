package model

import (
	"fmt"
	"math/rand"

	"github.com/golang-module/carbon"
)

const TxsSuperTable = "txs"
const TxsDBName = "avata_txs"

// TxData tx指标
type TxData struct {
	ChainID   uint64
	UserID    uint64
	ProjectID uint64
	Version   uint8
	Operation uint8
	Type      int
	Timestamp carbon.Carbon
	TxNum     uint64
	MsgNum    uint64

	// 自此之前时序库里最新的值
	BaseTxNum  uint64
	BaseMsgNum uint64
}

func (t TxData) DBName() string {
	return TxsDBName
}

func (t TxData) STableName() string {
	return TxsSuperTable
}

func (t TxData) TableName() string {
	// txs_<project_id>_<operation>_<type>
	if t.Version < 2 {
		return fmt.Sprintf("%s_%d_%d_%d", t.STableName(), t.ProjectID, t.Operation, t.Type)
	}
	// txs_<project_id>_<version>_<operation>_<type>
	return fmt.Sprintf("%s_%d_%d_%d_%d", t.STableName(), t.ProjectID, t.Version, t.Operation, t.Type)
}

func (t TxData) InsertSQL() string {
	return fmt.Sprintf(
		"INSERT INTO %s.%s USING %s TAGS(%d, %d, %d, %d, %d, %d) VALUES(%d, %d, %d)",
		t.DBName(),
		t.TableName(),
		t.STableName(),
		t.ChainID,
		t.UserID,
		t.ProjectID,
		t.Operation,
		t.Type,
		t.Version,
		t.Timestamp.TimestampMilli(),
		t.BaseTxNum+t.TxNum,
		t.BaseMsgNum+t.MsgNum,
	)
}

func (t TxData) GenerateTestData() TSData {
	return TxData{
		ChainID:   uint64(rand.Intn(1234121234)),
		UserID:    uint64(rand.Intn(12341241234)),
		ProjectID: uint64(rand.Intn(2342142143)),
		Version:   uint8(rand.Intn(200)),
		Operation: uint8(rand.Intn(10)),
		Type:      rand.Intn(10),
		Timestamp: carbon.Now(),
		TxNum:     1,
		MsgNum:    uint64(rand.Intn(10)),
	}
}
