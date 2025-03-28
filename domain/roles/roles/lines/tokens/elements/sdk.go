package elements

// Element represents an element
type Element interface {
	IsRole() bool
	Role() string
	IsPermission() bool
	Permission() string
}
