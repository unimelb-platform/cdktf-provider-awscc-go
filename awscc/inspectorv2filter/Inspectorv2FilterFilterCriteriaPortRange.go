package inspectorv2filter


type Inspectorv2FilterFilterCriteriaPortRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#begin_inclusive Inspectorv2Filter#begin_inclusive}.
	BeginInclusive *float64 `field:"optional" json:"beginInclusive" yaml:"beginInclusive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#end_inclusive Inspectorv2Filter#end_inclusive}.
	EndInclusive *float64 `field:"optional" json:"endInclusive" yaml:"endInclusive"`
}

