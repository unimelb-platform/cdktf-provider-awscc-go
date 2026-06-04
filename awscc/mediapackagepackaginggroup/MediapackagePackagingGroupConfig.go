package mediapackagepackaginggroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediapackagePackagingGroupConfig struct {
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
	// The ID of the PackagingGroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackage_packaging_group#packaging_group_id MediapackagePackagingGroup#packaging_group_id}
	PackagingGroupId *string `field:"required" json:"packagingGroupId" yaml:"packagingGroupId"`
	// CDN Authorization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackage_packaging_group#authorization MediapackagePackagingGroup#authorization}
	Authorization *MediapackagePackagingGroupAuthorization `field:"optional" json:"authorization" yaml:"authorization"`
	// The configuration parameters for egress access logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackage_packaging_group#egress_access_logs MediapackagePackagingGroup#egress_access_logs}
	EgressAccessLogs *MediapackagePackagingGroupEgressAccessLogs `field:"optional" json:"egressAccessLogs" yaml:"egressAccessLogs"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackage_packaging_group#tags MediapackagePackagingGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

