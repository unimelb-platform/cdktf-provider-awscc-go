package evidentlyproject


type EvidentlyProjectDataDelivery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_project#log_group EvidentlyProject#log_group}.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_project#s3 EvidentlyProject#s3}.
	S3 *EvidentlyProjectDataDeliveryS3 `field:"optional" json:"s3" yaml:"s3"`
}

