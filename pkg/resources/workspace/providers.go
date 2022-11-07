package workspace

type ProviderName string

const (
	GitHub ProviderName = "GITHUB_PUBLIC"
)

type Provider interface {
	GetName() string
}
