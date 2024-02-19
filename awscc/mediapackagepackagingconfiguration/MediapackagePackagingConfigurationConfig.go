package mediapackagepackagingconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediapackagePackagingConfigurationConfig struct {
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
	// The ID of the PackagingConfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#id MediapackagePackagingConfiguration#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The ID of a PackagingGroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#packaging_group_id MediapackagePackagingConfiguration#packaging_group_id}
	PackagingGroupId *string `field:"required" json:"packagingGroupId" yaml:"packagingGroupId"`
	// A CMAF packaging configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#cmaf_package MediapackagePackagingConfiguration#cmaf_package}
	CmafPackage *MediapackagePackagingConfigurationCmafPackage `field:"optional" json:"cmafPackage" yaml:"cmafPackage"`
	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#dash_package MediapackagePackagingConfiguration#dash_package}
	DashPackage *MediapackagePackagingConfigurationDashPackage `field:"optional" json:"dashPackage" yaml:"dashPackage"`
	// An HTTP Live Streaming (HLS) packaging configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#hls_package MediapackagePackagingConfiguration#hls_package}
	HlsPackage *MediapackagePackagingConfigurationHlsPackage `field:"optional" json:"hlsPackage" yaml:"hlsPackage"`
	// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#mss_package MediapackagePackagingConfiguration#mss_package}
	MssPackage *MediapackagePackagingConfigurationMssPackage `field:"optional" json:"mssPackage" yaml:"mssPackage"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#tags MediapackagePackagingConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

