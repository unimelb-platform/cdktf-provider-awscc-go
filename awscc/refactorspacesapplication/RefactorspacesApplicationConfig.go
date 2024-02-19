package refactorspacesapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RefactorspacesApplicationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/refactorspaces_application#environment_identifier RefactorspacesApplication#environment_identifier}.
	EnvironmentIdentifier *string `field:"required" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/refactorspaces_application#name RefactorspacesApplication#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/refactorspaces_application#proxy_type RefactorspacesApplication#proxy_type}.
	ProxyType *string `field:"required" json:"proxyType" yaml:"proxyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/refactorspaces_application#vpc_id RefactorspacesApplication#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/refactorspaces_application#api_gateway_proxy RefactorspacesApplication#api_gateway_proxy}.
	ApiGatewayProxy *RefactorspacesApplicationApiGatewayProxy `field:"optional" json:"apiGatewayProxy" yaml:"apiGatewayProxy"`
	// Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/refactorspaces_application#tags RefactorspacesApplication#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

