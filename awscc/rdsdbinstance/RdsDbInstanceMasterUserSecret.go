package rdsdbinstance


type RdsDbInstanceMasterUserSecret struct {
	// The AWS KMS key identifier that is used to encrypt the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#kms_key_id RdsDbInstance#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

