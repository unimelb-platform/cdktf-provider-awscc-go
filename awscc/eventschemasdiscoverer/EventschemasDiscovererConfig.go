package eventschemasdiscoverer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventschemasDiscovererConfig struct {
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
	// The ARN of the event bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eventschemas_discoverer#source_arn EventschemasDiscoverer#source_arn}
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// Defines whether event schemas from other accounts are discovered. Default is True.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eventschemas_discoverer#cross_account EventschemasDiscoverer#cross_account}
	CrossAccount interface{} `field:"optional" json:"crossAccount" yaml:"crossAccount"`
	// A description for the discoverer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eventschemas_discoverer#description EventschemasDiscoverer#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags associated with the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eventschemas_discoverer#tags EventschemasDiscoverer#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

