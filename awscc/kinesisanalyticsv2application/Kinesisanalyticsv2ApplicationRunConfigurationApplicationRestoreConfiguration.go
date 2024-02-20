package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationRunConfigurationApplicationRestoreConfiguration struct {
	// Specifies how the application should be restored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#application_restore_type Kinesisanalyticsv2Application#application_restore_type}
	ApplicationRestoreType *string `field:"required" json:"applicationRestoreType" yaml:"applicationRestoreType"`
	// The identifier of an existing snapshot of application state to use to restart an application.
	//
	// The application uses this value if RESTORE_FROM_CUSTOM_SNAPSHOT is specified for the ApplicationRestoreType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#snapshot_name Kinesisanalyticsv2Application#snapshot_name}
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
}

