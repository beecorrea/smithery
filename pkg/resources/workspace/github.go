package workspace

type GitHubVcsProvider struct {
	Name ProviderName
}

func (gh GitHubVcsProvider) GetName() string {
	return string(GitHub)
}

// Creates a repository using a provider, according to the defined settings.
func (gh GitHubVcsProvider) CreateRepository(repo *VcsRepository) error {
	return nil
}
