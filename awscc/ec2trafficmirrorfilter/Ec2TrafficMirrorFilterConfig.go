package ec2trafficmirrorfilter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2TrafficMirrorFilterConfig struct {
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
	// The description of a traffic mirror filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_traffic_mirror_filter#description Ec2TrafficMirrorFilter#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The network service that is associated with the traffic mirror filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_traffic_mirror_filter#network_services Ec2TrafficMirrorFilter#network_services}
	NetworkServices *[]*string `field:"optional" json:"networkServices" yaml:"networkServices"`
	// The tags for a traffic mirror filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_traffic_mirror_filter#tags Ec2TrafficMirrorFilter#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

