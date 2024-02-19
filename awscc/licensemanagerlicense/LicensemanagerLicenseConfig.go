package licensemanagerlicense

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LicensemanagerLicenseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license#consumption_configuration LicensemanagerLicense#consumption_configuration}.
	ConsumptionConfiguration *LicensemanagerLicenseConsumptionConfiguration `field:"required" json:"consumptionConfiguration" yaml:"consumptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license#entitlements LicensemanagerLicense#entitlements}.
	Entitlements interface{} `field:"required" json:"entitlements" yaml:"entitlements"`
	// Home region for the created license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license#home_region LicensemanagerLicense#home_region}
	HomeRegion *string `field:"required" json:"homeRegion" yaml:"homeRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license#issuer LicensemanagerLicense#issuer}.
	Issuer *LicensemanagerLicenseIssuer `field:"required" json:"issuer" yaml:"issuer"`
	// Name for the created license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license#license_name LicensemanagerLicense#license_name}
	LicenseName *string `field:"required" json:"licenseName" yaml:"licenseName"`
	// Product name for the created license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license#product_name LicensemanagerLicense#product_name}
	ProductName *string `field:"required" json:"productName" yaml:"productName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license#validity LicensemanagerLicense#validity}.
	Validity *LicensemanagerLicenseValidity `field:"required" json:"validity" yaml:"validity"`
	// Beneficiary of the license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license#beneficiary LicensemanagerLicense#beneficiary}
	Beneficiary *string `field:"optional" json:"beneficiary" yaml:"beneficiary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license#license_metadata LicensemanagerLicense#license_metadata}.
	LicenseMetadata interface{} `field:"optional" json:"licenseMetadata" yaml:"licenseMetadata"`
	// ProductSKU of the license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license#product_sku LicensemanagerLicense#product_sku}
	ProductSku *string `field:"optional" json:"productSku" yaml:"productSku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license#status LicensemanagerLicense#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

