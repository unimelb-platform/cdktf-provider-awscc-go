package batchjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/batchjobdefinition/internal"
)

type BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference interface {
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
	Devices() BatchJobDefinitionContainerPropertiesLinuxParametersDevicesList
	DevicesInput() interface{}
	// Experimental.
	Fqn() *string
	InitProcessEnabled() interface{}
	SetInitProcessEnabled(val interface{})
	InitProcessEnabledInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxSwap() *float64
	SetMaxSwap(val *float64)
	MaxSwapInput() *float64
	SharedMemorySize() *float64
	SetSharedMemorySize(val *float64)
	SharedMemorySizeInput() *float64
	Swappiness() *float64
	SetSwappiness(val *float64)
	SwappinessInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tmpfs() BatchJobDefinitionContainerPropertiesLinuxParametersTmpfsList
	TmpfsInput() interface{}
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
	PutDevices(value interface{})
	PutTmpfs(value interface{})
	ResetDevices()
	ResetInitProcessEnabled()
	ResetMaxSwap()
	ResetSharedMemorySize()
	ResetSwappiness()
	ResetTmpfs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference
type jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) Devices() BatchJobDefinitionContainerPropertiesLinuxParametersDevicesList {
	var returns BatchJobDefinitionContainerPropertiesLinuxParametersDevicesList
	_jsii_.Get(
		j,
		"devices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) DevicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"devicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) InitProcessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initProcessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) InitProcessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initProcessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) MaxSwap() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSwap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) MaxSwapInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSwapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) SharedMemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharedMemorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) SharedMemorySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharedMemorySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) Swappiness() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swappiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) SwappinessInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swappinessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) Tmpfs() BatchJobDefinitionContainerPropertiesLinuxParametersTmpfsList {
	var returns BatchJobDefinitionContainerPropertiesLinuxParametersTmpfsList
	_jsii_.Get(
		j,
		"tmpfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) TmpfsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tmpfsInput",
		&returns,
	)
	return returns
}


func NewBatchJobDefinitionContainerPropertiesLinuxParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference {
	_init_.Initialize()

	if err := validateNewBatchJobDefinitionContainerPropertiesLinuxParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference{}

	_jsii_.Create(
		"awscc.batchJobDefinition.BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchJobDefinitionContainerPropertiesLinuxParametersOutputReference_Override(b BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.batchJobDefinition.BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference)SetInitProcessEnabled(val interface{}) {
	if err := j.validateSetInitProcessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initProcessEnabled",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference)SetMaxSwap(val *float64) {
	if err := j.validateSetMaxSwapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSwap",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference)SetSharedMemorySize(val *float64) {
	if err := j.validateSetSharedMemorySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedMemorySize",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference)SetSwappiness(val *float64) {
	if err := j.validateSetSwappinessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"swappiness",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) PutDevices(value interface{}) {
	if err := b.validatePutDevicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDevices",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) PutTmpfs(value interface{}) {
	if err := b.validatePutTmpfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTmpfs",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) ResetDevices() {
	_jsii_.InvokeVoid(
		b,
		"resetDevices",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) ResetInitProcessEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetInitProcessEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) ResetMaxSwap() {
	_jsii_.InvokeVoid(
		b,
		"resetMaxSwap",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) ResetSharedMemorySize() {
	_jsii_.InvokeVoid(
		b,
		"resetSharedMemorySize",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) ResetSwappiness() {
	_jsii_.InvokeVoid(
		b,
		"resetSwappiness",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) ResetTmpfs() {
	_jsii_.InvokeVoid(
		b,
		"resetTmpfs",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

