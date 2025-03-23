package references

// NewBuilder creates a new version instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a reference builder
type Builder interface {
	Create() Builder
	WithPath(path []string) Builder
	WithName(name string) Builder
	WithVersion(version uint) Builder
	Now() (Reference, error)
}

// Reference represents a reference
type Reference interface {
	Path() []string
	Name() string
	Version() uint
}
