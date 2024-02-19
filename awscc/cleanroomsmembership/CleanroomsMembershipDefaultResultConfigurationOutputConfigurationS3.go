package cleanroomsmembership


type CleanroomsMembershipDefaultResultConfigurationOutputConfigurationS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_membership#bucket CleanroomsMembership#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_membership#result_format CleanroomsMembership#result_format}.
	ResultFormat *string `field:"required" json:"resultFormat" yaml:"resultFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_membership#key_prefix CleanroomsMembership#key_prefix}.
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

