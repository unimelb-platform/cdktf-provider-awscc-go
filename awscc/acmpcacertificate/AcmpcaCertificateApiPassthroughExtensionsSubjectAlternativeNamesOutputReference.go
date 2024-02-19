package acmpcacertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/acmpcacertificate/internal"
)

type AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference interface {
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
	DirectoryName() AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameOutputReference
	DirectoryNameInput() interface{}
	DnsName() *string
	SetDnsName(val *string)
	DnsNameInput() *string
	EdiPartyName() AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesEdiPartyNameOutputReference
	EdiPartyNameInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	OtherName() AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOtherNameOutputReference
	OtherNameInput() interface{}
	RegisteredId() *string
	SetRegisteredId(val *string)
	RegisteredIdInput() *string
	Rfc822Name() *string
	SetRfc822Name(val *string)
	Rfc822NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UniformResourceIdentifier() *string
	SetUniformResourceIdentifier(val *string)
	UniformResourceIdentifierInput() *string
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
	PutDirectoryName(value *AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryName)
	PutEdiPartyName(value *AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesEdiPartyName)
	PutOtherName(value *AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOtherName)
	ResetDirectoryName()
	ResetDnsName()
	ResetEdiPartyName()
	ResetIpAddress()
	ResetOtherName()
	ResetRegisteredId()
	ResetRfc822Name()
	ResetUniformResourceIdentifier()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference
type jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) DirectoryName() AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameOutputReference {
	var returns AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameOutputReference
	_jsii_.Get(
		j,
		"directoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) DirectoryNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) DnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) EdiPartyName() AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesEdiPartyNameOutputReference {
	var returns AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesEdiPartyNameOutputReference
	_jsii_.Get(
		j,
		"ediPartyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) EdiPartyNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ediPartyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) OtherName() AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOtherNameOutputReference {
	var returns AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOtherNameOutputReference
	_jsii_.Get(
		j,
		"otherName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) OtherNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"otherNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) RegisteredId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registeredId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) RegisteredIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registeredIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) Rfc822Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rfc822Name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) Rfc822NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rfc822NameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) UniformResourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniformResourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) UniformResourceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniformResourceIdentifierInput",
		&returns,
	)
	return returns
}


func NewAcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference {
	_init_.Initialize()

	if err := validateNewAcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference{}

	_jsii_.Create(
		"awscc.acmpcaCertificate.AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference_Override(a AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.acmpcaCertificate.AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference)SetDnsName(val *string) {
	if err := j.validateSetDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsName",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference)SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference)SetRegisteredId(val *string) {
	if err := j.validateSetRegisteredIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registeredId",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference)SetRfc822Name(val *string) {
	if err := j.validateSetRfc822NameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rfc822Name",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference)SetUniformResourceIdentifier(val *string) {
	if err := j.validateSetUniformResourceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uniformResourceIdentifier",
		val,
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) PutDirectoryName(value *AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryName) {
	if err := a.validatePutDirectoryNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDirectoryName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) PutEdiPartyName(value *AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesEdiPartyName) {
	if err := a.validatePutEdiPartyNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEdiPartyName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) PutOtherName(value *AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOtherName) {
	if err := a.validatePutOtherNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOtherName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) ResetDirectoryName() {
	_jsii_.InvokeVoid(
		a,
		"resetDirectoryName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) ResetDnsName() {
	_jsii_.InvokeVoid(
		a,
		"resetDnsName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) ResetEdiPartyName() {
	_jsii_.InvokeVoid(
		a,
		"resetEdiPartyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) ResetOtherName() {
	_jsii_.InvokeVoid(
		a,
		"resetOtherName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) ResetRegisteredId() {
	_jsii_.InvokeVoid(
		a,
		"resetRegisteredId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) ResetRfc822Name() {
	_jsii_.InvokeVoid(
		a,
		"resetRfc822Name",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) ResetUniformResourceIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetUniformResourceIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

