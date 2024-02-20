package dataawsccacmpcacertificateauthority

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccacmpcacertificateauthority/internal"
)

type DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference interface {
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
	CrlDistributionPointExtensionConfiguration() DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationCrlDistributionPointExtensionConfigurationOutputReference
	CustomCname() *string
	Enabled() cdktf.IResolvable
	ExpirationInDays() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration
	SetInternalValue(val *DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration)
	S3BucketName() *string
	S3ObjectAcl() *string
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

// The jsii proxy struct for DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference
type jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) CrlDistributionPointExtensionConfiguration() DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationCrlDistributionPointExtensionConfigurationOutputReference {
	var returns DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationCrlDistributionPointExtensionConfigurationOutputReference
	_jsii_.Get(
		j,
		"crlDistributionPointExtensionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) CustomCname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) ExpirationInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) InternalValue() *DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration {
	var returns *DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) S3ObjectAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccAcmpcaCertificateAuthority.DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference_Override(d DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccAcmpcaCertificateAuthority.DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference)SetInternalValue(val *DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

