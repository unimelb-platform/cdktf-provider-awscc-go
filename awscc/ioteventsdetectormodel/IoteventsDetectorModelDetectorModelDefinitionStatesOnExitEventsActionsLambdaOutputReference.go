package ioteventsdetectormodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ioteventsdetectormodel/internal"
)

type IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference interface {
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
	// Experimental.
	Fqn() *string
	FunctionArn() *string
	SetFunctionArn(val *string)
	FunctionArnInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Payload() IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaPayloadOutputReference
	PayloadInput() interface{}
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
	PutPayload(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaPayload)
	ResetPayload()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference
type jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) FunctionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) Payload() IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaPayloadOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaPayloadOutputReference
	_jsii_.Get(
		j,
		"payload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) PayloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"payloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference {
	_init_.Initialize()

	if err := validateNewIoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference{}

	_jsii_.Create(
		"awscc.ioteventsDetectorModel.IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference_Override(i IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ioteventsDetectorModel.IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference)SetFunctionArn(val *string) {
	if err := j.validateSetFunctionArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionArn",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) PutPayload(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaPayload) {
	if err := i.validatePutPayloadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPayload",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) ResetPayload() {
	_jsii_.InvokeVoid(
		i,
		"resetPayload",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsLambdaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

