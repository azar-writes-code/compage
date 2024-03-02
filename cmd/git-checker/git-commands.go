package git_checker

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/memory"
	log "github.com/sirupsen/logrus"
)

func CheckIfSHACommitSimilar(repoPath, repoURL, branchName string) (bool, error) {
	// Get local commit SHA
	localSHA, err := GetLocalGitCommitSHA(repoPath)
	if err != nil {
		log.Error("Error fetching local Git commit SHA:", err)
		return false, err
	}

	// Get remote commit SHA
	remoteSHA, err := GetRemoteGitCommitSHA(repoURL, branchName)
	if err != nil {
		log.Errorf("Error fetching remote Git commit SHA: %s", err)
		return false, err
	}

	return localSHA == remoteSHA, nil
}

func GetLocalGitCommitSHA(repoPath string) (string, error) {
	log.Debugf("repoPath: %s", repoPath)
	// Open the given repository
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Errorf("Error opening repository: %s", err)
		return "", err
	}

	// Get the HEAD reference
	ref, err := r.Head()
	if err != nil {
		log.Errorf("Error getting HEAD reference: %s", err)
		return "", err
	}

	// Get the commit SHA
	commitSHA := ref.Hash().String()

	return commitSHA, nil
}

func GetRemoteGitCommitSHA(repoURL, branchName string) (string, error) {
	// List remote references
	references, err := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{repoURL},
	}).List(&git.ListOptions{})
	if err != nil {
		log.Errorf("Error listing remote references: %s", err)
		return "", err
	}

	// Look for the branch reference
	for _, ref := range references {
		if ref.Name().String() == "refs/heads/"+branchName {
			return ref.Hash().String(), nil
		}
	}

	return "", fmt.Errorf("branch %s not found", branchName)
}
