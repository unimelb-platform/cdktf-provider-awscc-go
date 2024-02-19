package opensearchserverlesssecurityconfig


type OpensearchserverlessSecurityConfigSamlOptions struct {
	// The XML saml provider metadata document that you want to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/opensearchserverless_security_config#metadata OpensearchserverlessSecurityConfig#metadata}
	Metadata *string `field:"required" json:"metadata" yaml:"metadata"`
	// Group attribute for this saml integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/opensearchserverless_security_config#group_attribute OpensearchserverlessSecurityConfig#group_attribute}
	GroupAttribute *string `field:"optional" json:"groupAttribute" yaml:"groupAttribute"`
	// Defines the session timeout in minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/opensearchserverless_security_config#session_timeout OpensearchserverlessSecurityConfig#session_timeout}
	SessionTimeout *float64 `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
	// Custom attribute for this saml integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/opensearchserverless_security_config#user_attribute OpensearchserverlessSecurityConfig#user_attribute}
	UserAttribute *string `field:"optional" json:"userAttribute" yaml:"userAttribute"`
}

