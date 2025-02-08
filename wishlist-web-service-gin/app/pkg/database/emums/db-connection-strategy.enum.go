package enums

type DbConnectionStrategy int

const (
	None DbConnectionStrategy = iota
	Postgre DbConnectionStrategy = iota
)