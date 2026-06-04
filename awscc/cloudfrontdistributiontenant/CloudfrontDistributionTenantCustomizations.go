package cloudfrontdistributiontenant


type CloudfrontDistributionTenantCustomizations struct {
	// The ACMlong (ACM) certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudfront_distribution_tenant#certificate CloudfrontDistributionTenant#certificate}
	Certificate *CloudfrontDistributionTenantCustomizationsCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// The geographic restrictions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudfront_distribution_tenant#geo_restrictions CloudfrontDistributionTenant#geo_restrictions}
	GeoRestrictions *CloudfrontDistributionTenantCustomizationsGeoRestrictions `field:"optional" json:"geoRestrictions" yaml:"geoRestrictions"`
	// The WAF web ACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudfront_distribution_tenant#web_acl CloudfrontDistributionTenant#web_acl}
	WebAcl *CloudfrontDistributionTenantCustomizationsWebAcl `field:"optional" json:"webAcl" yaml:"webAcl"`
}

