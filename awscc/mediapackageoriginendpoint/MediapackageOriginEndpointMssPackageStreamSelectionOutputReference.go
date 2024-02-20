package mediapackageoriginendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediapackageoriginendpoint/internal"
)

type MediapackageOriginEndpointMssPackageStreamSelectionOutputReference interface {
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
	MaxVideoBitsPerSecond() *float64
	SetMaxVideoBitsPerSecond(val *float64)
	MaxVideoBitsPerSecondInput() *float64
	MinVideoBitsPerSecond() *float64
	SetMinVideoBitsPerSecond(val *float64)
	MinVideoBitsPerSecondInput() *float64
	StreamOrder() *string
	SetStreamOrder(val *string)
	StreamOrderInput() *string
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
	ResetMaxVideoBitsPerSecond()
	ResetMinVideoBitsPerSecond()
	ResetStreamOrder()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediapackageOriginEndpointMssPackageStreamSelectionOutputReference
type jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) MaxVideoBitsPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVideoBitsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) MaxVideoBitsPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVideoBitsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) MinVideoBitsPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minVideoBitsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) MinVideoBitsPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minVideoBitsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) StreamOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) StreamOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediapackageOriginEndpointMssPackageStreamSelectionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediapackageOriginEndpointMssPackageStreamSelectionOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackageOriginEndpointMssPackageStreamSelectionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference{}

	_jsii_.Create(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpointMssPackageStreamSelectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediapackageOriginEndpointMssPackageStreamSelectionOutputReference_Override(m MediapackageOriginEndpointMssPackageStreamSelectionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpointMssPackageStreamSelectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference)SetMaxVideoBitsPerSecond(val *float64) {
	if err := j.validateSetMaxVideoBitsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxVideoBitsPerSecond",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference)SetMinVideoBitsPerSecond(val *float64) {
	if err := j.validateSetMinVideoBitsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minVideoBitsPerSecond",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference)SetStreamOrder(val *string) {
	if err := j.validateSetStreamOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamOrder",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) ResetMaxVideoBitsPerSecond() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxVideoBitsPerSecond",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) ResetMinVideoBitsPerSecond() {
	_jsii_.InvokeVoid(
		m,
		"resetMinVideoBitsPerSecond",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) ResetStreamOrder() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamOrder",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointMssPackageStreamSelectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

