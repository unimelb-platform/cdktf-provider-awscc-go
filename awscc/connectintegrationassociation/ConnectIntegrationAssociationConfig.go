package connectintegrationassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConnectIntegrationAssociationConfig struct {
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
	// Amazon Connect instance identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_integration_association#instance_id ConnectIntegrationAssociation#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// ARN of Integration being associated with the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_integration_association#integration_arn ConnectIntegrationAssociation#integration_arn}
	IntegrationArn *string `field:"required" json:"integrationArn" yaml:"integrationArn"`
	// Specifies the integration type to be associated with the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_integration_association#integration_type ConnectIntegrationAssociation#integration_type}
	IntegrationType *string `field:"required" json:"integrationType" yaml:"integrationType"`
}

