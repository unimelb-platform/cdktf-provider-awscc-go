package dataawscclexbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscclexbot/internal"
)

type DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference interface {
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
	InternalValue() *DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStruct
	SetInternalValue(val *DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStruct)
	Message() DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListMessageOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Variations() DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListVariationsList
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

// The jsii proxy struct for DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference
type jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) InternalValue() *DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStruct {
	var returns *DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStruct
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) Message() DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListMessageOutputReference {
	var returns DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListMessageOutputReference
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) Variations() DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListVariationsList {
	var returns DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListVariationsList
	_jsii_.Get(
		j,
		"variations",
		&returns,
	)
	return returns
}


func NewDataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccLexBot.DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference_Override(d DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccLexBot.DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference)SetInternalValue(val *DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStruct) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupsListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

