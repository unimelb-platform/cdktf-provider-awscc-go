package ecrrepository


type EcrRepositoryImageScanningConfiguration struct {
	// The setting that determines whether images are scanned after being pushed to a repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_repository#scan_on_push EcrRepository#scan_on_push}
	ScanOnPush interface{} `field:"optional" json:"scanOnPush" yaml:"scanOnPush"`
}

