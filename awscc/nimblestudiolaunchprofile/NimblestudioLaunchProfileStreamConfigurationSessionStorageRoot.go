package nimblestudiolaunchprofile


type NimblestudioLaunchProfileStreamConfigurationSessionStorageRoot struct {
	// <p>The folder path in Linux workstations where files are uploaded.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#linux NimblestudioLaunchProfile#linux}
	Linux *string `field:"optional" json:"linux" yaml:"linux"`
	// <p>The folder path in Windows workstations where files are uploaded.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#windows NimblestudioLaunchProfile#windows}
	Windows *string `field:"optional" json:"windows" yaml:"windows"`
}

