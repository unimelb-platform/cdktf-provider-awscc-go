package eventschemasregistry

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventschemasRegistryConfig struct {
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
	// A description of the registry to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eventschemas_registry#description EventschemasRegistry#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the schema registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eventschemas_registry#registry_name EventschemasRegistry#registry_name}
	RegistryName *string `field:"optional" json:"registryName" yaml:"registryName"`
	// Tags associated with the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eventschemas_registry#tags EventschemasRegistry#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

