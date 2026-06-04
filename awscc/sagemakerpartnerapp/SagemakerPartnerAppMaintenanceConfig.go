package sagemakerpartnerapp


type SagemakerPartnerAppMaintenanceConfig struct {
	// The maintenance window start day and time for the PartnerApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_partner_app#maintenance_window_start SagemakerPartnerApp#maintenance_window_start}
	MaintenanceWindowStart *string `field:"optional" json:"maintenanceWindowStart" yaml:"maintenanceWindowStart"`
}

