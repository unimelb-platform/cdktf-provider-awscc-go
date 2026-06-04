package bedrockagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/bedrockagent/internal"
)

type BedrockAgentActionGroupsOutputReference interface {
	cdktf.ComplexObject
	ActionGroupExecutor() BedrockAgentActionGroupsActionGroupExecutorOutputReference
	ActionGroupExecutorInput() interface{}
	ActionGroupName() *string
	SetActionGroupName(val *string)
	ActionGroupNameInput() *string
	ActionGroupState() *string
	SetActionGroupState(val *string)
	ActionGroupStateInput() *string
	ApiSchema() BedrockAgentActionGroupsApiSchemaOutputReference
	ApiSchemaInput() interface{}
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	FunctionSchema() BedrockAgentActionGroupsFunctionSchemaOutputReference
	FunctionSchemaInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ParentActionGroupSignature() *string
	SetParentActionGroupSignature(val *string)
	ParentActionGroupSignatureInput() *string
	SkipResourceInUseCheckOnDelete() interface{}
	SetSkipResourceInUseCheckOnDelete(val interface{})
	SkipResourceInUseCheckOnDeleteInput() interface{}
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
	PutActionGroupExecutor(value *BedrockAgentActionGroupsActionGroupExecutor)
	PutApiSchema(value *BedrockAgentActionGroupsApiSchema)
	PutFunctionSchema(value *BedrockAgentActionGroupsFunctionSchema)
	ResetActionGroupExecutor()
	ResetActionGroupName()
	ResetActionGroupState()
	ResetApiSchema()
	ResetDescription()
	ResetFunctionSchema()
	ResetParentActionGroupSignature()
	ResetSkipResourceInUseCheckOnDelete()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockAgentActionGroupsOutputReference
type jsiiProxy_BedrockAgentActionGroupsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) ActionGroupExecutor() BedrockAgentActionGroupsActionGroupExecutorOutputReference {
	var returns BedrockAgentActionGroupsActionGroupExecutorOutputReference
	_jsii_.Get(
		j,
		"actionGroupExecutor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) ActionGroupExecutorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionGroupExecutorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) ActionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) ActionGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) ActionGroupState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) ActionGroupStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) ApiSchema() BedrockAgentActionGroupsApiSchemaOutputReference {
	var returns BedrockAgentActionGroupsApiSchemaOutputReference
	_jsii_.Get(
		j,
		"apiSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) ApiSchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) FunctionSchema() BedrockAgentActionGroupsFunctionSchemaOutputReference {
	var returns BedrockAgentActionGroupsFunctionSchemaOutputReference
	_jsii_.Get(
		j,
		"functionSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) FunctionSchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) ParentActionGroupSignature() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentActionGroupSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) ParentActionGroupSignatureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentActionGroupSignatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) SkipResourceInUseCheckOnDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipResourceInUseCheckOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) SkipResourceInUseCheckOnDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipResourceInUseCheckOnDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockAgentActionGroupsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockAgentActionGroupsOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockAgentActionGroupsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockAgentActionGroupsOutputReference{}

	_jsii_.Create(
		"awscc.bedrockAgent.BedrockAgentActionGroupsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockAgentActionGroupsOutputReference_Override(b BedrockAgentActionGroupsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.bedrockAgent.BedrockAgentActionGroupsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference)SetActionGroupName(val *string) {
	if err := j.validateSetActionGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionGroupName",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference)SetActionGroupState(val *string) {
	if err := j.validateSetActionGroupStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionGroupState",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference)SetParentActionGroupSignature(val *string) {
	if err := j.validateSetParentActionGroupSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentActionGroupSignature",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference)SetSkipResourceInUseCheckOnDelete(val interface{}) {
	if err := j.validateSetSkipResourceInUseCheckOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipResourceInUseCheckOnDelete",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentActionGroupsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) PutActionGroupExecutor(value *BedrockAgentActionGroupsActionGroupExecutor) {
	if err := b.validatePutActionGroupExecutorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putActionGroupExecutor",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) PutApiSchema(value *BedrockAgentActionGroupsApiSchema) {
	if err := b.validatePutApiSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putApiSchema",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) PutFunctionSchema(value *BedrockAgentActionGroupsFunctionSchema) {
	if err := b.validatePutFunctionSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putFunctionSchema",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) ResetActionGroupExecutor() {
	_jsii_.InvokeVoid(
		b,
		"resetActionGroupExecutor",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) ResetActionGroupName() {
	_jsii_.InvokeVoid(
		b,
		"resetActionGroupName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) ResetActionGroupState() {
	_jsii_.InvokeVoid(
		b,
		"resetActionGroupState",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) ResetApiSchema() {
	_jsii_.InvokeVoid(
		b,
		"resetApiSchema",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) ResetFunctionSchema() {
	_jsii_.InvokeVoid(
		b,
		"resetFunctionSchema",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) ResetParentActionGroupSignature() {
	_jsii_.InvokeVoid(
		b,
		"resetParentActionGroupSignature",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) ResetSkipResourceInUseCheckOnDelete() {
	_jsii_.InvokeVoid(
		b,
		"resetSkipResourceInUseCheckOnDelete",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentActionGroupsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

