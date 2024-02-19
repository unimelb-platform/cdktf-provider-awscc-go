package nimblestudiostudiocomponent


type NimblestudioStudioComponentConfigurationComputeFarmConfiguration struct {
	// <p>The name of an Active Directory user that is used on ComputeFarm worker             instances.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#active_directory_user NimblestudioStudioComponent#active_directory_user}
	ActiveDirectoryUser *string `field:"optional" json:"activeDirectoryUser" yaml:"activeDirectoryUser"`
	// <p>The endpoint of the ComputeFarm that is accessed by the studio component             resource.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#endpoint NimblestudioStudioComponent#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

