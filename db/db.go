package db

type db struct {
}
type Db interface {
}

func New() Db {
	d := db{}
	return d
}
