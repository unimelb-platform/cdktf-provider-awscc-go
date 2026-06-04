package opensearchserverlesssecurityconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpensearchserverlessSecurityConfigConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Security config description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/opensearchserverless_security_config#description OpensearchserverlessSecurityConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Describes IAM Identity Center options for an OpenSearch Serverless security configuration in the form of a key-value map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/opensearchserverless_security_config#iam_identity_center_options OpensearchserverlessSecurityConfig#iam_identity_center_options}
	IamIdentityCenterOptions *OpensearchserverlessSecurityConfigIamIdentityCenterOptions `field:"optional" json:"iamIdentityCenterOptions" yaml:"iamIdentityCenterOptions"`
	// The friendly name of the security config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/opensearchserverless_security_config#name OpensearchserverlessSecurityConfig#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Describes saml options in form of key value map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/opensearchserverless_security_config#saml_options OpensearchserverlessSecurityConfig#saml_options}
	SamlOptions *OpensearchserverlessSecurityConfigSamlOptions `field:"optional" json:"samlOptions" yaml:"samlOptions"`
	// Config type for security config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/opensearchserverless_security_config#type OpensearchserverlessSecurityConfig#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

