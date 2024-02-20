package dynamodbtable


type DynamodbTableSseSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#sse_enabled DynamodbTable#sse_enabled}.
	SseEnabled interface{} `field:"required" json:"sseEnabled" yaml:"sseEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#kms_master_key_id DynamodbTable#kms_master_key_id}.
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#sse_type DynamodbTable#sse_type}.
	SseType *string `field:"optional" json:"sseType" yaml:"sseType"`
}

