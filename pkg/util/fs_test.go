package util

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"testing"
	"testing/fstest"
	"time"

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

func (s *CopyDirTestSuite) checkEqualDirEntries(src, dest fs.FS) {
	s.NoError(
		fs.WalkDir(
			src, ".",
			func(path string, d fs.DirEntry, err error) error {
				if d.IsDir() {
					_, readDestErr := fs.ReadDir(dest, path)
					s.NoError(readDestErr)
				} else {
					srcContents, readSrcFErr := fs.ReadFile(src, path)
					s.NoError(readSrcFErr)
					destContents, readDestFErr := fs.ReadFile(dest, path)
					s.NoError(readDestFErr)
					s.True(bytes.Equal(srcContents, destContents))
				}
				return nil
			},
		),
	)
}

type CopyDirTestSuite struct {
	suite.Suite
}

func (s *CopyDirTestSuite) TestCopyDirToEmptyDir() {
	src := fstest.MapFS{
		"test1": {
			Data:    []byte("test 1"),
			Mode:    0755,
			ModTime: time.Now(),
		},
		"nested/test2": {
			Data:    []byte("test 2"),
			Mode:    0755,
			ModTime: time.Now(),
		},
	}
	dest := s.T().TempDir()

	err := CopyDirToDir(src, dest)
	s.NoError(err)

	s.checkEqualDirEntries(src, os.DirFS(dest))
}

func (s *CopyDirTestSuite) TestCopyDirToDirOverwrite() {
	src := fstest.MapFS{
		"test1": {
			Data:    []byte("test 1"),
			Mode:    0755,
			ModTime: time.Now(),
		},
		"nested/test2": {
			Data:    []byte("test 2"),
			Mode:    0755,
			ModTime: time.Now(),
		},
	}

	dest := s.T().TempDir()
	s.NoError(os.WriteFile(path.Join(dest, "test1"), []byte("junk data"), 0755))
	s.NoError(os.WriteFile(path.Join(dest, "test3"), []byte("test 3"), 0755))

	err := CopyDirToDir(src, dest)
	s.NoError(err)

	s.checkEqualDirEntries(src, os.DirFS(dest))
}

func TestCopyDirTestSuiteSuite(t *testing.T) {
	suite.Run(t, new(CopyDirTestSuite))
}
