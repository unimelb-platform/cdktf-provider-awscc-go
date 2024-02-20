package personalizeschema

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PersonalizeSchemaConfig struct {
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
	// Name for the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/personalize_schema#name PersonalizeSchema#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A schema in Avro JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/personalize_schema#schema PersonalizeSchema#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// The domain of a Domain dataset group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/personalize_schema#domain PersonalizeSchema#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

