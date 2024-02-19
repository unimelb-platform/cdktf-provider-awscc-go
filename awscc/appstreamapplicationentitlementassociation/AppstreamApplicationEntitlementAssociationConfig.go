package appstreamapplicationentitlementassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppstreamApplicationEntitlementAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_application_entitlement_association#application_identifier AppstreamApplicationEntitlementAssociation#application_identifier}.
	ApplicationIdentifier *string `field:"required" json:"applicationIdentifier" yaml:"applicationIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_application_entitlement_association#entitlement_name AppstreamApplicationEntitlementAssociation#entitlement_name}.
	EntitlementName *string `field:"required" json:"entitlementName" yaml:"entitlementName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_application_entitlement_association#stack_name AppstreamApplicationEntitlementAssociation#stack_name}.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
}

