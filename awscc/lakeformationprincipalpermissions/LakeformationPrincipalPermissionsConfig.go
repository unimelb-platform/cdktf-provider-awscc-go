package lakeformationprincipalpermissions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LakeformationPrincipalPermissionsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#permissions LakeformationPrincipalPermissions#permissions}.
	Permissions *[]*string `field:"required" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#permissions_with_grant_option LakeformationPrincipalPermissions#permissions_with_grant_option}.
	PermissionsWithGrantOption *[]*string `field:"required" json:"permissionsWithGrantOption" yaml:"permissionsWithGrantOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#principal LakeformationPrincipalPermissions#principal}.
	Principal *LakeformationPrincipalPermissionsPrincipal `field:"required" json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#resource LakeformationPrincipalPermissions#resource}.
	Resource *LakeformationPrincipalPermissionsResource `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_principal_permissions#catalog LakeformationPrincipalPermissions#catalog}.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
}

