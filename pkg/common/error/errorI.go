package error

// AVA Error Interface
type ErrorI interface {
	Error() string
	Println() string
	String() string
	ToJSON() string
}
