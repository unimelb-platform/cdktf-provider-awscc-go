package servicecatalogserviceactionassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicecatalogServiceActionAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/servicecatalog_service_action_association#product_id ServicecatalogServiceActionAssociation#product_id}.
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/servicecatalog_service_action_association#provisioning_artifact_id ServicecatalogServiceActionAssociation#provisioning_artifact_id}.
	ProvisioningArtifactId *string `field:"required" json:"provisioningArtifactId" yaml:"provisioningArtifactId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/servicecatalog_service_action_association#service_action_id ServicecatalogServiceActionAssociation#service_action_id}.
	ServiceActionId *string `field:"required" json:"serviceActionId" yaml:"serviceActionId"`
}

