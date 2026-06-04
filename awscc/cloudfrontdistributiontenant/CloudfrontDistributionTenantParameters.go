package cloudfrontdistributiontenant


type CloudfrontDistributionTenantParameters struct {
	// The parameter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudfront_distribution_tenant#name CloudfrontDistributionTenant#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The parameter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudfront_distribution_tenant#value CloudfrontDistributionTenant#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

