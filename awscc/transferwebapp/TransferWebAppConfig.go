package transferwebapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransferWebAppConfig struct {
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
	// You can provide a structure that contains the details for the identity provider to use with your web app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/transfer_web_app#identity_provider_details TransferWebApp#identity_provider_details}
	IdentityProviderDetails *TransferWebAppIdentityProviderDetails `field:"required" json:"identityProviderDetails" yaml:"identityProviderDetails"`
	// The AccessEndpoint is the URL that you provide to your users for them to interact with the Transfer Family web app.
	//
	// You can specify a custom URL or use the default value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/transfer_web_app#access_endpoint TransferWebApp#access_endpoint}
	AccessEndpoint *string `field:"optional" json:"accessEndpoint" yaml:"accessEndpoint"`
	// Key-value pairs that can be used to group and search for web apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/transfer_web_app#tags TransferWebApp#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/transfer_web_app#web_app_customization TransferWebApp#web_app_customization}.
	WebAppCustomization *TransferWebAppWebAppCustomization `field:"optional" json:"webAppCustomization" yaml:"webAppCustomization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/transfer_web_app#web_app_endpoint_policy TransferWebApp#web_app_endpoint_policy}.
	WebAppEndpointPolicy *string `field:"optional" json:"webAppEndpointPolicy" yaml:"webAppEndpointPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/transfer_web_app#web_app_units TransferWebApp#web_app_units}.
	WebAppUnits *TransferWebAppWebAppUnits `field:"optional" json:"webAppUnits" yaml:"webAppUnits"`
}

