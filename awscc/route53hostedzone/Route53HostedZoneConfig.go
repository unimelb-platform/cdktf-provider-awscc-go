package route53hostedzone

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53HostedZoneConfig struct {
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
	// A complex type that contains an optional comment.
	//
	// If you don't want to specify a comment, omit the HostedZoneConfig and Comment elements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53_hosted_zone#hosted_zone_config Route53HostedZone#hosted_zone_config}
	HostedZoneConfig *Route53HostedZoneHostedZoneConfig `field:"optional" json:"hostedZoneConfig" yaml:"hostedZoneConfig"`
	// Adds, edits, or deletes tags for a health check or a hosted zone.
	//
	// For information about using tags for cost allocation, see Using Cost Allocation Tags in the AWS Billing and Cost Management User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53_hosted_zone#hosted_zone_tags Route53HostedZone#hosted_zone_tags}
	HostedZoneTags interface{} `field:"optional" json:"hostedZoneTags" yaml:"hostedZoneTags"`
	// The name of the domain.
	//
	// Specify a fully qualified domain name, for example, www.example.com. The trailing dot is optional; Amazon Route 53 assumes that the domain name is fully qualified. This means that Route 53 treats www.example.com (without a trailing dot) and www.example.com. (with a trailing dot) as identical.
	//
	// If you're creating a public hosted zone, this is the name you have registered with your DNS registrar. If your domain name is registered with a registrar other than Route 53, change the name servers for your domain to the set of NameServers that are returned by the Fn::GetAtt intrinsic function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53_hosted_zone#name Route53HostedZone#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A complex type that contains information about a configuration for DNS query logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53_hosted_zone#query_logging_config Route53HostedZone#query_logging_config}
	QueryLoggingConfig *Route53HostedZoneQueryLoggingConfig `field:"optional" json:"queryLoggingConfig" yaml:"queryLoggingConfig"`
	// A complex type that contains information about the VPCs that are associated with the specified hosted zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53_hosted_zone#vp_cs Route53HostedZone#vp_cs}
	VpCs interface{} `field:"optional" json:"vpCs" yaml:"vpCs"`
}

