package iotmitigationaction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/iotmitigationaction/internal"
)

type IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OverrideDynamicGroups() interface{}
	SetOverrideDynamicGroups(val interface{})
	OverrideDynamicGroupsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThingGroupNames() *[]*string
	SetThingGroupNames(val *[]*string)
	ThingGroupNamesInput() *[]*string
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
	ResetOverrideDynamicGroups()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference
type jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) OverrideDynamicGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideDynamicGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) OverrideDynamicGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideDynamicGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) ThingGroupNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"thingGroupNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) ThingGroupNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"thingGroupNamesInput",
		&returns,
	)
	return returns
}


func NewIotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference {
	_init_.Initialize()

	if err := validateNewIotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference{}

	_jsii_.Create(
		"awscc.iotMitigationAction.IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference_Override(i IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.iotMitigationAction.IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference)SetOverrideDynamicGroups(val interface{}) {
	if err := j.validateSetOverrideDynamicGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideDynamicGroups",
		val,
	)
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference)SetThingGroupNames(val *[]*string) {
	if err := j.validateSetThingGroupNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thingGroupNames",
		val,
	)
}

func (i *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) ResetOverrideDynamicGroups() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideDynamicGroups",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotMitigationActionActionParamsAddThingsToThingGroupParamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

