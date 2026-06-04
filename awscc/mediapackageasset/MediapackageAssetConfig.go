package mediapackageasset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediapackageAssetConfig struct {
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
	// The unique identifier for the Asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackage_asset#asset_id MediapackageAsset#asset_id}
	AssetId *string `field:"required" json:"assetId" yaml:"assetId"`
	// The ID of the PackagingGroup for the Asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackage_asset#packaging_group_id MediapackageAsset#packaging_group_id}
	PackagingGroupId *string `field:"required" json:"packagingGroupId" yaml:"packagingGroupId"`
	// ARN of the source object in S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackage_asset#source_arn MediapackageAsset#source_arn}
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// The IAM role_arn used to access the source S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackage_asset#source_role_arn MediapackageAsset#source_role_arn}
	SourceRoleArn *string `field:"required" json:"sourceRoleArn" yaml:"sourceRoleArn"`
	// The list of egress endpoints available for the Asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackage_asset#egress_endpoints MediapackageAsset#egress_endpoints}
	EgressEndpoints interface{} `field:"optional" json:"egressEndpoints" yaml:"egressEndpoints"`
	// The resource ID to include in SPEKE key requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackage_asset#resource_id MediapackageAsset#resource_id}
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackage_asset#tags MediapackageAsset#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

