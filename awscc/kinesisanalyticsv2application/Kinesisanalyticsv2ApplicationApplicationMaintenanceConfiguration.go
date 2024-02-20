package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationMaintenanceConfiguration struct {
	// The start time for the maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#application_maintenance_window_start_time Kinesisanalyticsv2Application#application_maintenance_window_start_time}
	ApplicationMaintenanceWindowStartTime *string `field:"required" json:"applicationMaintenanceWindowStartTime" yaml:"applicationMaintenanceWindowStartTime"`
}

