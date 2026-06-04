package codeguruprofilerprofilinggroup


type CodeguruprofilerProfilingGroupAgentPermissions struct {
	// The principals for the agent permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/codeguruprofiler_profiling_group#principals CodeguruprofilerProfilingGroup#principals}
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
}

