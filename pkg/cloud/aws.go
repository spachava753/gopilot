package cloud

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
)

// CheckAwsConfig checks that aws credentials are configured in the user's
// machine by using the AWS Go SDK
func CheckAwsConfig() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// Create a new AWS session using the default configuration
	cfg, loadConfigErr := config.LoadDefaultConfig(
		ctx,
		config.WithDefaultRegion("us-east-1"),
	)
	if loadConfigErr != nil {
		return fmt.Errorf("could not load AWS config: %w", loadConfigErr)
	}

	if _, retrieveCreds := cfg.Credentials.Retrieve(ctx); retrieveCreds != nil {
		return errors.New("config for AWS not set properly")
	}

	return nil
}
