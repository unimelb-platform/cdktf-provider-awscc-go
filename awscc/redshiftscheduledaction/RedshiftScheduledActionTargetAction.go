package redshiftscheduledaction


type RedshiftScheduledActionTargetAction struct {
	// Describes a pause cluster operation. For example, a scheduled action to run the `PauseCluster` API operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshift_scheduled_action#pause_cluster RedshiftScheduledAction#pause_cluster}
	PauseCluster *RedshiftScheduledActionTargetActionPauseCluster `field:"optional" json:"pauseCluster" yaml:"pauseCluster"`
	// Describes a resize cluster operation. For example, a scheduled action to run the `ResizeCluster` API operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshift_scheduled_action#resize_cluster RedshiftScheduledAction#resize_cluster}
	ResizeCluster *RedshiftScheduledActionTargetActionResizeCluster `field:"optional" json:"resizeCluster" yaml:"resizeCluster"`
	// Describes a resume cluster operation. For example, a scheduled action to run the `ResumeCluster` API operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshift_scheduled_action#resume_cluster RedshiftScheduledAction#resume_cluster}
	ResumeCluster *RedshiftScheduledActionTargetActionResumeCluster `field:"optional" json:"resumeCluster" yaml:"resumeCluster"`
}

