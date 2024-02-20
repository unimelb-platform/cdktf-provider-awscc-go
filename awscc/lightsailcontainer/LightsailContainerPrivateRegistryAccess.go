package lightsailcontainer


type LightsailContainerPrivateRegistryAccess struct {
	// An object to describe a request to activate or deactivate the role that you can use to grant an Amazon Lightsail container service access to Amazon Elastic Container Registry (Amazon ECR) private repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_container#ecr_image_puller_role LightsailContainer#ecr_image_puller_role}
	EcrImagePullerRole *LightsailContainerPrivateRegistryAccessEcrImagePullerRole `field:"optional" json:"ecrImagePullerRole" yaml:"ecrImagePullerRole"`
}

