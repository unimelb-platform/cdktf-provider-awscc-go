package appconfigextensionassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppconfigExtensionAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_extension_association#extension_identifier AppconfigExtensionAssociation#extension_identifier}.
	ExtensionIdentifier *string `field:"optional" json:"extensionIdentifier" yaml:"extensionIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_extension_association#extension_version_number AppconfigExtensionAssociation#extension_version_number}.
	ExtensionVersionNumber *float64 `field:"optional" json:"extensionVersionNumber" yaml:"extensionVersionNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_extension_association#parameters AppconfigExtensionAssociation#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_extension_association#resource_identifier AppconfigExtensionAssociation#resource_identifier}.
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_extension_association#tags AppconfigExtensionAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

