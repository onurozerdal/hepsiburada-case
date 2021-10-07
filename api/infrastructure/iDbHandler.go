package infrastructure

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}