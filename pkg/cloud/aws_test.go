package cloud

import "testing"

func TestCheckAwsConfig(t *testing.T) {
	err := CheckAwsConfig()
	if err != nil {
		t.Fatal("could not load aws config:", err)
	}
}
