package ecsservice


type EcsServiceDeploymentConfigurationAlarms struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#alarm_names EcsService#alarm_names}.
	AlarmNames *[]*string `field:"required" json:"alarmNames" yaml:"alarmNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#enable EcsService#enable}.
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#rollback EcsService#rollback}.
	Rollback interface{} `field:"required" json:"rollback" yaml:"rollback"`
}

