package datatype

type Resulter interface {
	Result(operator string, value interface{}) (bool, error)
}
