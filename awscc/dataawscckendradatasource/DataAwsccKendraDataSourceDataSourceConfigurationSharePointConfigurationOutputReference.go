package dataawscckendradatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscckendradatasource/internal"
)

type DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference interface {
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
	CrawlAttachments() cdktf.IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisableLocalGroups() cdktf.IResolvable
	DocumentTitleFieldName() *string
	ExclusionPatterns() *[]*string
	FieldMappings() DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationFieldMappingsList
	// Experimental.
	Fqn() *string
	InclusionPatterns() *[]*string
	InternalValue() *DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfiguration
	SetInternalValue(val *DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfiguration)
	SecretArn() *string
	SharePointVersion() *string
	SslCertificateS3Path() DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationSslCertificateS3PathOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Urls() *[]*string
	UseChangeLog() cdktf.IResolvable
	VpcConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationVpcConfigurationOutputReference
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

// The jsii proxy struct for DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference
type jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) CrawlAttachments() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"crawlAttachments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) DisableLocalGroups() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableLocalGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) DocumentTitleFieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentTitleFieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ExclusionPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusionPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) FieldMappings() DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationFieldMappingsList {
	var returns DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationFieldMappingsList
	_jsii_.Get(
		j,
		"fieldMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) InclusionPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) InternalValue() *DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfiguration {
	var returns *DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) SharePointVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePointVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) SslCertificateS3Path() DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationSslCertificateS3PathOutputReference {
	var returns DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationSslCertificateS3PathOutputReference
	_jsii_.Get(
		j,
		"sslCertificateS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) Urls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) UseChangeLog() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"useChangeLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) VpcConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationVpcConfigurationOutputReference {
	var returns DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationVpcConfigurationOutputReference
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}


func NewDataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccKendraDataSource.DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference_Override(d DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccKendraDataSource.DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetInternalValue(val *DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

