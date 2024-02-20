package lightsailcontainer


type LightsailContainerContainerServiceDeploymentContainersPorts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_container#port LightsailContainer#port}.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_container#protocol LightsailContainer#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

