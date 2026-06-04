package nimblestudiolaunchprofile


type NimblestudioLaunchProfileStreamConfigurationVolumeConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_launch_profile#iops NimblestudioLaunchProfile#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_launch_profile#size NimblestudioLaunchProfile#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_launch_profile#throughput NimblestudioLaunchProfile#throughput}.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
}

