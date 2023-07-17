package tf

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EnsureTerraformTestSuite struct {
	suite.Suite
}

func (s *EnsureTerraformTestSuite) TestValidBaseDir() {
	baseDir := s.T().TempDir()
	execPath, err := EnsureTerraform(context.TODO(), baseDir)
	s.Assert().NoError(err)
	s.Assert().FileExists(execPath)
}

func (s *EnsureTerraformTestSuite) TestInvalidBaseDir() {
	execPath, err := EnsureTerraform(context.TODO(), "")
	s.Assert().EqualError(
		err, "could not stat path: stat : no such file or directory",
	)
	s.Assert().NoFileExists(execPath)
}

func TestEnsureTerraformTestSuite(t *testing.T) {
	suite.Run(t, new(EnsureTerraformTestSuite))
}
