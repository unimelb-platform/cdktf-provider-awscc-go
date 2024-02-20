package route53recoveryreadinessresourceset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53RecoveryreadinessResourceSetConfig struct {
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
	// A list of resource objects in the resource set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoveryreadiness_resource_set#resources Route53RecoveryreadinessResourceSet#resources}
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
	// The resource type of the resources in the resource set.
	//
	// Enter one of the following values for resource type:
	//
	// AWS: :AutoScaling: :AutoScalingGroup, AWS: :CloudWatch: :Alarm, AWS: :EC2: :CustomerGateway, AWS: :DynamoDB: :Table, AWS: :EC2: :Volume, AWS: :ElasticLoadBalancing: :LoadBalancer, AWS: :ElasticLoadBalancingV2: :LoadBalancer, AWS: :MSK: :Cluster, AWS: :RDS: :DBCluster, AWS: :Route53: :HealthCheck, AWS: :SQS: :Queue, AWS: :SNS: :Topic, AWS: :SNS: :Subscription, AWS: :EC2: :VPC, AWS: :EC2: :VPNConnection, AWS: :EC2: :VPNGateway, AWS::Route53RecoveryReadiness::DNSTargetResource
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoveryreadiness_resource_set#resource_set_type Route53RecoveryreadinessResourceSet#resource_set_type}
	ResourceSetType *string `field:"required" json:"resourceSetType" yaml:"resourceSetType"`
	// The name of the resource set to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoveryreadiness_resource_set#resource_set_name Route53RecoveryreadinessResourceSet#resource_set_name}
	ResourceSetName *string `field:"optional" json:"resourceSetName" yaml:"resourceSetName"`
	// A tag to associate with the parameters for a resource set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoveryreadiness_resource_set#tags Route53RecoveryreadinessResourceSet#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

