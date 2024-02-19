package sagemakerspace


type SagemakerSpaceOwnershipSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#owner_user_profile_name SagemakerSpace#owner_user_profile_name}.
	OwnerUserProfileName *string `field:"required" json:"ownerUserProfileName" yaml:"ownerUserProfileName"`
}

