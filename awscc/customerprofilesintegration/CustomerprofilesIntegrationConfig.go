package customerprofilesintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomerprofilesIntegrationConfig struct {
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
	// The unique name of the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#domain_name CustomerprofilesIntegration#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#flow_definition CustomerprofilesIntegration#flow_definition}.
	FlowDefinition *CustomerprofilesIntegrationFlowDefinition `field:"optional" json:"flowDefinition" yaml:"flowDefinition"`
	// The name of the ObjectType defined for the 3rd party data in Profile Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#object_type_name CustomerprofilesIntegration#object_type_name}
	ObjectTypeName *string `field:"optional" json:"objectTypeName" yaml:"objectTypeName"`
	// The mapping between 3rd party event types and ObjectType names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#object_type_names CustomerprofilesIntegration#object_type_names}
	ObjectTypeNames interface{} `field:"optional" json:"objectTypeNames" yaml:"objectTypeNames"`
	// The tags (keys and values) associated with the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#tags CustomerprofilesIntegration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The URI of the S3 bucket or any other type of data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#uri CustomerprofilesIntegration#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

