package dynamodbglobaltable


type DynamodbGlobalTableReplicasSseSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/dynamodb_global_table#kms_master_key_id DynamodbGlobalTable#kms_master_key_id}.
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
}

