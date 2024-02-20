package ioteventsdetectormodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ioteventsdetectormodel/internal"
)

type IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference interface {
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
	InternalValue() *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValue
	SetInternalValue(val *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValue)
	Quality() *string
	SetQuality(val *string)
	QualityInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timestamp() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueTimestampOutputReference
	TimestampInput() interface{}
	Value() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueValueOutputReference
	ValueInput() *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueValue
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
	PutTimestamp(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueTimestamp)
	PutValue(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueValue)
	ResetQuality()
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

// The jsii proxy struct for IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference
type jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) InternalValue() *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValue {
	var returns *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) Quality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) QualityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) Timestamp() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueTimestampOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueTimestampOutputReference
	_jsii_.Get(
		j,
		"timestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) TimestampInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) Value() IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueValueOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueValueOutputReference
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) ValueInput() *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueValue {
	var returns *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueValue
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference {
	_init_.Initialize()

	if err := validateNewIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference{}

	_jsii_.Create(
		"awscc.ioteventsDetectorModel.IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference_Override(i IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ioteventsDetectorModel.IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference)SetInternalValue(val *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValue) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference)SetQuality(val *string) {
	if err := j.validateSetQualityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quality",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) PutTimestamp(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueTimestamp) {
	if err := i.validatePutTimestampParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimestamp",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) PutValue(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueValue) {
	if err := i.validatePutValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putValue",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) ResetQuality() {
	_jsii_.InvokeVoid(
		i,
		"resetQuality",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) ResetTimestamp() {
	_jsii_.InvokeVoid(
		i,
		"resetTimestamp",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

