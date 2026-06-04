package appsyncfunctionconfiguration


type AppsyncFunctionConfigurationRuntime struct {
	// The name of the runtime to use. Currently, the only allowed value is APPSYNC_JS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_function_configuration#name AppsyncFunctionConfiguration#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The version of the runtime to use. Currently, the only allowed version is 1.0.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_function_configuration#runtime_version AppsyncFunctionConfiguration#runtime_version}
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
}

