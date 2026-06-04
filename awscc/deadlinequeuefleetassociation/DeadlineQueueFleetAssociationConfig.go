package deadlinequeuefleetassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeadlineQueueFleetAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/deadline_queue_fleet_association#farm_id DeadlineQueueFleetAssociation#farm_id}.
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/deadline_queue_fleet_association#fleet_id DeadlineQueueFleetAssociation#fleet_id}.
	FleetId *string `field:"required" json:"fleetId" yaml:"fleetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/deadline_queue_fleet_association#queue_id DeadlineQueueFleetAssociation#queue_id}.
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
}

