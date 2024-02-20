package databrewschedule


type DatabrewScheduleTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_schedule#key DatabrewSchedule#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_schedule#value DatabrewSchedule#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

