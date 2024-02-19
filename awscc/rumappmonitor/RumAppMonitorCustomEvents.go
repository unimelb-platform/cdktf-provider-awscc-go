package rumappmonitor


type RumAppMonitorCustomEvents struct {
	// Indicates whether AppMonitor accepts custom events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rum_app_monitor#status RumAppMonitor#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

