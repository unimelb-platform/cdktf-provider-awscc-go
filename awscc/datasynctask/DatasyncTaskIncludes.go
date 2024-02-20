package datasynctask


type DatasyncTaskIncludes struct {
	// The type of filter rule to apply. AWS DataSync only supports the SIMPLE_PATTERN rule type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#filter_type DatasyncTask#filter_type}
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
	// A single filter string that consists of the patterns to include or exclude. The patterns are delimited by "|".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#value DatasyncTask#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

