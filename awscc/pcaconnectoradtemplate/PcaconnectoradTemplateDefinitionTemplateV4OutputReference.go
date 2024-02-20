package pcaconnectoradtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/pcaconnectoradtemplate/internal"
)

type PcaconnectoradTemplateDefinitionTemplateV4OutputReference interface {
	cdktf.ComplexObject
	CertificateValidity() PcaconnectoradTemplateDefinitionTemplateV4CertificateValidityOutputReference
	CertificateValidityInput() *PcaconnectoradTemplateDefinitionTemplateV4CertificateValidity
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
	EnrollmentFlags() PcaconnectoradTemplateDefinitionTemplateV4EnrollmentFlagsOutputReference
	EnrollmentFlagsInput() *PcaconnectoradTemplateDefinitionTemplateV4EnrollmentFlags
	Extensions() PcaconnectoradTemplateDefinitionTemplateV4ExtensionsOutputReference
	ExtensionsInput() *PcaconnectoradTemplateDefinitionTemplateV4Extensions
	// Experimental.
	Fqn() *string
	GeneralFlags() PcaconnectoradTemplateDefinitionTemplateV4GeneralFlagsOutputReference
	GeneralFlagsInput() *PcaconnectoradTemplateDefinitionTemplateV4GeneralFlags
	HashAlgorithm() *string
	SetHashAlgorithm(val *string)
	HashAlgorithmInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PrivateKeyAttributes() PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyAttributesOutputReference
	PrivateKeyAttributesInput() *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyAttributes
	PrivateKeyFlags() PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference
	PrivateKeyFlagsInput() *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlags
	SubjectNameFlags() PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference
	SubjectNameFlagsInput() *PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlags
	SupersededTemplates() *[]*string
	SetSupersededTemplates(val *[]*string)
	SupersededTemplatesInput() *[]*string
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
	PutCertificateValidity(value *PcaconnectoradTemplateDefinitionTemplateV4CertificateValidity)
	PutEnrollmentFlags(value *PcaconnectoradTemplateDefinitionTemplateV4EnrollmentFlags)
	PutExtensions(value *PcaconnectoradTemplateDefinitionTemplateV4Extensions)
	PutGeneralFlags(value *PcaconnectoradTemplateDefinitionTemplateV4GeneralFlags)
	PutPrivateKeyAttributes(value *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyAttributes)
	PutPrivateKeyFlags(value *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlags)
	PutSubjectNameFlags(value *PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlags)
	ResetHashAlgorithm()
	ResetSupersededTemplates()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PcaconnectoradTemplateDefinitionTemplateV4OutputReference
type jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) CertificateValidity() PcaconnectoradTemplateDefinitionTemplateV4CertificateValidityOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV4CertificateValidityOutputReference
	_jsii_.Get(
		j,
		"certificateValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) CertificateValidityInput() *PcaconnectoradTemplateDefinitionTemplateV4CertificateValidity {
	var returns *PcaconnectoradTemplateDefinitionTemplateV4CertificateValidity
	_jsii_.Get(
		j,
		"certificateValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) EnrollmentFlags() PcaconnectoradTemplateDefinitionTemplateV4EnrollmentFlagsOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV4EnrollmentFlagsOutputReference
	_jsii_.Get(
		j,
		"enrollmentFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) EnrollmentFlagsInput() *PcaconnectoradTemplateDefinitionTemplateV4EnrollmentFlags {
	var returns *PcaconnectoradTemplateDefinitionTemplateV4EnrollmentFlags
	_jsii_.Get(
		j,
		"enrollmentFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) Extensions() PcaconnectoradTemplateDefinitionTemplateV4ExtensionsOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV4ExtensionsOutputReference
	_jsii_.Get(
		j,
		"extensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) ExtensionsInput() *PcaconnectoradTemplateDefinitionTemplateV4Extensions {
	var returns *PcaconnectoradTemplateDefinitionTemplateV4Extensions
	_jsii_.Get(
		j,
		"extensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) GeneralFlags() PcaconnectoradTemplateDefinitionTemplateV4GeneralFlagsOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV4GeneralFlagsOutputReference
	_jsii_.Get(
		j,
		"generalFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) GeneralFlagsInput() *PcaconnectoradTemplateDefinitionTemplateV4GeneralFlags {
	var returns *PcaconnectoradTemplateDefinitionTemplateV4GeneralFlags
	_jsii_.Get(
		j,
		"generalFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) HashAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) HashAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) PrivateKeyAttributes() PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyAttributesOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyAttributesOutputReference
	_jsii_.Get(
		j,
		"privateKeyAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) PrivateKeyAttributesInput() *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyAttributes {
	var returns *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyAttributes
	_jsii_.Get(
		j,
		"privateKeyAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) PrivateKeyFlags() PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference
	_jsii_.Get(
		j,
		"privateKeyFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) PrivateKeyFlagsInput() *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlags {
	var returns *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlags
	_jsii_.Get(
		j,
		"privateKeyFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) SubjectNameFlags() PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference
	_jsii_.Get(
		j,
		"subjectNameFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) SubjectNameFlagsInput() *PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlags {
	var returns *PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlags
	_jsii_.Get(
		j,
		"subjectNameFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) SupersededTemplates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supersededTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) SupersededTemplatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supersededTemplatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPcaconnectoradTemplateDefinitionTemplateV4OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PcaconnectoradTemplateDefinitionTemplateV4OutputReference {
	_init_.Initialize()

	if err := validateNewPcaconnectoradTemplateDefinitionTemplateV4OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference{}

	_jsii_.Create(
		"awscc.pcaconnectoradTemplate.PcaconnectoradTemplateDefinitionTemplateV4OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPcaconnectoradTemplateDefinitionTemplateV4OutputReference_Override(p PcaconnectoradTemplateDefinitionTemplateV4OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.pcaconnectoradTemplate.PcaconnectoradTemplateDefinitionTemplateV4OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference)SetHashAlgorithm(val *string) {
	if err := j.validateSetHashAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashAlgorithm",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference)SetSupersededTemplates(val *[]*string) {
	if err := j.validateSetSupersededTemplatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supersededTemplates",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) PutCertificateValidity(value *PcaconnectoradTemplateDefinitionTemplateV4CertificateValidity) {
	if err := p.validatePutCertificateValidityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCertificateValidity",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) PutEnrollmentFlags(value *PcaconnectoradTemplateDefinitionTemplateV4EnrollmentFlags) {
	if err := p.validatePutEnrollmentFlagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEnrollmentFlags",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) PutExtensions(value *PcaconnectoradTemplateDefinitionTemplateV4Extensions) {
	if err := p.validatePutExtensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putExtensions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) PutGeneralFlags(value *PcaconnectoradTemplateDefinitionTemplateV4GeneralFlags) {
	if err := p.validatePutGeneralFlagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGeneralFlags",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) PutPrivateKeyAttributes(value *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyAttributes) {
	if err := p.validatePutPrivateKeyAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPrivateKeyAttributes",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) PutPrivateKeyFlags(value *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlags) {
	if err := p.validatePutPrivateKeyFlagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPrivateKeyFlags",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) PutSubjectNameFlags(value *PcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlags) {
	if err := p.validatePutSubjectNameFlagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSubjectNameFlags",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) ResetHashAlgorithm() {
	_jsii_.InvokeVoid(
		p,
		"resetHashAlgorithm",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) ResetSupersededTemplates() {
	_jsii_.InvokeVoid(
		p,
		"resetSupersededTemplates",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV4OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

