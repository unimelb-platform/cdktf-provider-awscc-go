package mediatailorvodsource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediatailorVodSourceConfig struct {
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
	// <p>A list of HTTP package configuration parameters for this VOD source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediatailor_vod_source#http_package_configurations MediatailorVodSource#http_package_configurations}
	HttpPackageConfigurations interface{} `field:"required" json:"httpPackageConfigurations" yaml:"httpPackageConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediatailor_vod_source#source_location_name MediatailorVodSource#source_location_name}.
	SourceLocationName *string `field:"required" json:"sourceLocationName" yaml:"sourceLocationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediatailor_vod_source#vod_source_name MediatailorVodSource#vod_source_name}.
	VodSourceName *string `field:"required" json:"vodSourceName" yaml:"vodSourceName"`
	// The tags to assign to the VOD source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediatailor_vod_source#tags MediatailorVodSource#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

