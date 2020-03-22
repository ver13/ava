package error

// AVA Error Interface
type ErrorI interface {
	Error() error
	Println() string
	String() string
	ToJSON() string
}
