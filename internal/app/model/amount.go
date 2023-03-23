package model

import (
	"fmt"
	"math/rand"

	"github.com/golang-module/carbon"
	"github.com/shopspring/decimal"
)

const AmountSuperTable = "amounts"
const AmountsDBName = "avata_amounts"

// AmountData amounts指标
type AmountData struct {
	ChainID   uint64
	UserID    uint64
	ProjectID uint64
	Version   uint8
	Operation int
	Type      int // 1: 资金账户+代金券 2: 代金券
	Timestamp carbon.Carbon
	Amount    decimal.Decimal

	// 自此之前时序库里最新的值
	BaseAmount decimal.Decimal
}

func (t AmountData) DBName() string {
	return AmountsDBName
}

func (t AmountData) STableName() string {
	return AmountSuperTable
}

func (t AmountData) TableName() string {
	// amounts_<user_id>_<project_id>_<operation>_<type>
	if t.Version < 2 {
		return fmt.Sprintf("%s_%d_%d_%d_%d", t.STableName(), t.UserID, t.ProjectID, t.Operation, t.Type)
	}
	// amounts_<user_id>_<project_id>_<version>_<operation>_<type>
	return fmt.Sprintf("%s_%d_%d_%d_%d_%d", t.STableName(), t.UserID, t.ProjectID, t.Version, t.Operation, t.Type)
}

func (t AmountData) InsertSQL() string {
	return fmt.Sprintf(
		"INSERT INTO %s.%s USING %s TAGS(%d, %d, %d, %d, %d, %d) VALUES(%d, %s)",
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
		t.BaseAmount.Add(t.Amount),
	)
}

func (t AmountData) GenerateTestData() TSData {
	return AmountData{
		ChainID:   uint64(rand.Intn(1234121234)),
		UserID:    uint64(rand.Intn(12341241234)),
		ProjectID: uint64(rand.Intn(2342142143)),
		Version:   uint8(rand.Intn(200)),
		Operation: rand.Intn(10),
		Type:      rand.Intn(10),
		Timestamp: carbon.Now(),
		Amount:    decimal.NewFromFloat(rand.Float64()),
	}
}
