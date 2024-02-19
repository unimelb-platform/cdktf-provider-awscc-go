package mediapackageoriginendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediapackageoriginendpoint/internal"
)

type MediapackageOriginEndpointCmafPackageEncryptionOutputReference interface {
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
	ConstantInitializationVector() *string
	SetConstantInitializationVector(val *string)
	ConstantInitializationVectorInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EncryptionMethod() *string
	SetEncryptionMethod(val *string)
	EncryptionMethodInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyRotationIntervalSeconds() *float64
	SetKeyRotationIntervalSeconds(val *float64)
	KeyRotationIntervalSecondsInput() *float64
	SpekeKeyProvider() MediapackageOriginEndpointCmafPackageEncryptionSpekeKeyProviderOutputReference
	SpekeKeyProviderInput() *MediapackageOriginEndpointCmafPackageEncryptionSpekeKeyProvider
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
	PutSpekeKeyProvider(value *MediapackageOriginEndpointCmafPackageEncryptionSpekeKeyProvider)
	ResetConstantInitializationVector()
	ResetEncryptionMethod()
	ResetKeyRotationIntervalSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediapackageOriginEndpointCmafPackageEncryptionOutputReference
type jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) ConstantInitializationVector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constantInitializationVector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) ConstantInitializationVectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constantInitializationVectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) EncryptionMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) EncryptionMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) KeyRotationIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyRotationIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) KeyRotationIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyRotationIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) SpekeKeyProvider() MediapackageOriginEndpointCmafPackageEncryptionSpekeKeyProviderOutputReference {
	var returns MediapackageOriginEndpointCmafPackageEncryptionSpekeKeyProviderOutputReference
	_jsii_.Get(
		j,
		"spekeKeyProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) SpekeKeyProviderInput() *MediapackageOriginEndpointCmafPackageEncryptionSpekeKeyProvider {
	var returns *MediapackageOriginEndpointCmafPackageEncryptionSpekeKeyProvider
	_jsii_.Get(
		j,
		"spekeKeyProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediapackageOriginEndpointCmafPackageEncryptionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediapackageOriginEndpointCmafPackageEncryptionOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackageOriginEndpointCmafPackageEncryptionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference{}

	_jsii_.Create(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpointCmafPackageEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediapackageOriginEndpointCmafPackageEncryptionOutputReference_Override(m MediapackageOriginEndpointCmafPackageEncryptionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpointCmafPackageEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference)SetConstantInitializationVector(val *string) {
	if err := j.validateSetConstantInitializationVectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"constantInitializationVector",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference)SetEncryptionMethod(val *string) {
	if err := j.validateSetEncryptionMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionMethod",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference)SetKeyRotationIntervalSeconds(val *float64) {
	if err := j.validateSetKeyRotationIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyRotationIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) PutSpekeKeyProvider(value *MediapackageOriginEndpointCmafPackageEncryptionSpekeKeyProvider) {
	if err := m.validatePutSpekeKeyProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSpekeKeyProvider",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) ResetConstantInitializationVector() {
	_jsii_.InvokeVoid(
		m,
		"resetConstantInitializationVector",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) ResetEncryptionMethod() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryptionMethod",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) ResetKeyRotationIntervalSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyRotationIntervalSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediapackageOriginEndpointCmafPackageEncryptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

