package servicecatalogappregistryresourceassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicecatalogappregistryResourceAssociationConfig struct {
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
	// The name or the Id of the Application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/servicecatalogappregistry_resource_association#application ServicecatalogappregistryResourceAssociation#application}
	Application *string `field:"required" json:"application" yaml:"application"`
	// The name or the Id of the Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/servicecatalogappregistry_resource_association#resource ServicecatalogappregistryResourceAssociation#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// The type of the CFN Resource for now it's enum CFN_STACK.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/servicecatalogappregistry_resource_association#resource_type ServicecatalogappregistryResourceAssociation#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

