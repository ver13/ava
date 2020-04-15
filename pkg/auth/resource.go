package auth

// Resource is an entity such as a user or
type Resource struct {
	// Name of the resource
	Name string
	// Type of resource, e.g.
	Type string
	// Endpoint resource e.g NotesService.Create
	Endpoint string
}
