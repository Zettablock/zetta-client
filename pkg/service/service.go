package service

import (
	"os"
	"os/exec"

	"github.com/Zettablock/zetta-sdk/pkg/config"
)

type Zetta struct {
	f *config.Config
}

func NewZetta(f *config.Config) *Zetta {
	return &Zetta{
		f: f,
	}
}

func (s *Zetta) CreateRepository(name string, repoType string) error {
	if err := s.runCommand("git", "init", name); err != nil {
		return err
	}

	if repoType == "Model" {
		// Set up Git LFS
		if err := s.runCommand("git", "lfs", "install"); err != nil {
			return err
		}
		if err := s.runCommand("git", "lfs", "track", "*.model"); err != nil {
			return err
		}
	}

	return nil
}

func (s *Zetta) runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
