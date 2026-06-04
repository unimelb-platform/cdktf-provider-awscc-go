package dmsreplicationconfig


type DmsReplicationConfigTags struct {
	// <p>Tag key.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/dms_replication_config#key DmsReplicationConfig#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// <p>Tag value.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/dms_replication_config#value DmsReplicationConfig#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

