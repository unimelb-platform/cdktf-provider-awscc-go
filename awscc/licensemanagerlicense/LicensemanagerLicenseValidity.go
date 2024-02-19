package licensemanagerlicense


type LicensemanagerLicenseValidity struct {
	// Validity begin date for the license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license#begin LicensemanagerLicense#begin}
	Begin *string `field:"required" json:"begin" yaml:"begin"`
	// Validity begin date for the license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license#end LicensemanagerLicense#end}
	End *string `field:"required" json:"end" yaml:"end"`
}

