package aws_test

import (
	"testing"

	"github.com/khulnasoft-lab/driftctl/test"
	"github.com/khulnasoft-lab/driftctl/test/acceptance"
)

func TestAcc_Aws_S3Bucket_BucketInUsEast1(t *testing.T) {
	acceptance.Run(t, acceptance.AccTestCase{
		TerraformVersion: "1.4.6",
		Paths:            []string{"./testdata/acc/aws_s3_bucket"},
		Args:             []string{"scan"},
		Checks: []acceptance.AccCheck{
			{
				Env: map[string]string{
					"AWS_REGION": "us-east-1",
				},
				Check: func(result *test.ScanResult, stdout string, err error) {
					if err != nil {
						t.Fatal(err)
					}
					result.AssertManagedCount(1)
				},
			},
		},
	})
}
