package references

type reference struct {
	path    []string
	name    string
	version uint
}

func createReference(
	path []string,
	name string,
	version uint,
) Reference {
	out := reference{
		path:    path,
		name:    name,
		version: version,
	}

	return &out
}

// Path returns the path
func (obj *reference) Path() []string {
	return obj.path
}

// Name returns the name
func (obj *reference) Name() string {
	return obj.name
}

// Version returns the version
func (obj *reference) Version() uint {
	return obj.version
}
