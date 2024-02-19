package lightsailinstance


type LightsailInstanceAddOnsAutoSnapshotAddOnRequest struct {
	// The daily time when an automatic snapshot will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_instance#snapshot_time_of_day LightsailInstance#snapshot_time_of_day}
	SnapshotTimeOfDay *string `field:"optional" json:"snapshotTimeOfDay" yaml:"snapshotTimeOfDay"`
}

