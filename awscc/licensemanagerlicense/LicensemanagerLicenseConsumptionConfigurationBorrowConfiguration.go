package licensemanagerlicense


type LicensemanagerLicenseConsumptionConfigurationBorrowConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license#allow_early_check_in LicensemanagerLicense#allow_early_check_in}.
	AllowEarlyCheckIn interface{} `field:"required" json:"allowEarlyCheckIn" yaml:"allowEarlyCheckIn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license#max_time_to_live_in_minutes LicensemanagerLicense#max_time_to_live_in_minutes}.
	MaxTimeToLiveInMinutes *float64 `field:"required" json:"maxTimeToLiveInMinutes" yaml:"maxTimeToLiveInMinutes"`
}

