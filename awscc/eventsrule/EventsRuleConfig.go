package eventsrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventsRuleConfig struct {
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
	// The description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#description EventsRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name or ARN of the event bus associated with the rule.
	//
	// If you omit this, the default event bus is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#event_bus_name EventsRule#event_bus_name}
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// The event pattern of the rule.
	//
	// For more information, see Events and Event Patterns in the Amazon EventBridge User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#event_pattern EventsRule#event_pattern}
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// The name of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#name EventsRule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the role that is used for target invocation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#role_arn EventsRule#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The scheduling expression.
	//
	// For example, "cron(0 20 * * ? *)", "rate(5 minutes)". For more information, see Creating an Amazon EventBridge rule that runs on a schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#schedule_expression EventsRule#schedule_expression}
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	// The state of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#state EventsRule#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// Adds the specified targets to the specified rule, or updates the targets if they are already associated with the rule.
	//
	// Targets are the resources that are invoked when a rule is triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#targets EventsRule#targets}
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}

