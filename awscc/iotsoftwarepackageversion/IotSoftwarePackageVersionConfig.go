package iotsoftwarepackageversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotSoftwarePackageVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_software_package_version#package_name IotSoftwarePackageVersion#package_name}.
	PackageName *string `field:"required" json:"packageName" yaml:"packageName"`
	// The artifact location of the package version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_software_package_version#artifact IotSoftwarePackageVersion#artifact}
	Artifact *IotSoftwarePackageVersionArtifact `field:"optional" json:"artifact" yaml:"artifact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_software_package_version#attributes IotSoftwarePackageVersion#attributes}.
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_software_package_version#description IotSoftwarePackageVersion#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The inline json job document associated with a software package version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_software_package_version#recipe IotSoftwarePackageVersion#recipe}
	Recipe *string `field:"optional" json:"recipe" yaml:"recipe"`
	// The sbom zip archive location of the package version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_software_package_version#sbom IotSoftwarePackageVersion#sbom}
	Sbom *IotSoftwarePackageVersionSbom `field:"optional" json:"sbom" yaml:"sbom"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_software_package_version#tags IotSoftwarePackageVersion#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_software_package_version#version_name IotSoftwarePackageVersion#version_name}.
	VersionName *string `field:"optional" json:"versionName" yaml:"versionName"`
}

