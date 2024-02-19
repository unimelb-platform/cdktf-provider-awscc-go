package codedeploydeploymentconfig


type CodedeployDeploymentConfigZonalConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config#first_zone_monitor_duration_in_seconds CodedeployDeploymentConfig#first_zone_monitor_duration_in_seconds}.
	FirstZoneMonitorDurationInSeconds *float64 `field:"optional" json:"firstZoneMonitorDurationInSeconds" yaml:"firstZoneMonitorDurationInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config#minimum_healthy_hosts_per_zone CodedeployDeploymentConfig#minimum_healthy_hosts_per_zone}.
	MinimumHealthyHostsPerZone *CodedeployDeploymentConfigZonalConfigMinimumHealthyHostsPerZone `field:"optional" json:"minimumHealthyHostsPerZone" yaml:"minimumHealthyHostsPerZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config#monitor_duration_in_seconds CodedeployDeploymentConfig#monitor_duration_in_seconds}.
	MonitorDurationInSeconds *float64 `field:"optional" json:"monitorDurationInSeconds" yaml:"monitorDurationInSeconds"`
}

