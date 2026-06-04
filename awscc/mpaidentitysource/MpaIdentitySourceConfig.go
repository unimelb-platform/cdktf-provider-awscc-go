package mpaidentitysource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MpaIdentitySourceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mpa_identity_source#identity_source_parameters MpaIdentitySource#identity_source_parameters}.
	IdentitySourceParameters *MpaIdentitySourceIdentitySourceParameters `field:"required" json:"identitySourceParameters" yaml:"identitySourceParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mpa_identity_source#tags MpaIdentitySource#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

