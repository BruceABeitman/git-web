package utility

import (
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	gitObj "gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

const githubURL = "https://github.com/"

// ResolveTreeFromBranch given a user, repo, & branchname
// Retrieves the most recent commit's worktree for designated branch
func ResolveTreeFromBranch(user, repo, branchName string) (*gitObj.Tree, string, error) {
	path := user + "/" + repo + ".git"
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: githubURL + path,
	})
	if err != nil {
		return nil, "", err
	}
	commit, err := resolveCommitFromBranch(r, branchName)
	if err != nil {
		return nil, "", err
	}
	commitObject, err := r.CommitObject(plumbing.NewHash(commit))
	if err != nil {
		return nil, commit, err
	}
	tree, err := commitObject.Tree()
	if err != nil {
		return nil, commit, err
	}
	return tree, commit, nil
}

// resolveCommitFromBranch given a branchname, returns it's corresponding commit hash
func resolveCommitFromBranch(r *git.Repository, branchName string) (string, error) {
	refs, err := r.References()
	if err != nil {
		return "", err
	}
	commit := ""
	refs.ForEach(func(ref *plumbing.Reference) error {
		// The HEAD is omitted in a `git show-ref` so we ignore the symbolic references
		if ref.Type() == plumbing.SymbolicReference {
			return nil
		}
		if ref.Name().String() == "refs/remotes/origin/"+branchName {
			commit = ref.Hash().String()
		}
		return nil
	})
	if commit == "" {
		return "", err
	}
	return commit, nil
}
