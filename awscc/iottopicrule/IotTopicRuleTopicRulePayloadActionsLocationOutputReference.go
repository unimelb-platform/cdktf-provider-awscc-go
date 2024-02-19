package iottopicrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/iottopicrule/internal"
)

type IotTopicRuleTopicRulePayloadActionsLocationOutputReference interface {
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
	DeviceId() *string
	SetDeviceId(val *string)
	DeviceIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Latitude() *string
	SetLatitude(val *string)
	LatitudeInput() *string
	Longitude() *string
	SetLongitude(val *string)
	LongitudeInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timestamp() IotTopicRuleTopicRulePayloadActionsLocationTimestampOutputReference
	TimestampInput() interface{}
	TrackerName() *string
	SetTrackerName(val *string)
	TrackerNameInput() *string
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
	PutTimestamp(value *IotTopicRuleTopicRulePayloadActionsLocationTimestamp)
	ResetTimestamp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotTopicRuleTopicRulePayloadActionsLocationOutputReference
type jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) DeviceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) DeviceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) Latitude() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) LatitudeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) Longitude() *string {
	var returns *string
	_jsii_.Get(
		j,
		"longitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) LongitudeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"longitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) Timestamp() IotTopicRuleTopicRulePayloadActionsLocationTimestampOutputReference {
	var returns IotTopicRuleTopicRulePayloadActionsLocationTimestampOutputReference
	_jsii_.Get(
		j,
		"timestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) TimestampInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) TrackerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) TrackerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackerNameInput",
		&returns,
	)
	return returns
}


func NewIotTopicRuleTopicRulePayloadActionsLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IotTopicRuleTopicRulePayloadActionsLocationOutputReference {
	_init_.Initialize()

	if err := validateNewIotTopicRuleTopicRulePayloadActionsLocationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference{}

	_jsii_.Create(
		"awscc.iotTopicRule.IotTopicRuleTopicRulePayloadActionsLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotTopicRuleTopicRulePayloadActionsLocationOutputReference_Override(i IotTopicRuleTopicRulePayloadActionsLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.iotTopicRule.IotTopicRuleTopicRulePayloadActionsLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference)SetDeviceId(val *string) {
	if err := j.validateSetDeviceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceId",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference)SetLatitude(val *string) {
	if err := j.validateSetLatitudeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latitude",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference)SetLongitude(val *string) {
	if err := j.validateSetLongitudeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longitude",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference)SetTrackerName(val *string) {
	if err := j.validateSetTrackerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackerName",
		val,
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) PutTimestamp(value *IotTopicRuleTopicRulePayloadActionsLocationTimestamp) {
	if err := i.validatePutTimestampParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimestamp",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) ResetTimestamp() {
	_jsii_.InvokeVoid(
		i,
		"resetTimestamp",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleTopicRulePayloadActionsLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

