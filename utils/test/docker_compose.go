package test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gruntwork-io/terratest/modules/shell"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/gruntwork-io/terratest/modules/testing"
	. "github.com/onsi/ginkgo/v2"
)

type DockerComposeSetup struct {
	T        GinkgoTInterface
	DBPort   int
	options  *docker.Options
	useSetup bool
}

func NewDockerComposeSetup() *DockerComposeSetup {
	return &DockerComposeSetup{
		useSetup: _viper.GetBool("use.setup"),
		DBPort:   _viper.GetInt("db.port"),
	}
}

func (s *DockerComposeSetup) Setup() {
	t := WrapGinkgoT(GinkgoT())
	if !s.useSetup {
		return
	}

	_ = os.Setenv("DOCKER_DEFAULT_PLATFORM", "linux/amd64")

	options := &docker.Options{
		WorkingDir:  "../",
		Logger:      nil,
		ProjectName: fmt.Sprintf("unit-test-%s", gofakeit.LetterN(5)),
	}

	envVarsFile := s.configureDockerComposeEnvVars()

	docker.RunDockerCompose(t, options, "--env-file", envVarsFile, "up", "-d")
	dsn := fmt.Sprintf("root:secret@tcp(localhost:%d)/example?charset=utf8mb4&parseTime=True&loc=Local", s.DBPort)
	WaitForMySQL(dsn, 60*time.Second)

	// Give the system enough time to start.
	time.Sleep(15 * time.Second)

	s.options = options
	s.T = t
}

func (s *DockerComposeSetup) Teardown() {
	if !s.useSetup {
		return
	}

	docker.RunDockerCompose(s.T, s.options, "down")
	docker.RunDockerCompose(s.T, s.options, "rm", "-v")

	RemoveDanglingDockerVolumes(s.T)
}

func (s *DockerComposeSetup) configureDockerComposeEnvVars() string {
	envVars := map[string]string{
		"DB_PORT": strconv.Itoa(s.DBPort),
	}

	return s.writeDockerComposeEnvVars(envVars)
}

func (s *DockerComposeSetup) writeDockerComposeEnvVars(envVars map[string]string) string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("os.Getwd failed with error: %s", err.Error()))
	}

	filePath := fmt.Sprintf("%s/.env", pwd)

	f, err := os.Create(filePath)
	if err != nil {
		panic(fmt.Sprintf("os.OpenFile failed with error: %s", err.Error()))
	}

	defer f.Close()

	content := ""
	for k, v := range envVars {
		content += fmt.Sprintf("%s=%s\n", k, v)
	}

	if _, err := f.WriteString(content); err != nil {
		panic(fmt.Sprintf("f.WriteString failed with error: %s", err.Error()))
	}

	return filePath
}

func RemoveDanglingDockerVolumes(t testing.TestingT) {
	output := shell.RunCommandAndGetOutput(t, shell.Command{
		Command: "docker",
		Args:    []string{"volume", "ls", "-qf", "dangling=true"},
	})
	volumes := strings.Split(output, "\n")
	shell.RunCommand(t, shell.Command{
		Command: "docker",
		Args:    append([]string{"volume", "rm"}, volumes...),
	})
}
