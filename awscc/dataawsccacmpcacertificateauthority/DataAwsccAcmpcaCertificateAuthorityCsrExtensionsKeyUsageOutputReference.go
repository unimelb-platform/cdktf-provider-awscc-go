package dataawsccacmpcacertificateauthority

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccacmpcacertificateauthority/internal"
)

type DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference interface {
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
	CrlSign() cdktf.IResolvable
	DataEncipherment() cdktf.IResolvable
	DecipherOnly() cdktf.IResolvable
	DigitalSignature() cdktf.IResolvable
	EncipherOnly() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsage
	SetInternalValue(val *DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsage)
	KeyAgreement() cdktf.IResolvable
	KeyCertSign() cdktf.IResolvable
	KeyEncipherment() cdktf.IResolvable
	NonRepudiation() cdktf.IResolvable
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference
type jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) CrlSign() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"crlSign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) DataEncipherment() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"dataEncipherment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) DecipherOnly() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"decipherOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) DigitalSignature() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"digitalSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) EncipherOnly() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"encipherOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) InternalValue() *DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsage {
	var returns *DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) KeyAgreement() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"keyAgreement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) KeyCertSign() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"keyCertSign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) KeyEncipherment() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"keyEncipherment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) NonRepudiation() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"nonRepudiation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccAcmpcaCertificateAuthority.DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference_Override(d DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccAcmpcaCertificateAuthority.DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetInternalValue(val *DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityCsrExtensionsKeyUsageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

