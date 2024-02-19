package lookoutmetricsalert


type LookoutmetricsAlertAction struct {
	// Configuration options for a Lambda alert action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_alert#lambda_configuration LookoutmetricsAlert#lambda_configuration}
	LambdaConfiguration *LookoutmetricsAlertActionLambdaConfiguration `field:"optional" json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
	// Configuration options for an SNS alert action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_alert#sns_configuration LookoutmetricsAlert#sns_configuration}
	SnsConfiguration *LookoutmetricsAlertActionSnsConfiguration `field:"optional" json:"snsConfiguration" yaml:"snsConfiguration"`
}

