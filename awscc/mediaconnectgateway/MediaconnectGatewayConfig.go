package mediaconnectgateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaconnectGatewayConfig struct {
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
	// The range of IP addresses that contribute content or initiate output requests for flows communicating with this gateway.
	//
	// These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_gateway#egress_cidr_blocks MediaconnectGateway#egress_cidr_blocks}
	EgressCidrBlocks *[]*string `field:"required" json:"egressCidrBlocks" yaml:"egressCidrBlocks"`
	// The name of the gateway. This name can not be modified after the gateway is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_gateway#name MediaconnectGateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of networks in the gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_gateway#networks MediaconnectGateway#networks}
	Networks interface{} `field:"required" json:"networks" yaml:"networks"`
}

