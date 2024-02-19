package apigatewaydeployment


type ApigatewayDeploymentDeploymentCanarySettings struct {
	// The percentage (0.0-100.0) of traffic routed to the canary deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_deployment#percent_traffic ApigatewayDeployment#percent_traffic}
	PercentTraffic *float64 `field:"optional" json:"percentTraffic" yaml:"percentTraffic"`
	// A stage variable overrides used for the canary release deployment.
	//
	// They can override existing stage variables or add new stage variables for the canary release deployment. These stage variables are represented as a string-to-string map between stage variable names and their values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_deployment#stage_variable_overrides ApigatewayDeployment#stage_variable_overrides}
	StageVariableOverrides *map[string]*string `field:"optional" json:"stageVariableOverrides" yaml:"stageVariableOverrides"`
	// A Boolean flag to indicate whether the canary release deployment uses the stage cache or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_deployment#use_stage_cache ApigatewayDeployment#use_stage_cache}
	UseStageCache interface{} `field:"optional" json:"useStageCache" yaml:"useStageCache"`
}

