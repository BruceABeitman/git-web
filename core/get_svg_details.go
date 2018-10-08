package core

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// GetSVGDetails retrieves details of SVG a file from the latest commit at the designated user's repo
func (s ServiceImpl) GetSVGDetails(user, repo, branchName, filename string) (*model.SVGDetail, error) {
	// retrieve the work tree for the given branch
	tree, commit, err := utility.ResolveTreeFromBranch(user, repo, branchName)
	if err != nil {
		fmt.Printf("Error interacting with git-go. %v\n", err)
		return nil, ErrResolvingAsset
	}
	// walk the tree and find the file specified
	colors := make([]model.Color, 0)
	svgDetail := &model.SVGDetail{Colors: colors}
	tree.Files().ForEach(func(f *object.File) error {
		if f.Name == filename {
			svgDetail.Repo = &model.Repo{
				Name: repo,
			}
			svgDetail.Branch = &model.Branch{
				Name:   branchName,
				Commit: commit,
			}
			svgDetail.Filename = filename
			svgDetail.FileSize = f.Size
			svgDetail.Contents, err = f.Contents()
			if err != nil {
				fmt.Printf("Error retrieving file contents. %v\n", err)
				return err
			}
			// retrieve all color meta-data from the file's contents
			svgDetail.Colors = utility.FindAllColors(svgDetail.Contents)
		}
		return nil
	})
	if len(svgDetail.Filename) == 0 {
		return nil, ErrResolvingAsset
	}
	return svgDetail, nil
}
