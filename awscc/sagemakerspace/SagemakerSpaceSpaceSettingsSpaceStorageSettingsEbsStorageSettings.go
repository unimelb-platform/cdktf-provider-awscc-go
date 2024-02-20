package sagemakerspace


type SagemakerSpaceSpaceSettingsSpaceStorageSettingsEbsStorageSettings struct {
	// Size of the Amazon EBS volume in Gb.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#ebs_volume_size_in_gb SagemakerSpace#ebs_volume_size_in_gb}
	EbsVolumeSizeInGb *float64 `field:"required" json:"ebsVolumeSizeInGb" yaml:"ebsVolumeSizeInGb"`
}

