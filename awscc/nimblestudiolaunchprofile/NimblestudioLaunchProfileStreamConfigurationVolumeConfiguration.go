package nimblestudiolaunchprofile


type NimblestudioLaunchProfileStreamConfigurationVolumeConfiguration struct {
	// <p>The number of I/O operations per second for the root volume that is attached to             streaming session.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#iops NimblestudioLaunchProfile#iops}
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// <p>The size of the root volume that is attached to the streaming session.
	//
	// The root volume
	//             size is measured in GiBs.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#size NimblestudioLaunchProfile#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// <p>The throughput to provision for the root volume that is attached to the streaming             session.
	//
	// The throughput is measured in MiB/s.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#throughput NimblestudioLaunchProfile#throughput}
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
}

