package appsyncapi


type AppsyncApiEventConfigDefaultSubscribeAuthModes struct {
	// Security configuration for your AppSync API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_api#auth_type AppsyncApi#auth_type}
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
}

