package dataawscccloudfrontdistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscccloudfrontdistribution/internal"
)

type DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference interface {
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
	ConnectionAttempts() *float64
	ConnectionTimeout() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomOriginConfig() DataAwsccCloudfrontDistributionDistributionConfigOriginsCustomOriginConfigOutputReference
	DomainName() *string
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *DataAwsccCloudfrontDistributionDistributionConfigOrigins
	SetInternalValue(val *DataAwsccCloudfrontDistributionDistributionConfigOrigins)
	OriginAccessControlId() *string
	OriginCustomHeaders() DataAwsccCloudfrontDistributionDistributionConfigOriginsOriginCustomHeadersList
	OriginPath() *string
	OriginShield() DataAwsccCloudfrontDistributionDistributionConfigOriginsOriginShieldOutputReference
	S3OriginConfig() DataAwsccCloudfrontDistributionDistributionConfigOriginsS3OriginConfigOutputReference
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

// The jsii proxy struct for DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference
type jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) ConnectionAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) ConnectionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) CustomOriginConfig() DataAwsccCloudfrontDistributionDistributionConfigOriginsCustomOriginConfigOutputReference {
	var returns DataAwsccCloudfrontDistributionDistributionConfigOriginsCustomOriginConfigOutputReference
	_jsii_.Get(
		j,
		"customOriginConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) InternalValue() *DataAwsccCloudfrontDistributionDistributionConfigOrigins {
	var returns *DataAwsccCloudfrontDistributionDistributionConfigOrigins
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) OriginAccessControlId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessControlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) OriginCustomHeaders() DataAwsccCloudfrontDistributionDistributionConfigOriginsOriginCustomHeadersList {
	var returns DataAwsccCloudfrontDistributionDistributionConfigOriginsOriginCustomHeadersList
	_jsii_.Get(
		j,
		"originCustomHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) OriginPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) OriginShield() DataAwsccCloudfrontDistributionDistributionConfigOriginsOriginShieldOutputReference {
	var returns DataAwsccCloudfrontDistributionDistributionConfigOriginsOriginShieldOutputReference
	_jsii_.Get(
		j,
		"originShield",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) S3OriginConfig() DataAwsccCloudfrontDistributionDistributionConfigOriginsS3OriginConfigOutputReference {
	var returns DataAwsccCloudfrontDistributionDistributionConfigOriginsS3OriginConfigOutputReference
	_jsii_.Get(
		j,
		"s3OriginConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccCloudfrontDistribution.DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference_Override(d DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccCloudfrontDistribution.DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference)SetInternalValue(val *DataAwsccCloudfrontDistributionDistributionConfigOrigins) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOriginsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

