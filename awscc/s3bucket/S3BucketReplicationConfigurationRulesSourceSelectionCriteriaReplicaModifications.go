package s3bucket


type S3BucketReplicationConfigurationRulesSourceSelectionCriteriaReplicaModifications struct {
	// Specifies whether Amazon S3 replicates modifications on replicas.   *Allowed values*: ``Enabled`` | ``Disabled``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3_bucket#status S3Bucket#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

