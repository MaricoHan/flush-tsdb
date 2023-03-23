package model

type TSData interface {
	DBName() string
	STableName() string
	TableName() string
	InsertSQL() string
}

type Generator interface {
	TSData
	GenerateTestData() TSData
}
