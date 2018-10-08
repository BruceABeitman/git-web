package core

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// GetSVGs retrieves a list of SVG files from the latest commit at the designated user's repo & branch
func (s ServiceImpl) GetSVGs(user, repo, branchName string) (*model.SVGList, error) {
	// retrieve the work tree for the given branch
	tree, commit, err := utility.ResolveTreeFromBranch(user, repo, branchName)
	if err != nil {
		fmt.Printf("Error interacting with git-go. %v\n", err)
		return nil, ErrResolvingAsset
	}
	// walk the tree and build array of svg files
	filenames := make([]string, 0)
	tree.Files().ForEach(func(f *object.File) error {
		if f.Name[len(f.Name)-4:] == ".svg" {
			filenames = append(filenames, f.Name)
		}
		return nil
	})
	return &model.SVGList{
		Repo: &model.Repo{
			Name: repo,
		},
		Branch: &model.Branch{
			Name:   branchName,
			Commit: commit,
		},
		TotalFiles: len(filenames),
		Filenames:  filenames,
	}, nil
}
