package commands

import (
	"context"
	"fmt"
	"io/fs"
	"path"

	"github.com/mitchellh/go-homedir"

	"github.com/hashicorp/terraform-exec/tfexec"

	"github.com/spachava753/gopilot/pkg/config"
	"github.com/spachava753/gopilot/pkg/tf"
	"github.com/spachava753/gopilot/pkg/util"
)

// Bootstrap creates the gopilot control plane
func Bootstrap(ctx context.Context, tfDir fs.FS) error {
	// 1. Ensure right terraform version is installed
	dir, homedirErr := homedir.Dir()
	if homedirErr != nil {
		return fmt.Errorf(
			"could not fetch home directory: %w",
			homedirErr,
		)
	}

	tfExecPath, ensureErr := tf.EnsureEnvironment(ctx, dir)
	if ensureErr != nil {
		return fmt.Errorf(
			"could not ensure terraform is installed: %w",
			ensureErr,
		)
	}

	workingDir := path.Join(dir, config.Path)
	tf, newTerraformErr := tfexec.NewTerraform(workingDir, tfExecPath)
	if newTerraformErr != nil {
		return fmt.Errorf(
			"could not instantiate terraform client: %w",
			newTerraformErr,
		)
	}

	// 2. Create VM in cloud for control plane
	if copyErr := util.CopyDirToDir(tfDir, workingDir); copyErr != nil {
		return fmt.Errorf("could not copy terraform modules: %w", copyErr)
	}

	if applyErr := tf.Apply(ctx); applyErr != nil {
		return fmt.Errorf("could not apply terraform: %w", applyErr)
	}
	// 3. Use terraform output to install keys
	// 4. Install gopilot through SSH?
	return nil
}
