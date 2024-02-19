package licensemanagerlicense


type LicensemanagerLicenseIssuer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license#name LicensemanagerLicense#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license#sign_key LicensemanagerLicense#sign_key}.
	SignKey *string `field:"optional" json:"signKey" yaml:"signKey"`
}

