package ioteventsdetectormodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ioteventsdetectormodel/internal"
)

type IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference interface {
	cdktf.ComplexObject
	AssetId() *string
	SetAssetId(val *string)
	AssetIdInput() *string
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
	EntryId() *string
	SetEntryId(val *string)
	EntryIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PropertyAlias() *string
	SetPropertyAlias(val *string)
	PropertyAliasInput() *string
	PropertyId() *string
	SetPropertyId(val *string)
	PropertyIdInput() *string
	PropertyValue() IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWisePropertyValueOutputReference
	PropertyValueInput() *IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWisePropertyValue
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
	PutPropertyValue(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWisePropertyValue)
	ResetAssetId()
	ResetEntryId()
	ResetPropertyAlias()
	ResetPropertyId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference
type jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) AssetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) AssetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) EntryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) EntryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) PropertyAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) PropertyAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) PropertyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) PropertyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) PropertyValue() IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWisePropertyValueOutputReference {
	var returns IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWisePropertyValueOutputReference
	_jsii_.Get(
		j,
		"propertyValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) PropertyValueInput() *IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWisePropertyValue {
	var returns *IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWisePropertyValue
	_jsii_.Get(
		j,
		"propertyValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference {
	_init_.Initialize()

	if err := validateNewIoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference{}

	_jsii_.Create(
		"awscc.ioteventsDetectorModel.IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference_Override(i IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ioteventsDetectorModel.IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference)SetAssetId(val *string) {
	if err := j.validateSetAssetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetId",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference)SetEntryId(val *string) {
	if err := j.validateSetEntryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryId",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference)SetPropertyAlias(val *string) {
	if err := j.validateSetPropertyAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyAlias",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference)SetPropertyId(val *string) {
	if err := j.validateSetPropertyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyId",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) PutPropertyValue(value *IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWisePropertyValue) {
	if err := i.validatePutPropertyValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPropertyValue",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) ResetAssetId() {
	_jsii_.InvokeVoid(
		i,
		"resetAssetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) ResetEntryId() {
	_jsii_.InvokeVoid(
		i,
		"resetEntryId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) ResetPropertyAlias() {
	_jsii_.InvokeVoid(
		i,
		"resetPropertyAlias",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) ResetPropertyId() {
	_jsii_.InvokeVoid(
		i,
		"resetPropertyId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWiseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

