package dataawscccassandratable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscccassandratable/internal"
)

type DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisableScaleIn() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfiguration
	SetInternalValue(val *DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfiguration)
	ScaleInCooldown() *float64
	ScaleOutCooldown() *float64
	TargetValue() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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

// The jsii proxy struct for DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference
type jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) DisableScaleIn() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) InternalValue() *DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfiguration {
	var returns *DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) ScaleInCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) ScaleOutCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) TargetValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccCassandraTable.DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference_Override(d DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccCassandraTable.DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference)SetInternalValue(val *DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

