package services

import (
	"os"
	"os/exec"
	"time"

	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/nepackage/nepackage/utils"
)

func InitLocalRepository(path string) (err error) { // Create a repository model
	os.Chdir(path)
	initRepo := exec.Command("git", "init")
	if err := initRepo.Run(); err != nil {
		log.Println("could not init git local repository: ", err)
		return err
	}
	log.Println(initRepo)
	return nil
}

func FirstCommitLocalRepository(path string) (err error) { // Create a repository model
	systemUsername, err := utils.GetCurrentUser()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	repository, err := git.PlainOpen(path)
	if err != nil {
		log.Println("error opening git repository", err.Error())
		return err
	}
	repositoryWorktree, err := repository.Worktree()
	if err != nil {
		log.Println("error getting repository worktree", err.Error())
		return err
	}

	_, err = repositoryWorktree.Add(".")
	if err != nil {
		log.Println("error adding files to stage area ", err.Error())
	}

	commit, err := repositoryWorktree.Commit("initial commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  systemUsername,
			Email: systemUsername + "@nepackage.org",
			When:  time.Now(),
		},
	})
	if err != nil {
		log.Println("error creating first commit ", err.Error())
	}
	log.Println(commit)
	return nil
}
