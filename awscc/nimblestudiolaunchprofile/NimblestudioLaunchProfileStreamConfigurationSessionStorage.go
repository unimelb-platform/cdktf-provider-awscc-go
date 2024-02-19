package nimblestudiolaunchprofile


type NimblestudioLaunchProfileStreamConfigurationSessionStorage struct {
	// <p>Allows artists to upload files to their workstations.
	//
	// The only valid option is
	//                 <code>UPLOAD</code>.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#mode NimblestudioLaunchProfile#mode}
	Mode *[]*string `field:"required" json:"mode" yaml:"mode"`
	// <p>The upload storage root location (folder) on streaming workstations where files are             uploaded.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#root NimblestudioLaunchProfile#root}
	Root *NimblestudioLaunchProfileStreamConfigurationSessionStorageRoot `field:"optional" json:"root" yaml:"root"`
}

