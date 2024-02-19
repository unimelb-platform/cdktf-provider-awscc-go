package mediapackagepackagingconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediapackagepackagingconfiguration/internal"
)

type MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference interface {
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

// The jsii proxy struct for MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference
type jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) MaxVideoBitsPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVideoBitsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) MaxVideoBitsPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVideoBitsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) MinVideoBitsPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minVideoBitsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) MinVideoBitsPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minVideoBitsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) StreamOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) StreamOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference{}

	_jsii_.Create(
		"awscc.mediapackagePackagingConfiguration.MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference_Override(m MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediapackagePackagingConfiguration.MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference)SetMaxVideoBitsPerSecond(val *float64) {
	if err := j.validateSetMaxVideoBitsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxVideoBitsPerSecond",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference)SetMinVideoBitsPerSecond(val *float64) {
	if err := j.validateSetMinVideoBitsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minVideoBitsPerSecond",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference)SetStreamOrder(val *string) {
	if err := j.validateSetStreamOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamOrder",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) ResetMaxVideoBitsPerSecond() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxVideoBitsPerSecond",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) ResetMinVideoBitsPerSecond() {
	_jsii_.InvokeVoid(
		m,
		"resetMinVideoBitsPerSecond",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) ResetStreamOrder() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamOrder",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationHlsPackageHlsManifestsStreamSelectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

