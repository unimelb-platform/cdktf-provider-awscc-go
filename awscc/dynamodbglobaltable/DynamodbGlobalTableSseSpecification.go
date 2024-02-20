package dynamodbglobaltable


type DynamodbGlobalTableSseSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#sse_enabled DynamodbGlobalTable#sse_enabled}.
	SseEnabled interface{} `field:"required" json:"sseEnabled" yaml:"sseEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#sse_type DynamodbGlobalTable#sse_type}.
	SseType *string `field:"optional" json:"sseType" yaml:"sseType"`
}

