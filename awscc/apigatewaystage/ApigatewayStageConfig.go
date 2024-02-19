package apigatewaystage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigatewayStageConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The string identifier of the associated RestApi.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#rest_api_id ApigatewayStage#rest_api_id}
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Access log settings, including the access log format and access log destination ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#access_log_setting ApigatewayStage#access_log_setting}
	AccessLogSetting *ApigatewayStageAccessLogSetting `field:"optional" json:"accessLogSetting" yaml:"accessLogSetting"`
	// Specifies whether a cache cluster is enabled for the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#cache_cluster_enabled ApigatewayStage#cache_cluster_enabled}
	CacheClusterEnabled interface{} `field:"optional" json:"cacheClusterEnabled" yaml:"cacheClusterEnabled"`
	// The stage's cache capacity in GB.
	//
	// For more information about choosing a cache size, see [Enabling API caching to enhance responsiveness](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#cache_cluster_size ApigatewayStage#cache_cluster_size}
	CacheClusterSize *string `field:"optional" json:"cacheClusterSize" yaml:"cacheClusterSize"`
	// Settings for the canary deployment in this stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#canary_setting ApigatewayStage#canary_setting}
	CanarySetting *ApigatewayStageCanarySetting `field:"optional" json:"canarySetting" yaml:"canarySetting"`
	// The identifier of a client certificate for an API stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#client_certificate_id ApigatewayStage#client_certificate_id}
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// The identifier of the Deployment that the stage points to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#deployment_id ApigatewayStage#deployment_id}
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// The stage's description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#description ApigatewayStage#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The version of the associated API documentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#documentation_version ApigatewayStage#documentation_version}
	DocumentationVersion *string `field:"optional" json:"documentationVersion" yaml:"documentationVersion"`
	// A map that defines the method settings for a Stage resource.
	//
	// Keys (designated as ``/{method_setting_key`` below) are method paths defined as ``{resource_path}/{http_method}`` for an individual method override, or ``/\* /\*`` for overriding all methods in the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#method_settings ApigatewayStage#method_settings}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	MethodSettings interface{} `field:"optional" json:"methodSettings" yaml:"methodSettings"`
	// The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway.
	//
	// Stage names can only contain alphanumeric characters, hyphens, and underscores. Maximum length is 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#stage_name ApigatewayStage#stage_name}
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
	// The collection of tags. Each tag element is associated with a given resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#tags ApigatewayStage#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Specifies whether active tracing with X-ray is enabled for the Stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#tracing_enabled ApigatewayStage#tracing_enabled}
	TracingEnabled interface{} `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
	// A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value.
	//
	// Variable names are limited to alphanumeric characters. Values must match the following regular expression: ``[A-Za-z0-9-._~:/?#&=,]+``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#variables ApigatewayStage#variables}
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
}

