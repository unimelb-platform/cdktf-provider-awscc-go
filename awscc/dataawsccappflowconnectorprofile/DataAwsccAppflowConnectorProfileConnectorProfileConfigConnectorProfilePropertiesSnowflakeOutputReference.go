package dataawsccappflowconnectorprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccappflowconnectorprofile/internal"
)

type DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference interface {
	cdktf.ComplexObject
	AccountName() *string
	BucketName() *string
	BucketPrefix() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflake
	SetInternalValue(val *DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflake)
	PrivateLinkServiceName() *string
	Region() *string
	Stage() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Warehouse() *string
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

// The jsii proxy struct for DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference
type jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) BucketPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) InternalValue() *DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflake {
	var returns *DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflake
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) PrivateLinkServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) Warehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouse",
		&returns,
	)
	return returns
}


func NewDataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccAppflowConnectorProfile.DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference_Override(d DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccAppflowConnectorProfile.DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference)SetInternalValue(val *DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflake) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

