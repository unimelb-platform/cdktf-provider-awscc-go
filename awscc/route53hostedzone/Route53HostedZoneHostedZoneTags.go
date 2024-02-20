package route53hostedzone


type Route53HostedZoneHostedZoneTags struct {
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53_hosted_zone#key Route53HostedZone#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53_hosted_zone#value Route53HostedZone#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

