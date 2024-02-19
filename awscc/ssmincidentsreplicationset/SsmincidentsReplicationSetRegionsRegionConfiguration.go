package ssmincidentsreplicationset


type SsmincidentsReplicationSetRegionsRegionConfiguration struct {
	// The ARN of the ReplicationSet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmincidents_replication_set#sse_kms_key_id SsmincidentsReplicationSet#sse_kms_key_id}
	SseKmsKeyId *string `field:"required" json:"sseKmsKeyId" yaml:"sseKmsKeyId"`
}

