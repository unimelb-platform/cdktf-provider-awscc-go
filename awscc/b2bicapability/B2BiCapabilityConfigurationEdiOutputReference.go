package b2bicapability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/b2bicapability/internal"
)

type B2BiCapabilityConfigurationEdiOutputReference interface {
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
	InputLocation() B2BiCapabilityConfigurationEdiInputLocationOutputReference
	InputLocationInput() *B2BiCapabilityConfigurationEdiInputLocation
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OutputLocation() B2BiCapabilityConfigurationEdiOutputLocationOutputReference
	OutputLocationInput() *B2BiCapabilityConfigurationEdiOutputLocation
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransformerId() *string
	SetTransformerId(val *string)
	TransformerIdInput() *string
	Type() B2BiCapabilityConfigurationEdiTypeOutputReference
	TypeInput() *B2BiCapabilityConfigurationEdiType
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
	PutInputLocation(value *B2BiCapabilityConfigurationEdiInputLocation)
	PutOutputLocation(value *B2BiCapabilityConfigurationEdiOutputLocation)
	PutType(value *B2BiCapabilityConfigurationEdiType)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for B2BiCapabilityConfigurationEdiOutputReference
type jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) InputLocation() B2BiCapabilityConfigurationEdiInputLocationOutputReference {
	var returns B2BiCapabilityConfigurationEdiInputLocationOutputReference
	_jsii_.Get(
		j,
		"inputLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) InputLocationInput() *B2BiCapabilityConfigurationEdiInputLocation {
	var returns *B2BiCapabilityConfigurationEdiInputLocation
	_jsii_.Get(
		j,
		"inputLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) OutputLocation() B2BiCapabilityConfigurationEdiOutputLocationOutputReference {
	var returns B2BiCapabilityConfigurationEdiOutputLocationOutputReference
	_jsii_.Get(
		j,
		"outputLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) OutputLocationInput() *B2BiCapabilityConfigurationEdiOutputLocation {
	var returns *B2BiCapabilityConfigurationEdiOutputLocation
	_jsii_.Get(
		j,
		"outputLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) TransformerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) TransformerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) Type() B2BiCapabilityConfigurationEdiTypeOutputReference {
	var returns B2BiCapabilityConfigurationEdiTypeOutputReference
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) TypeInput() *B2BiCapabilityConfigurationEdiType {
	var returns *B2BiCapabilityConfigurationEdiType
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewB2BiCapabilityConfigurationEdiOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) B2BiCapabilityConfigurationEdiOutputReference {
	_init_.Initialize()

	if err := validateNewB2BiCapabilityConfigurationEdiOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference{}

	_jsii_.Create(
		"awscc.b2BiCapability.B2BiCapabilityConfigurationEdiOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewB2BiCapabilityConfigurationEdiOutputReference_Override(b B2BiCapabilityConfigurationEdiOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.b2BiCapability.B2BiCapabilityConfigurationEdiOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference)SetTransformerId(val *string) {
	if err := j.validateSetTransformerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transformerId",
		val,
	)
}

func (b *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) PutInputLocation(value *B2BiCapabilityConfigurationEdiInputLocation) {
	if err := b.validatePutInputLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putInputLocation",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) PutOutputLocation(value *B2BiCapabilityConfigurationEdiOutputLocation) {
	if err := b.validatePutOutputLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putOutputLocation",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) PutType(value *B2BiCapabilityConfigurationEdiType) {
	if err := b.validatePutTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putType",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_B2BiCapabilityConfigurationEdiOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

