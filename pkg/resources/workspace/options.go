package workspace

type VcsOptions struct {
	RepoVisibility string // Options: public, private.
	BranchOptions  VcsBranchOptions
}

type VcsBranchOptions struct {
	MergeType         string   // Options: squash, rebase, merge.
	DeleteOnMerge     bool     // Delete branch after it's been merged.
	ProtectedBranches []string // List of branches that can't be deleted.
}
