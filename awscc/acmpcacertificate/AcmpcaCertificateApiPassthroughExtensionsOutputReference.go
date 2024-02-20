package acmpcacertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/acmpcacertificate/internal"
)

type AcmpcaCertificateApiPassthroughExtensionsOutputReference interface {
	cdktf.ComplexObject
	CertificatePolicies() AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesList
	CertificatePoliciesInput() interface{}
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
	CustomExtensions() AcmpcaCertificateApiPassthroughExtensionsCustomExtensionsList
	CustomExtensionsInput() interface{}
	ExtendedKeyUsage() AcmpcaCertificateApiPassthroughExtensionsExtendedKeyUsageList
	ExtendedKeyUsageInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyUsage() AcmpcaCertificateApiPassthroughExtensionsKeyUsageOutputReference
	KeyUsageInput() interface{}
	SubjectAlternativeNames() AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesList
	SubjectAlternativeNamesInput() interface{}
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
	PutCertificatePolicies(value interface{})
	PutCustomExtensions(value interface{})
	PutExtendedKeyUsage(value interface{})
	PutKeyUsage(value *AcmpcaCertificateApiPassthroughExtensionsKeyUsage)
	PutSubjectAlternativeNames(value interface{})
	ResetCertificatePolicies()
	ResetCustomExtensions()
	ResetExtendedKeyUsage()
	ResetKeyUsage()
	ResetSubjectAlternativeNames()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AcmpcaCertificateApiPassthroughExtensionsOutputReference
type jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) CertificatePolicies() AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesList {
	var returns AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesList
	_jsii_.Get(
		j,
		"certificatePolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) CertificatePoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificatePoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) CustomExtensions() AcmpcaCertificateApiPassthroughExtensionsCustomExtensionsList {
	var returns AcmpcaCertificateApiPassthroughExtensionsCustomExtensionsList
	_jsii_.Get(
		j,
		"customExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) CustomExtensionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customExtensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) ExtendedKeyUsage() AcmpcaCertificateApiPassthroughExtensionsExtendedKeyUsageList {
	var returns AcmpcaCertificateApiPassthroughExtensionsExtendedKeyUsageList
	_jsii_.Get(
		j,
		"extendedKeyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) ExtendedKeyUsageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extendedKeyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) KeyUsage() AcmpcaCertificateApiPassthroughExtensionsKeyUsageOutputReference {
	var returns AcmpcaCertificateApiPassthroughExtensionsKeyUsageOutputReference
	_jsii_.Get(
		j,
		"keyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) KeyUsageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) SubjectAlternativeNames() AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesList {
	var returns AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesList
	_jsii_.Get(
		j,
		"subjectAlternativeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) SubjectAlternativeNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subjectAlternativeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAcmpcaCertificateApiPassthroughExtensionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AcmpcaCertificateApiPassthroughExtensionsOutputReference {
	_init_.Initialize()

	if err := validateNewAcmpcaCertificateApiPassthroughExtensionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference{}

	_jsii_.Create(
		"awscc.acmpcaCertificate.AcmpcaCertificateApiPassthroughExtensionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAcmpcaCertificateApiPassthroughExtensionsOutputReference_Override(a AcmpcaCertificateApiPassthroughExtensionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.acmpcaCertificate.AcmpcaCertificateApiPassthroughExtensionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) PutCertificatePolicies(value interface{}) {
	if err := a.validatePutCertificatePoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCertificatePolicies",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) PutCustomExtensions(value interface{}) {
	if err := a.validatePutCustomExtensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomExtensions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) PutExtendedKeyUsage(value interface{}) {
	if err := a.validatePutExtendedKeyUsageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExtendedKeyUsage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) PutKeyUsage(value *AcmpcaCertificateApiPassthroughExtensionsKeyUsage) {
	if err := a.validatePutKeyUsageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKeyUsage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) PutSubjectAlternativeNames(value interface{}) {
	if err := a.validatePutSubjectAlternativeNamesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSubjectAlternativeNames",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) ResetCertificatePolicies() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificatePolicies",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) ResetCustomExtensions() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomExtensions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) ResetExtendedKeyUsage() {
	_jsii_.InvokeVoid(
		a,
		"resetExtendedKeyUsage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) ResetKeyUsage() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyUsage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) ResetSubjectAlternativeNames() {
	_jsii_.InvokeVoid(
		a,
		"resetSubjectAlternativeNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

