package sagemakerdomain


type SagemakerDomainDefaultUserSettingsCustomPosixUserConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#gid SagemakerDomain#gid}.
	Gid *float64 `field:"required" json:"gid" yaml:"gid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#uid SagemakerDomain#uid}.
	Uid *float64 `field:"required" json:"uid" yaml:"uid"`
}

