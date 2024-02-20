package apigatewayapikey


type ApigatewayApiKeyStageKeys struct {
	// The string identifier of the associated RestApi.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_api_key#rest_api_id ApigatewayApiKey#rest_api_id}
	RestApiId *string `field:"optional" json:"restApiId" yaml:"restApiId"`
	// The stage name associated with the stage key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_api_key#stage_name ApigatewayApiKey#stage_name}
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

