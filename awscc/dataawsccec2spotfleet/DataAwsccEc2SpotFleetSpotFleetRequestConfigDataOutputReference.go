package dataawsccec2spotfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccec2spotfleet/internal"
)

type DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference interface {
	cdktf.ComplexObject
	AllocationStrategy() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Context() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExcessCapacityTerminationPolicy() *string
	// Experimental.
	Fqn() *string
	IamFleetRole() *string
	InstanceInterruptionBehavior() *string
	InstancePoolsToUseCount() *float64
	InternalValue() *DataAwsccEc2SpotFleetSpotFleetRequestConfigData
	SetInternalValue(val *DataAwsccEc2SpotFleetSpotFleetRequestConfigData)
	LaunchSpecifications() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsList
	LaunchTemplateConfigs() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsList
	LoadBalancersConfig() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLoadBalancersConfigOutputReference
	OnDemandAllocationStrategy() *string
	OnDemandMaxTotalPrice() *string
	OnDemandTargetCapacity() *float64
	ReplaceUnhealthyInstances() cdktf.IResolvable
	SpotMaintenanceStrategies() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataSpotMaintenanceStrategiesOutputReference
	SpotMaxTotalPrice() *string
	SpotPrice() *string
	TagSpecifications() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataTagSpecificationsList
	TargetCapacity() *float64
	TargetCapacityUnitType() *string
	TerminateInstancesWithExpiration() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	ValidFrom() *string
	ValidUntil() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference
type jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) AllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) ExcessCapacityTerminationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excessCapacityTerminationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) IamFleetRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamFleetRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) InstanceInterruptionBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInterruptionBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) InstancePoolsToUseCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancePoolsToUseCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) InternalValue() *DataAwsccEc2SpotFleetSpotFleetRequestConfigData {
	var returns *DataAwsccEc2SpotFleetSpotFleetRequestConfigData
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) LaunchSpecifications() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsList {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsList
	_jsii_.Get(
		j,
		"launchSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) LaunchTemplateConfigs() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsList {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsList
	_jsii_.Get(
		j,
		"launchTemplateConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) LoadBalancersConfig() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLoadBalancersConfigOutputReference {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLoadBalancersConfigOutputReference
	_jsii_.Get(
		j,
		"loadBalancersConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) OnDemandAllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandAllocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) OnDemandMaxTotalPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandMaxTotalPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) OnDemandTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) ReplaceUnhealthyInstances() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"replaceUnhealthyInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) SpotMaintenanceStrategies() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataSpotMaintenanceStrategiesOutputReference {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataSpotMaintenanceStrategiesOutputReference
	_jsii_.Get(
		j,
		"spotMaintenanceStrategies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) SpotMaxTotalPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotMaxTotalPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) SpotPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) TagSpecifications() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataTagSpecificationsList {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataTagSpecificationsList
	_jsii_.Get(
		j,
		"tagSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) TargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) TargetCapacityUnitType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCapacityUnitType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) TerminateInstancesWithExpiration() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"terminateInstancesWithExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) ValidFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) ValidUntil() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntil",
		&returns,
	)
	return returns
}


func NewDataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccEc2SpotFleet.DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference_Override(d DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccEc2SpotFleet.DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference)SetInternalValue(val *DataAwsccEc2SpotFleetSpotFleetRequestConfigData) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

