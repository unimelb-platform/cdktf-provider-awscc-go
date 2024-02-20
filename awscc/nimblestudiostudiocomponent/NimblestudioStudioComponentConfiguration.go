package nimblestudiostudiocomponent


type NimblestudioStudioComponentConfiguration struct {
	// <p>The configuration for a Microsoft Active Directory (Microsoft AD) studio             resource.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#active_directory_configuration NimblestudioStudioComponent#active_directory_configuration}
	ActiveDirectoryConfiguration *NimblestudioStudioComponentConfigurationActiveDirectoryConfiguration `field:"optional" json:"activeDirectoryConfiguration" yaml:"activeDirectoryConfiguration"`
	// <p>The configuration for a render farm that is associated with a studio resource.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#compute_farm_configuration NimblestudioStudioComponent#compute_farm_configuration}
	ComputeFarmConfiguration *NimblestudioStudioComponentConfigurationComputeFarmConfiguration `field:"optional" json:"computeFarmConfiguration" yaml:"computeFarmConfiguration"`
	// <p>The configuration for a license service that is associated with a studio             resource.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#license_service_configuration NimblestudioStudioComponent#license_service_configuration}
	LicenseServiceConfiguration *NimblestudioStudioComponentConfigurationLicenseServiceConfiguration `field:"optional" json:"licenseServiceConfiguration" yaml:"licenseServiceConfiguration"`
	// <p>The configuration for a shared file storage system that is associated with a studio             resource.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#shared_file_system_configuration NimblestudioStudioComponent#shared_file_system_configuration}
	SharedFileSystemConfiguration *NimblestudioStudioComponentConfigurationSharedFileSystemConfiguration `field:"optional" json:"sharedFileSystemConfiguration" yaml:"sharedFileSystemConfiguration"`
}

