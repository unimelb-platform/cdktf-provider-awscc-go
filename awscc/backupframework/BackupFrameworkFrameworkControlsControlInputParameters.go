package backupframework


type BackupFrameworkFrameworkControlsControlInputParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_framework#parameter_name BackupFramework#parameter_name}.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_framework#parameter_value BackupFramework#parameter_value}.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

