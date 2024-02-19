package appstreamappblock


type AppstreamAppBlockSetupScriptDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block#executable_path AppstreamAppBlock#executable_path}.
	ExecutablePath *string `field:"required" json:"executablePath" yaml:"executablePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block#script_s3_location AppstreamAppBlock#script_s3_location}.
	ScriptS3Location *AppstreamAppBlockSetupScriptDetailsScriptS3Location `field:"required" json:"scriptS3Location" yaml:"scriptS3Location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block#timeout_in_seconds AppstreamAppBlock#timeout_in_seconds}.
	TimeoutInSeconds *float64 `field:"required" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block#executable_parameters AppstreamAppBlock#executable_parameters}.
	ExecutableParameters *string `field:"optional" json:"executableParameters" yaml:"executableParameters"`
}

