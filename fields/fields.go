package fields

type IField interface {
	Name() string
	DataType() string
	Annotation() string
	Comment() string
	NameSnake() string
	GenSql() string
}
