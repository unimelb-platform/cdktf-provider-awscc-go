package dynamodbglobaltable


type DynamodbGlobalTableReplicasSseSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#kms_master_key_id DynamodbGlobalTable#kms_master_key_id}.
	KmsMasterKeyId *string `field:"required" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
}

