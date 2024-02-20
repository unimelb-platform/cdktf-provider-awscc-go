package nimblestudiostudiocomponent


type NimblestudioStudioComponentInitializationScripts struct {
	// <p>The version number of the protocol that is used by the launch profile.
	//
	// The only valid
	//             version is "2021-03-31".</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#launch_profile_protocol_version NimblestudioStudioComponent#launch_profile_protocol_version}
	LaunchProfileProtocolVersion *string `field:"optional" json:"launchProfileProtocolVersion" yaml:"launchProfileProtocolVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#platform NimblestudioStudioComponent#platform}.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#run_context NimblestudioStudioComponent#run_context}.
	RunContext *string `field:"optional" json:"runContext" yaml:"runContext"`
	// <p>The initialization script.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#script NimblestudioStudioComponent#script}
	Script *string `field:"optional" json:"script" yaml:"script"`
}

