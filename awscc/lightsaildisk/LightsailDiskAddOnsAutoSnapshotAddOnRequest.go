package lightsaildisk


type LightsailDiskAddOnsAutoSnapshotAddOnRequest struct {
	// The daily time when an automatic snapshot will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_disk#snapshot_time_of_day LightsailDisk#snapshot_time_of_day}
	SnapshotTimeOfDay *string `field:"optional" json:"snapshotTimeOfDay" yaml:"snapshotTimeOfDay"`
}

