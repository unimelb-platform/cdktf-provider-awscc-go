package mediapackageoriginendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediapackageoriginendpoint/internal"
)

type MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference interface {
	cdktf.ComplexObject
	CertificateArn() *string
	SetCertificateArn(val *string)
	CertificateArnInput() *string
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
	EncryptionContractConfiguration() MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderEncryptionContractConfigurationOutputReference
	EncryptionContractConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProvider
	SetInternalValue(val *MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProvider)
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	SystemIds() *[]*string
	SetSystemIds(val *[]*string)
	SystemIdsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	PutEncryptionContractConfiguration(value *MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderEncryptionContractConfiguration)
	ResetCertificateArn()
	ResetEncryptionContractConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference
type jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) EncryptionContractConfiguration() MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderEncryptionContractConfigurationOutputReference {
	var returns MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderEncryptionContractConfigurationOutputReference
	_jsii_.Get(
		j,
		"encryptionContractConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) EncryptionContractConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionContractConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) InternalValue() *MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProvider {
	var returns *MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProvider
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) SystemIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) SystemIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewMediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference{}

	_jsii_.Create(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference_Override(m MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediapackageOriginEndpoint.MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference)SetCertificateArn(val *string) {
	if err := j.validateSetCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference)SetInternalValue(val *MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProvider) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference)SetSystemIds(val *[]*string) {
	if err := j.validateSetSystemIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemIds",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) PutEncryptionContractConfiguration(value *MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderEncryptionContractConfiguration) {
	if err := m.validatePutEncryptionContractConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEncryptionContractConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		m,
		"resetCertificateArn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) ResetEncryptionContractConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryptionContractConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProviderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

