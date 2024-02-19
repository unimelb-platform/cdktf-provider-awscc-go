package mediapackagepackagingconfiguration


type MediapackagePackagingConfigurationMssPackage struct {
	// A list of MSS manifest configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#mss_manifests MediapackagePackagingConfiguration#mss_manifests}
	MssManifests interface{} `field:"required" json:"mssManifests" yaml:"mssManifests"`
	// A CMAF encryption configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#encryption MediapackagePackagingConfiguration#encryption}
	Encryption *MediapackagePackagingConfigurationMssPackageEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments will be rounded to the nearest multiple of the source fragment duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#segment_duration_seconds MediapackagePackagingConfiguration#segment_duration_seconds}
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
}

