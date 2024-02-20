package sagemakerspace


type SagemakerSpaceSpaceSharingSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#sharing_type SagemakerSpace#sharing_type}.
	SharingType *string `field:"required" json:"sharingType" yaml:"sharingType"`
}

