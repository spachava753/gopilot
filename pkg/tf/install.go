package tf

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/hashicorp/go-version"
	install "github.com/hashicorp/hc-install"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/hc-install/src"

	"github.com/spachava753/gopilot/pkg/config"
	"github.com/spachava753/gopilot/pkg/util"
)

// EnsureTerraform ensures that the right version of the terraform cli is
// installed, and returns the exec path
func EnsureTerraform(ctx context.Context, baseDir string) (string, error) {
	info, err := os.Stat(baseDir)
	if err != nil {
		return "", fmt.Errorf("could not stat path: %w", err)
	}

	if !info.IsDir() {
		return "", fmt.Errorf(
			"need to provide a path for base directory: %w",
			err,
		)
	}

	configPath := path.Join(baseDir, config.Path)
	if createDirErr := util.CreateFolderIfNotExist(configPath); createDirErr != nil {
		return "", createDirErr
	}

	installDir := path.Join(configPath, "terraform-install")
	if createDirErr := util.CreateFolderIfNotExist(installDir); createDirErr != nil {
		return "", createDirErr
	}

	exactVersion := &releases.ExactVersion{
		Product:    product.Terraform,
		Version:    version.Must(version.NewVersion(config.TerraformVersion)),
		InstallDir: installDir,
	}

	installer := install.NewInstaller()

	execPath, err := installer.Ensure(
		ctx, []src.Source{
			exactVersion,
		},
	)
	if err != nil {
		return "", fmt.Errorf(
			"error ensuring terraform is installed: %w",
			err,
		)
	}

	return execPath, nil
}
