package tokens

const (
	// KIND_READ represents the read kind
	KIND_READ (uint8) = iota

	// KIND_INSERT represents the insert kind
	KIND_INSERT

	// KIND_UPDATE represents the update kind
	KIND_UPDATE

	// KIND_DELETE represents the delete kind
	KIND_DELETE
)

// Tokens represents tokens
type Tokens interface {
	List() []Token
}

// Token represents a token
type Token interface {
	Kind() uint8
	Resource() string
	IsNot() bool
	MustNot() bool
	MustBe() bool
}
