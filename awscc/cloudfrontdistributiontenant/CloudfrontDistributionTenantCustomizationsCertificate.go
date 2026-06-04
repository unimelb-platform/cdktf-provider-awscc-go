package cloudfrontdistributiontenant


type CloudfrontDistributionTenantCustomizationsCertificate struct {
	// The Amazon Resource Name (ARN) of the ACM certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudfront_distribution_tenant#arn CloudfrontDistributionTenant#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

