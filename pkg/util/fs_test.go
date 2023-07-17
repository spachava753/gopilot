package util

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

// TODO: get rid of dependency of Terraform artifact store,
//  create a proxy server that will serve artifact

type CreateFolderIfNotExistSuite struct {
	suite.Suite
}

func (s *CreateFolderIfNotExistSuite) TestFolderAlreadyExists() {
	tempDir := s.T().TempDir()
	s.NoError(os.MkdirAll(tempDir, 0755))
	s.NoError(CreateFolderIfNotExist(tempDir))
}

func (s *CreateFolderIfNotExistSuite) TestFolderDoesNotExist() {
	tempDir := s.T().TempDir()
	createDirPath := path.Join(tempDir, "terraform-install")
	s.NoError(CreateFolderIfNotExist(createDirPath))
	s.DirExists(createDirPath)
}

func (s *CreateFolderIfNotExistSuite) TestParentFolderDoesNotExist() {
	tempDir := s.T().TempDir()
	createPath := "./testdata/nested/folder"
	createDirPath := path.Join(tempDir, createPath)
	expectedErrMsg := fmt.Sprintf(
		"unable to get permissions of parent folder: stat %s: no such "+
			"file or directory",
		filepath.Dir(createDirPath),
	)
	s.EqualError(CreateFolderIfNotExist(createDirPath), expectedErrMsg)
}

func TestCreateFolderIfNotExistSuite(t *testing.T) {
	suite.Run(t, new(CreateFolderIfNotExistSuite))
}
