package xraysamplingrule


type XraySamplingRuleSamplingRuleRecordSamplingRule struct {
	// The percentage of matching requests to instrument, after the reservoir is exhausted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#fixed_rate XraySamplingRule#fixed_rate}
	FixedRate *float64 `field:"required" json:"fixedRate" yaml:"fixedRate"`
	// Matches the hostname from a request URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#host XraySamplingRule#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Matches the HTTP method from a request URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#http_method XraySamplingRule#http_method}
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// The priority of the sampling rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#priority XraySamplingRule#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// A fixed number of matching requests to instrument per second, prior to applying the fixed rate.
	//
	// The reservoir is not used directly by services, but applies to all services using the rule collectively.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#reservoir_size XraySamplingRule#reservoir_size}
	ReservoirSize *float64 `field:"required" json:"reservoirSize" yaml:"reservoirSize"`
	// Matches the ARN of the AWS resource on which the service runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#resource_arn XraySamplingRule#resource_arn}
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Matches the name that the service uses to identify itself in segments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#service_name XraySamplingRule#service_name}
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// Matches the origin that the service uses to identify its type in segments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#service_type XraySamplingRule#service_type}
	ServiceType *string `field:"required" json:"serviceType" yaml:"serviceType"`
	// Matches the path from a request URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#url_path XraySamplingRule#url_path}
	UrlPath *string `field:"required" json:"urlPath" yaml:"urlPath"`
	// Matches attributes derived from the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#attributes XraySamplingRule#attributes}
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#rule_arn XraySamplingRule#rule_arn}
	RuleArn *string `field:"optional" json:"ruleArn" yaml:"ruleArn"`
	// The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#rule_name XraySamplingRule#rule_name}
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// The version of the sampling rule format (1).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#version XraySamplingRule#version}
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

