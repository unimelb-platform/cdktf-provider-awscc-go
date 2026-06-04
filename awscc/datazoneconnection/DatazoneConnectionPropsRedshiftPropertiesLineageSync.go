package datazoneconnection


type DatazoneConnectionPropsRedshiftPropertiesLineageSync struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_connection#enabled DatazoneConnection#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Lineage Sync Schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_connection#schedule DatazoneConnection#schedule}
	Schedule *DatazoneConnectionPropsRedshiftPropertiesLineageSyncSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

