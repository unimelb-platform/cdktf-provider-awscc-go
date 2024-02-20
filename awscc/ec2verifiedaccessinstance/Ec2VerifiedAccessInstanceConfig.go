package ec2verifiedaccessinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2VerifiedAccessInstanceConfig struct {
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
	// A description for the AWS Verified Access instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_instance#description Ec2VerifiedAccessInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether FIPS is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_instance#fips_enabled Ec2VerifiedAccessInstance#fips_enabled}
	FipsEnabled interface{} `field:"optional" json:"fipsEnabled" yaml:"fipsEnabled"`
	// The configuration options for AWS Verified Access instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_instance#logging_configurations Ec2VerifiedAccessInstance#logging_configurations}
	LoggingConfigurations *Ec2VerifiedAccessInstanceLoggingConfigurations `field:"optional" json:"loggingConfigurations" yaml:"loggingConfigurations"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_instance#tags Ec2VerifiedAccessInstance#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The IDs of the AWS Verified Access trust providers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_instance#verified_access_trust_provider_ids Ec2VerifiedAccessInstance#verified_access_trust_provider_ids}
	VerifiedAccessTrustProviderIds *[]*string `field:"optional" json:"verifiedAccessTrustProviderIds" yaml:"verifiedAccessTrustProviderIds"`
	// AWS Verified Access trust providers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_instance#verified_access_trust_providers Ec2VerifiedAccessInstance#verified_access_trust_providers}
	VerifiedAccessTrustProviders interface{} `field:"optional" json:"verifiedAccessTrustProviders" yaml:"verifiedAccessTrustProviders"`
}

