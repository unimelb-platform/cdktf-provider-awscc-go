package lambdalayerversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LambdaLayerVersionConfig struct {
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
	// The function layer archive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_layer_version#content LambdaLayerVersion#content}
	Content *LambdaLayerVersionContent `field:"required" json:"content" yaml:"content"`
	// A list of compatible instruction set architectures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_layer_version#compatible_architectures LambdaLayerVersion#compatible_architectures}
	CompatibleArchitectures *[]*string `field:"optional" json:"compatibleArchitectures" yaml:"compatibleArchitectures"`
	// A list of compatible function runtimes. Used for filtering with ListLayers and ListLayerVersions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_layer_version#compatible_runtimes LambdaLayerVersion#compatible_runtimes}
	CompatibleRuntimes *[]*string `field:"optional" json:"compatibleRuntimes" yaml:"compatibleRuntimes"`
	// The description of the version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_layer_version#description LambdaLayerVersion#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name or Amazon Resource Name (ARN) of the layer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_layer_version#layer_name LambdaLayerVersion#layer_name}
	LayerName *string `field:"optional" json:"layerName" yaml:"layerName"`
	// The layer's software license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_layer_version#license_info LambdaLayerVersion#license_info}
	LicenseInfo *string `field:"optional" json:"licenseInfo" yaml:"licenseInfo"`
}

