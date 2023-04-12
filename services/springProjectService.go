package services

import (
	"os"
	"os/exec"

	"log"

	"github.com/google/uuid"
	"github.com/nepackage/nepackage/models"
	"github.com/nepackage/nepackage/utils"
)

func SpringProjectGenerator(springProject models.SpringProject) (springProjectOut *models.SpringProject, projectPath string, err error) {
	uuidProject := uuid.New().String()
	tmpFolderCreation, err := utils.TmpFolderCreation(uuidProject)
	if err != nil {
		return nil, "", err
	}

	cmd := exec.Command("spring", "init", "--artifactId="+springProject.ArtifactId, "--bootVersion="+springProject.BootVersion, `--description="`+springProject.Description+`"`, "--groupId="+springProject.GroupId, "--javaVersion="+springProject.JavaVersion, "--language="+springProject.Language, "--name="+springProject.Name, "--packageName="+springProject.PackageName, "--packaging="+springProject.Packaging, "--type="+springProject.Type, "--version="+springProject.Version, uuidProject)
	log.Println(cmd)

	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Println("could not run command: ", err)
		return nil, "", err
	}

	log.Println("Project created in: ", tmpFolderCreation)
	return &springProject, tmpFolderCreation, nil
}
