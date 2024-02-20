package apprunnervpcingressconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApprunnerVpcIngressConnectionConfig struct {
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
	// The configuration of customer?s VPC and related VPC endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_vpc_ingress_connection#ingress_vpc_configuration ApprunnerVpcIngressConnection#ingress_vpc_configuration}
	IngressVpcConfiguration *ApprunnerVpcIngressConnectionIngressVpcConfiguration `field:"required" json:"ingressVpcConfiguration" yaml:"ingressVpcConfiguration"`
	// The Amazon Resource Name (ARN) of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_vpc_ingress_connection#service_arn ApprunnerVpcIngressConnection#service_arn}
	ServiceArn *string `field:"required" json:"serviceArn" yaml:"serviceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_vpc_ingress_connection#tags ApprunnerVpcIngressConnection#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The customer-provided Vpc Ingress Connection name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_vpc_ingress_connection#vpc_ingress_connection_name ApprunnerVpcIngressConnection#vpc_ingress_connection_name}
	VpcIngressConnectionName *string `field:"optional" json:"vpcIngressConnectionName" yaml:"vpcIngressConnectionName"`
}

