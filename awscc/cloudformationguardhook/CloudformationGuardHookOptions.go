package cloudformationguardhook


type CloudformationGuardHookOptions struct {
	// S3 Source Location for the Guard files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudformation_guard_hook#input_params CloudformationGuardHook#input_params}
	InputParams *CloudformationGuardHookOptionsInputParams `field:"optional" json:"inputParams" yaml:"inputParams"`
}

