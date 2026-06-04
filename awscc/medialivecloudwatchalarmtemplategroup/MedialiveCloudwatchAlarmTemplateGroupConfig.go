package medialivecloudwatchalarmtemplategroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveCloudwatchAlarmTemplateGroupConfig struct {
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
	// A resource's name. Names must be unique within the scope of a resource type in a specific region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_cloudwatch_alarm_template_group#name MedialiveCloudwatchAlarmTemplateGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A resource's optional description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_cloudwatch_alarm_template_group#description MedialiveCloudwatchAlarmTemplateGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Represents the tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_cloudwatch_alarm_template_group#tags MedialiveCloudwatchAlarmTemplateGroup#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

