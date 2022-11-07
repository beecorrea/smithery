package workspace

// Defines a repository structure.
type VcsRepository struct {
	Name    string
	options VcsOptions
}

// VCS provider integrator.
type VcsRepositoryProvider interface {
	// Creates a repository using a provider, according to the defined settings.
	CreateRepository(repo VcsRepository) error
}
