package athenaworkgroup


type AthenaWorkGroupWorkGroupConfigurationUpdatesEngineVersion struct {
	// The engine version requested by the user.
	//
	// Possible values are determined by the output of ListEngineVersions, including Auto. The default is Auto.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#selected_engine_version AthenaWorkGroup#selected_engine_version}
	SelectedEngineVersion *string `field:"optional" json:"selectedEngineVersion" yaml:"selectedEngineVersion"`
}

