package main

import (
	"compress/gzip"
	"fmt"
	"github.com/docker/docker/api/types"
	dockerCli "github.com/docker/docker/client"
	"golang.org/x/net/context"
	"io"
	"os"
)

type Builder struct {
	docker         *dockerCli.Client
	projectsFolder string
}

func NewBuilder(docker *dockerCli.Client) *Builder {
	return &Builder{
		docker:         docker,
		projectsFolder: "/projects",
	}
}

// BuildFromFile assumes a .tar.gz file
func (b *Builder) BuildFromFile(project string, sourceFile string) error {
	// unzip untar file
	err := b.extractProject(project, sourceFile)
	if err != nil {
		return err
	}

	err = b.buildDockerImage(project)

	return nil
}

func (b *Builder) buildDockerImage(project string) error {
	file, err := os.Open(fmt.Sprintf("%s/%s", b.projectsFolder, project))
	if err != nil {
		return err
	}

	buildOptions := types.ImageBuildOptions{
		Tags:           []string{fmt.Sprintf("%s:v1", project)},
		SuppressOutput: true,
	}

	_, err = b.docker.ImageBuild(context.Background(), file, buildOptions)
	return err
}

func (b *Builder) extractProject(project string, sourceFile string) error {
	target := fmt.Sprintf("%s/%s", b.projectsFolder, project)
	fileReader, err := os.Open(sourceFile)
	if err != nil {
		return err
	}

	gzReader, err := gzip.NewReader(fileReader)
	if err != nil {
		return err
	}

	// TODO:
	// we should be able to send tar reader on the fly
	// but seems is not working because is a gz file
	// So save the file ungz and save it as tar
	f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(0755))
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.Copy(f, gzReader); err != nil {
		return err
	}

	return nil
}
