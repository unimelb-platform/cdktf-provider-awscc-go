package ceanomalymonitor


type CeAnomalyMonitorResourceTags struct {
	// The key name for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ce_anomaly_monitor#key CeAnomalyMonitor#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ce_anomaly_monitor#value CeAnomalyMonitor#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

