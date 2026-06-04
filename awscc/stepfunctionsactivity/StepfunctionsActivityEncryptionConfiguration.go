package stepfunctionsactivity


type StepfunctionsActivityEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/stepfunctions_activity#kms_data_key_reuse_period_seconds StepfunctionsActivity#kms_data_key_reuse_period_seconds}.
	KmsDataKeyReusePeriodSeconds *float64 `field:"optional" json:"kmsDataKeyReusePeriodSeconds" yaml:"kmsDataKeyReusePeriodSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/stepfunctions_activity#kms_key_id StepfunctionsActivity#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/stepfunctions_activity#type StepfunctionsActivity#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

