package licensemanagergrant

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LicensemanagerGrantConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_grant#allowed_operations LicensemanagerGrant#allowed_operations}.
	AllowedOperations *[]*string `field:"optional" json:"allowedOperations" yaml:"allowedOperations"`
	// Name for the created Grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_grant#grant_name LicensemanagerGrant#grant_name}
	GrantName *string `field:"optional" json:"grantName" yaml:"grantName"`
	// Home region for the created grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_grant#home_region LicensemanagerGrant#home_region}
	HomeRegion *string `field:"optional" json:"homeRegion" yaml:"homeRegion"`
	// License Arn for the grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_grant#license_arn LicensemanagerGrant#license_arn}
	LicenseArn *string `field:"optional" json:"licenseArn" yaml:"licenseArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_grant#principals LicensemanagerGrant#principals}.
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_grant#status LicensemanagerGrant#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

