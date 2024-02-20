package appsyncresolver


type AppsyncResolverRuntime struct {
	// The name of the runtime to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_resolver#name AppsyncResolver#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The version of the runtime to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_resolver#runtime_version AppsyncResolver#runtime_version}
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
}

