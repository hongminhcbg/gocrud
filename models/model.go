package models

type SqlKeyType uint8

const (
	SqlKeyType_NONE SqlKeyType = iota
	SqlKeyType_CANDIDATE
	SqlKeyType_PRIMARY
	SqlKeyType_UNIQUE
)

type SqlHint struct {
	DataType   string
	IsNotNull  bool
	DefaultVal string
	KeyType    SqlKeyType
}
