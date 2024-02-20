package applicationautoscalingscalingpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationautoscalingScalingPolicyConfig struct {
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
	// The name of the scaling policy.
	//
	// Updates to the name of a target tracking scaling policy are not supported, unless you also update the metric used for scaling. To change only a target tracking scaling policy's name, first delete the policy by removing the existing AWS::ApplicationAutoScaling::ScalingPolicy resource from the template and updating the stack. Then, recreate the resource with the same settings and a different name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#policy_name ApplicationautoscalingScalingPolicy#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The scaling policy type.
	//
	// The following policy types are supported:
	//
	// TargetTrackingScaling Not supported for Amazon EMR
	//
	// StepScaling Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon Keyspaces, Amazon MSK, Amazon ElastiCache, or Neptune.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#policy_type ApplicationautoscalingScalingPolicy#policy_type}
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
	// The identifier of the resource associated with the scaling policy.
	//
	// This string consists of the resource type and unique identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#resource_id ApplicationautoscalingScalingPolicy#resource_id}
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// The scalable dimension. This string consists of the service namespace, resource type, and scaling property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#scalable_dimension ApplicationautoscalingScalingPolicy#scalable_dimension}
	ScalableDimension *string `field:"optional" json:"scalableDimension" yaml:"scalableDimension"`
	// The CloudFormation-generated ID of an Application Auto Scaling scalable target.
	//
	// For more information about the ID, see the Return Value section of the AWS::ApplicationAutoScaling::ScalableTarget resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#scaling_target_id ApplicationautoscalingScalingPolicy#scaling_target_id}
	ScalingTargetId *string `field:"optional" json:"scalingTargetId" yaml:"scalingTargetId"`
	// The namespace of the AWS service that provides the resource, or a custom-resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#service_namespace ApplicationautoscalingScalingPolicy#service_namespace}
	ServiceNamespace *string `field:"optional" json:"serviceNamespace" yaml:"serviceNamespace"`
	// A step scaling policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#step_scaling_policy_configuration ApplicationautoscalingScalingPolicy#step_scaling_policy_configuration}
	StepScalingPolicyConfiguration *ApplicationautoscalingScalingPolicyStepScalingPolicyConfiguration `field:"optional" json:"stepScalingPolicyConfiguration" yaml:"stepScalingPolicyConfiguration"`
	// A target tracking scaling policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#target_tracking_scaling_policy_configuration ApplicationautoscalingScalingPolicy#target_tracking_scaling_policy_configuration}
	TargetTrackingScalingPolicyConfiguration *ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfiguration `field:"optional" json:"targetTrackingScalingPolicyConfiguration" yaml:"targetTrackingScalingPolicyConfiguration"`
}

