package ioteventsdetectormodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ioteventsdetectormodel/internal"
)

type IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference interface {
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
	ContentExpression() *string
	SetContentExpression(val *string)
	ContentExpressionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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

// The jsii proxy struct for IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference
type jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) ContentExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) ContentExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference {
	_init_.Initialize()

	if err := validateNewIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference{}

	_jsii_.Create(
		"awscc.ioteventsDetectorModel.IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference_Override(i IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ioteventsDetectorModel.IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference)SetContentExpression(val *string) {
	if err := j.validateSetContentExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentExpression",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotTopicPublishPayloadOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
