package mediapackagepackagingconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediapackagepackagingconfiguration/internal"
)

type MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference interface {
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
	EncryptionContractConfiguration() MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderEncryptionContractConfigurationOutputReference
	EncryptionContractConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProvider
	SetInternalValue(val *MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProvider)
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
	PutEncryptionContractConfiguration(value *MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderEncryptionContractConfiguration)
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

// The jsii proxy struct for MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference
type jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) EncryptionContractConfiguration() MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderEncryptionContractConfigurationOutputReference {
	var returns MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderEncryptionContractConfigurationOutputReference
	_jsii_.Get(
		j,
		"encryptionContractConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) EncryptionContractConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionContractConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) InternalValue() *MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProvider {
	var returns *MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProvider
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) SystemIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) SystemIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewMediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference {
	_init_.Initialize()

	if err := validateNewMediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference{}

	_jsii_.Create(
		"awscc.mediapackagePackagingConfiguration.MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference_Override(m MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediapackagePackagingConfiguration.MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference)SetInternalValue(val *MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProvider) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference)SetSystemIds(val *[]*string) {
	if err := j.validateSetSystemIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemIds",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) PutEncryptionContractConfiguration(value *MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderEncryptionContractConfiguration) {
	if err := m.validatePutEncryptionContractConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEncryptionContractConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) ResetEncryptionContractConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryptionContractConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediapackagePackagingConfigurationDashPackageEncryptionSpekeKeyProviderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

