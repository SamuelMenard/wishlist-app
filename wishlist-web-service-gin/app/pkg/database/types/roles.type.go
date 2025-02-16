package types

type UserRole string

const (
	Owner UserRole = "owner"
	Contributor UserRole = "contributor"
)