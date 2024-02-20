package omicsrungroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OmicsRunGroupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_run_group#max_cpus OmicsRunGroup#max_cpus}.
	MaxCpus *float64 `field:"optional" json:"maxCpus" yaml:"maxCpus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_run_group#max_duration OmicsRunGroup#max_duration}.
	MaxDuration *float64 `field:"optional" json:"maxDuration" yaml:"maxDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_run_group#max_gpus OmicsRunGroup#max_gpus}.
	MaxGpus *float64 `field:"optional" json:"maxGpus" yaml:"maxGpus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_run_group#max_runs OmicsRunGroup#max_runs}.
	MaxRuns *float64 `field:"optional" json:"maxRuns" yaml:"maxRuns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_run_group#name OmicsRunGroup#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A map of resource tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_run_group#tags OmicsRunGroup#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

