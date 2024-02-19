package route53hostedzone


type Route53HostedZoneVpCs struct {
	// The ID of an Amazon VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53_hosted_zone#vpc_id Route53HostedZone#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The region that an Amazon VPC was created in. See https://docs.aws.amazon.com/general/latest/gr/rande.html for a list of up to date regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53_hosted_zone#vpc_region Route53HostedZone#vpc_region}
	VpcRegion *string `field:"required" json:"vpcRegion" yaml:"vpcRegion"`
}

