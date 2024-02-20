package apigatewaystage


type ApigatewayStageCanarySetting struct {
	// The ID of the canary deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#deployment_id ApigatewayStage#deployment_id}
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// The percent (0-100) of traffic diverted to a canary deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#percent_traffic ApigatewayStage#percent_traffic}
	PercentTraffic *float64 `field:"optional" json:"percentTraffic" yaml:"percentTraffic"`
	// Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary.
	//
	// These stage variables are represented as a string-to-string map between stage variable names and their values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#stage_variable_overrides ApigatewayStage#stage_variable_overrides}
	StageVariableOverrides *map[string]*string `field:"optional" json:"stageVariableOverrides" yaml:"stageVariableOverrides"`
	// A Boolean flag to indicate whether the canary deployment uses the stage cache or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#use_stage_cache ApigatewayStage#use_stage_cache}
	UseStageCache interface{} `field:"optional" json:"useStageCache" yaml:"useStageCache"`
}

