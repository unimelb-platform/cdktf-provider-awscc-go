package cloudfrontdistribution


type CloudfrontDistributionDistributionConfigOrigins struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#domain_name CloudfrontDistribution#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#id CloudfrontDistribution#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#connection_attempts CloudfrontDistribution#connection_attempts}.
	ConnectionAttempts *float64 `field:"optional" json:"connectionAttempts" yaml:"connectionAttempts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#connection_timeout CloudfrontDistribution#connection_timeout}.
	ConnectionTimeout *float64 `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#custom_origin_config CloudfrontDistribution#custom_origin_config}.
	CustomOriginConfig *CloudfrontDistributionDistributionConfigOriginsCustomOriginConfig `field:"optional" json:"customOriginConfig" yaml:"customOriginConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#origin_access_control_id CloudfrontDistribution#origin_access_control_id}.
	OriginAccessControlId *string `field:"optional" json:"originAccessControlId" yaml:"originAccessControlId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#origin_custom_headers CloudfrontDistribution#origin_custom_headers}.
	OriginCustomHeaders interface{} `field:"optional" json:"originCustomHeaders" yaml:"originCustomHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#origin_path CloudfrontDistribution#origin_path}.
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#origin_shield CloudfrontDistribution#origin_shield}.
	OriginShield *CloudfrontDistributionDistributionConfigOriginsOriginShield `field:"optional" json:"originShield" yaml:"originShield"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#s3_origin_config CloudfrontDistribution#s3_origin_config}.
	S3OriginConfig *CloudfrontDistributionDistributionConfigOriginsS3OriginConfig `field:"optional" json:"s3OriginConfig" yaml:"s3OriginConfig"`
}

