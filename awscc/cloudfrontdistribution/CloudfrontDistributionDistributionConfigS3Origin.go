package cloudfrontdistribution


type CloudfrontDistributionDistributionConfigS3Origin struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#dns_name CloudfrontDistribution#dns_name}.
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#origin_access_identity CloudfrontDistribution#origin_access_identity}.
	OriginAccessIdentity *string `field:"optional" json:"originAccessIdentity" yaml:"originAccessIdentity"`
}

