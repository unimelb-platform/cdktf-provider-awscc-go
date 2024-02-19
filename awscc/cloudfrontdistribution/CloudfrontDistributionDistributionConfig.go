package cloudfrontdistribution


type CloudfrontDistributionDistributionConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#default_cache_behavior CloudfrontDistribution#default_cache_behavior}.
	DefaultCacheBehavior *CloudfrontDistributionDistributionConfigDefaultCacheBehavior `field:"required" json:"defaultCacheBehavior" yaml:"defaultCacheBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#enabled CloudfrontDistribution#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#aliases CloudfrontDistribution#aliases}.
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#cache_behaviors CloudfrontDistribution#cache_behaviors}.
	CacheBehaviors interface{} `field:"optional" json:"cacheBehaviors" yaml:"cacheBehaviors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#cnames CloudfrontDistribution#cnames}.
	Cnames *[]*string `field:"optional" json:"cnames" yaml:"cnames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#comment CloudfrontDistribution#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#continuous_deployment_policy_id CloudfrontDistribution#continuous_deployment_policy_id}.
	ContinuousDeploymentPolicyId *string `field:"optional" json:"continuousDeploymentPolicyId" yaml:"continuousDeploymentPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#custom_error_responses CloudfrontDistribution#custom_error_responses}.
	CustomErrorResponses interface{} `field:"optional" json:"customErrorResponses" yaml:"customErrorResponses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#custom_origin CloudfrontDistribution#custom_origin}.
	CustomOrigin *CloudfrontDistributionDistributionConfigCustomOrigin `field:"optional" json:"customOrigin" yaml:"customOrigin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#default_root_object CloudfrontDistribution#default_root_object}.
	DefaultRootObject *string `field:"optional" json:"defaultRootObject" yaml:"defaultRootObject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#http_version CloudfrontDistribution#http_version}.
	HttpVersion *string `field:"optional" json:"httpVersion" yaml:"httpVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#ipv6_enabled CloudfrontDistribution#ipv6_enabled}.
	Ipv6Enabled interface{} `field:"optional" json:"ipv6Enabled" yaml:"ipv6Enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#logging CloudfrontDistribution#logging}.
	Logging *CloudfrontDistributionDistributionConfigLogging `field:"optional" json:"logging" yaml:"logging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#origin_groups CloudfrontDistribution#origin_groups}.
	OriginGroups *CloudfrontDistributionDistributionConfigOriginGroups `field:"optional" json:"originGroups" yaml:"originGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#origins CloudfrontDistribution#origins}.
	Origins interface{} `field:"optional" json:"origins" yaml:"origins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#price_class CloudfrontDistribution#price_class}.
	PriceClass *string `field:"optional" json:"priceClass" yaml:"priceClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#restrictions CloudfrontDistribution#restrictions}.
	Restrictions *CloudfrontDistributionDistributionConfigRestrictions `field:"optional" json:"restrictions" yaml:"restrictions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#s3_origin CloudfrontDistribution#s3_origin}.
	S3Origin *CloudfrontDistributionDistributionConfigS3Origin `field:"optional" json:"s3Origin" yaml:"s3Origin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#staging CloudfrontDistribution#staging}.
	Staging interface{} `field:"optional" json:"staging" yaml:"staging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#viewer_certificate CloudfrontDistribution#viewer_certificate}.
	ViewerCertificate *CloudfrontDistributionDistributionConfigViewerCertificate `field:"optional" json:"viewerCertificate" yaml:"viewerCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#web_acl_id CloudfrontDistribution#web_acl_id}.
	WebAclId *string `field:"optional" json:"webAclId" yaml:"webAclId"`
}

