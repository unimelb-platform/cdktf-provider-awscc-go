package ioteventsalarmmodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ioteventsalarmmodel/internal"
)

type IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference interface {
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
	HashKeyField() *string
	SetHashKeyField(val *string)
	HashKeyFieldInput() *string
	HashKeyType() *string
	SetHashKeyType(val *string)
	HashKeyTypeInput() *string
	HashKeyValue() *string
	SetHashKeyValue(val *string)
	HashKeyValueInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Operation() *string
	SetOperation(val *string)
	OperationInput() *string
	Payload() IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbPayloadOutputReference
	PayloadField() *string
	SetPayloadField(val *string)
	PayloadFieldInput() *string
	PayloadInput() interface{}
	RangeKeyField() *string
	SetRangeKeyField(val *string)
	RangeKeyFieldInput() *string
	RangeKeyType() *string
	SetRangeKeyType(val *string)
	RangeKeyTypeInput() *string
	RangeKeyValue() *string
	SetRangeKeyValue(val *string)
	RangeKeyValueInput() *string
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
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
	PutPayload(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbPayload)
	ResetHashKeyType()
	ResetOperation()
	ResetPayload()
	ResetPayloadField()
	ResetRangeKeyField()
	ResetRangeKeyType()
	ResetRangeKeyValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference
type jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) HashKeyField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) HashKeyFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) HashKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) HashKeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) HashKeyValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) HashKeyValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) Operation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) OperationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) Payload() IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbPayloadOutputReference {
	var returns IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbPayloadOutputReference
	_jsii_.Get(
		j,
		"payload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) PayloadField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) PayloadFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) PayloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"payloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) RangeKeyField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) RangeKeyFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) RangeKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) RangeKeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) RangeKeyValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) RangeKeyValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference {
	_init_.Initialize()

	if err := validateNewIoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference{}

	_jsii_.Create(
		"awscc.ioteventsAlarmModel.IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference_Override(i IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ioteventsAlarmModel.IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference)SetHashKeyField(val *string) {
	if err := j.validateSetHashKeyFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashKeyField",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference)SetHashKeyType(val *string) {
	if err := j.validateSetHashKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashKeyType",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference)SetHashKeyValue(val *string) {
	if err := j.validateSetHashKeyValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashKeyValue",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference)SetOperation(val *string) {
	if err := j.validateSetOperationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operation",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference)SetPayloadField(val *string) {
	if err := j.validateSetPayloadFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payloadField",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference)SetRangeKeyField(val *string) {
	if err := j.validateSetRangeKeyFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rangeKeyField",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference)SetRangeKeyType(val *string) {
	if err := j.validateSetRangeKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rangeKeyType",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference)SetRangeKeyValue(val *string) {
	if err := j.validateSetRangeKeyValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rangeKeyValue",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) PutPayload(value *IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbPayload) {
	if err := i.validatePutPayloadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPayload",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) ResetHashKeyType() {
	_jsii_.InvokeVoid(
		i,
		"resetHashKeyType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) ResetOperation() {
	_jsii_.InvokeVoid(
		i,
		"resetOperation",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) ResetPayload() {
	_jsii_.InvokeVoid(
		i,
		"resetPayload",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) ResetPayloadField() {
	_jsii_.InvokeVoid(
		i,
		"resetPayloadField",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) ResetRangeKeyField() {
	_jsii_.InvokeVoid(
		i,
		"resetRangeKeyField",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) ResetRangeKeyType() {
	_jsii_.InvokeVoid(
		i,
		"resetRangeKeyType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) ResetRangeKeyValue() {
	_jsii_.InvokeVoid(
		i,
		"resetRangeKeyValue",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

