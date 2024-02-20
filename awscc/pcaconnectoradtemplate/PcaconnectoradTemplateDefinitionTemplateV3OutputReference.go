package pcaconnectoradtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/pcaconnectoradtemplate/internal"
)

type PcaconnectoradTemplateDefinitionTemplateV3OutputReference interface {
	cdktf.ComplexObject
	CertificateValidity() PcaconnectoradTemplateDefinitionTemplateV3CertificateValidityOutputReference
	CertificateValidityInput() *PcaconnectoradTemplateDefinitionTemplateV3CertificateValidity
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
	EnrollmentFlags() PcaconnectoradTemplateDefinitionTemplateV3EnrollmentFlagsOutputReference
	EnrollmentFlagsInput() *PcaconnectoradTemplateDefinitionTemplateV3EnrollmentFlags
	Extensions() PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference
	ExtensionsInput() *PcaconnectoradTemplateDefinitionTemplateV3Extensions
	// Experimental.
	Fqn() *string
	GeneralFlags() PcaconnectoradTemplateDefinitionTemplateV3GeneralFlagsOutputReference
	GeneralFlagsInput() *PcaconnectoradTemplateDefinitionTemplateV3GeneralFlags
	HashAlgorithm() *string
	SetHashAlgorithm(val *string)
	HashAlgorithmInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PrivateKeyAttributes() PcaconnectoradTemplateDefinitionTemplateV3PrivateKeyAttributesOutputReference
	PrivateKeyAttributesInput() *PcaconnectoradTemplateDefinitionTemplateV3PrivateKeyAttributes
	PrivateKeyFlags() PcaconnectoradTemplateDefinitionTemplateV3PrivateKeyFlagsOutputReference
	PrivateKeyFlagsInput() *PcaconnectoradTemplateDefinitionTemplateV3PrivateKeyFlags
	SubjectNameFlags() PcaconnectoradTemplateDefinitionTemplateV3SubjectNameFlagsOutputReference
	SubjectNameFlagsInput() *PcaconnectoradTemplateDefinitionTemplateV3SubjectNameFlags
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
	PutCertificateValidity(value *PcaconnectoradTemplateDefinitionTemplateV3CertificateValidity)
	PutEnrollmentFlags(value *PcaconnectoradTemplateDefinitionTemplateV3EnrollmentFlags)
	PutExtensions(value *PcaconnectoradTemplateDefinitionTemplateV3Extensions)
	PutGeneralFlags(value *PcaconnectoradTemplateDefinitionTemplateV3GeneralFlags)
	PutPrivateKeyAttributes(value *PcaconnectoradTemplateDefinitionTemplateV3PrivateKeyAttributes)
	PutPrivateKeyFlags(value *PcaconnectoradTemplateDefinitionTemplateV3PrivateKeyFlags)
	PutSubjectNameFlags(value *PcaconnectoradTemplateDefinitionTemplateV3SubjectNameFlags)
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

// The jsii proxy struct for PcaconnectoradTemplateDefinitionTemplateV3OutputReference
type jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) CertificateValidity() PcaconnectoradTemplateDefinitionTemplateV3CertificateValidityOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV3CertificateValidityOutputReference
	_jsii_.Get(
		j,
		"certificateValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) CertificateValidityInput() *PcaconnectoradTemplateDefinitionTemplateV3CertificateValidity {
	var returns *PcaconnectoradTemplateDefinitionTemplateV3CertificateValidity
	_jsii_.Get(
		j,
		"certificateValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) EnrollmentFlags() PcaconnectoradTemplateDefinitionTemplateV3EnrollmentFlagsOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV3EnrollmentFlagsOutputReference
	_jsii_.Get(
		j,
		"enrollmentFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) EnrollmentFlagsInput() *PcaconnectoradTemplateDefinitionTemplateV3EnrollmentFlags {
	var returns *PcaconnectoradTemplateDefinitionTemplateV3EnrollmentFlags
	_jsii_.Get(
		j,
		"enrollmentFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) Extensions() PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV3ExtensionsOutputReference
	_jsii_.Get(
		j,
		"extensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) ExtensionsInput() *PcaconnectoradTemplateDefinitionTemplateV3Extensions {
	var returns *PcaconnectoradTemplateDefinitionTemplateV3Extensions
	_jsii_.Get(
		j,
		"extensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) GeneralFlags() PcaconnectoradTemplateDefinitionTemplateV3GeneralFlagsOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV3GeneralFlagsOutputReference
	_jsii_.Get(
		j,
		"generalFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) GeneralFlagsInput() *PcaconnectoradTemplateDefinitionTemplateV3GeneralFlags {
	var returns *PcaconnectoradTemplateDefinitionTemplateV3GeneralFlags
	_jsii_.Get(
		j,
		"generalFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) HashAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) HashAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) PrivateKeyAttributes() PcaconnectoradTemplateDefinitionTemplateV3PrivateKeyAttributesOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV3PrivateKeyAttributesOutputReference
	_jsii_.Get(
		j,
		"privateKeyAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) PrivateKeyAttributesInput() *PcaconnectoradTemplateDefinitionTemplateV3PrivateKeyAttributes {
	var returns *PcaconnectoradTemplateDefinitionTemplateV3PrivateKeyAttributes
	_jsii_.Get(
		j,
		"privateKeyAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) PrivateKeyFlags() PcaconnectoradTemplateDefinitionTemplateV3PrivateKeyFlagsOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV3PrivateKeyFlagsOutputReference
	_jsii_.Get(
		j,
		"privateKeyFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) PrivateKeyFlagsInput() *PcaconnectoradTemplateDefinitionTemplateV3PrivateKeyFlags {
	var returns *PcaconnectoradTemplateDefinitionTemplateV3PrivateKeyFlags
	_jsii_.Get(
		j,
		"privateKeyFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) SubjectNameFlags() PcaconnectoradTemplateDefinitionTemplateV3SubjectNameFlagsOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV3SubjectNameFlagsOutputReference
	_jsii_.Get(
		j,
		"subjectNameFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) SubjectNameFlagsInput() *PcaconnectoradTemplateDefinitionTemplateV3SubjectNameFlags {
	var returns *PcaconnectoradTemplateDefinitionTemplateV3SubjectNameFlags
	_jsii_.Get(
		j,
		"subjectNameFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) SupersededTemplates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supersededTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) SupersededTemplatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supersededTemplatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPcaconnectoradTemplateDefinitionTemplateV3OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PcaconnectoradTemplateDefinitionTemplateV3OutputReference {
	_init_.Initialize()

	if err := validateNewPcaconnectoradTemplateDefinitionTemplateV3OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference{}

	_jsii_.Create(
		"awscc.pcaconnectoradTemplate.PcaconnectoradTemplateDefinitionTemplateV3OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPcaconnectoradTemplateDefinitionTemplateV3OutputReference_Override(p PcaconnectoradTemplateDefinitionTemplateV3OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.pcaconnectoradTemplate.PcaconnectoradTemplateDefinitionTemplateV3OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference)SetHashAlgorithm(val *string) {
	if err := j.validateSetHashAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashAlgorithm",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference)SetSupersededTemplates(val *[]*string) {
	if err := j.validateSetSupersededTemplatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supersededTemplates",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) PutCertificateValidity(value *PcaconnectoradTemplateDefinitionTemplateV3CertificateValidity) {
	if err := p.validatePutCertificateValidityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCertificateValidity",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) PutEnrollmentFlags(value *PcaconnectoradTemplateDefinitionTemplateV3EnrollmentFlags) {
	if err := p.validatePutEnrollmentFlagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEnrollmentFlags",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) PutExtensions(value *PcaconnectoradTemplateDefinitionTemplateV3Extensions) {
	if err := p.validatePutExtensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putExtensions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) PutGeneralFlags(value *PcaconnectoradTemplateDefinitionTemplateV3GeneralFlags) {
	if err := p.validatePutGeneralFlagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGeneralFlags",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) PutPrivateKeyAttributes(value *PcaconnectoradTemplateDefinitionTemplateV3PrivateKeyAttributes) {
	if err := p.validatePutPrivateKeyAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPrivateKeyAttributes",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) PutPrivateKeyFlags(value *PcaconnectoradTemplateDefinitionTemplateV3PrivateKeyFlags) {
	if err := p.validatePutPrivateKeyFlagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPrivateKeyFlags",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) PutSubjectNameFlags(value *PcaconnectoradTemplateDefinitionTemplateV3SubjectNameFlags) {
	if err := p.validatePutSubjectNameFlagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSubjectNameFlags",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) ResetSupersededTemplates() {
	_jsii_.InvokeVoid(
		p,
		"resetSupersededTemplates",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV3OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

