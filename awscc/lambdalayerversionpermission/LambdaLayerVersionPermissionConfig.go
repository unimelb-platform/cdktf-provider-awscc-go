package lambdalayerversionpermission

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LambdaLayerVersionPermissionConfig struct {
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
	// The API action that grants access to the layer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_layer_version_permission#action LambdaLayerVersionPermission#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The name or Amazon Resource Name (ARN) of the layer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_layer_version_permission#layer_version_arn LambdaLayerVersionPermission#layer_version_arn}
	LayerVersionArn *string `field:"required" json:"layerVersionArn" yaml:"layerVersionArn"`
	// An account ID, or * to grant layer usage permission to all accounts in an organization, or all AWS accounts (if organizationId is not specified).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_layer_version_permission#principal LambdaLayerVersionPermission#principal}
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// With the principal set to *, grant permission to all accounts in the specified organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_layer_version_permission#organization_id LambdaLayerVersionPermission#organization_id}
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
}

