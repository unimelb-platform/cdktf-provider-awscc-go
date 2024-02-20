package nimblestudiolaunchprofile


type NimblestudioLaunchProfileStreamConfigurationSessionBackup struct {
	// <p>The maximum number of backups that each streaming session created from this launch             profile can have.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#max_backups_to_retain NimblestudioLaunchProfile#max_backups_to_retain}
	MaxBackupsToRetain *float64 `field:"optional" json:"maxBackupsToRetain" yaml:"maxBackupsToRetain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#mode NimblestudioLaunchProfile#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

