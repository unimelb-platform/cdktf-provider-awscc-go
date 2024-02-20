package schedulerschedulegroup


type SchedulerScheduleGroupTags struct {
	// Key for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/scheduler_schedule_group#key SchedulerScheduleGroup#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/scheduler_schedule_group#value SchedulerScheduleGroup#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

