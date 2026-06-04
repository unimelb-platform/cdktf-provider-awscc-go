package gluejob


type GlueJobNotificationProperty struct {
	// It is the number of minutes to wait before sending a job run delay notification after a job run starts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/glue_job#notify_delay_after GlueJob#notify_delay_after}
	NotifyDelayAfter *float64 `field:"optional" json:"notifyDelayAfter" yaml:"notifyDelayAfter"`
}

