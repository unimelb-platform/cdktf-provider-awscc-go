package customerprofileseventstream

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomerprofilesEventStreamConfig struct {
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
	// The unique name of the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_event_stream#domain_name CustomerprofilesEventStream#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The name of the event stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_event_stream#event_stream_name CustomerprofilesEventStream#event_stream_name}
	EventStreamName *string `field:"required" json:"eventStreamName" yaml:"eventStreamName"`
	// The StreamARN of the destination to deliver profile events to. For example, arn:aws:kinesis:region:account-id:stream/stream-name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_event_stream#uri CustomerprofilesEventStream#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// The tags used to organize, track, or control access for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_event_stream#tags CustomerprofilesEventStream#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

