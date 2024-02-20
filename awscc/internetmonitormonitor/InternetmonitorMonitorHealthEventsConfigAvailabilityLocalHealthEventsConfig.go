package internetmonitormonitor


type InternetmonitorMonitorHealthEventsConfigAvailabilityLocalHealthEventsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/internetmonitor_monitor#health_score_threshold InternetmonitorMonitor#health_score_threshold}.
	HealthScoreThreshold *float64 `field:"optional" json:"healthScoreThreshold" yaml:"healthScoreThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/internetmonitor_monitor#min_traffic_impact InternetmonitorMonitor#min_traffic_impact}.
	MinTrafficImpact *float64 `field:"optional" json:"minTrafficImpact" yaml:"minTrafficImpact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/internetmonitor_monitor#status InternetmonitorMonitor#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

