package dataawsccdatazoneconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccdatazoneconnection/internal"
)

type DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference interface {
	cdktf.ComplexObject
	AthenaProperties() cdktf.StringMap
	AuthenticationConfiguration() DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference
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
	ConnectionProperties() cdktf.StringMap
	ConnectionType() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInput
	SetInternalValue(val *DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInput)
	MatchCriteria() *string
	Name() *string
	PhysicalConnectionRequirements() DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputPhysicalConnectionRequirementsOutputReference
	PythonProperties() cdktf.StringMap
	SparkProperties() cdktf.StringMap
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValidateCredentials() cdktf.IResolvable
	ValidateForComputeEnvironments() *[]*string
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

// The jsii proxy struct for DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference
type jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) AthenaProperties() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"athenaProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) AuthenticationConfiguration() DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference {
	var returns DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference
	_jsii_.Get(
		j,
		"authenticationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ConnectionProperties() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"connectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ConnectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) InternalValue() *DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInput {
	var returns *DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) MatchCriteria() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) PhysicalConnectionRequirements() DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputPhysicalConnectionRequirementsOutputReference {
	var returns DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputPhysicalConnectionRequirementsOutputReference
	_jsii_.Get(
		j,
		"physicalConnectionRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) PythonProperties() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"pythonProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) SparkProperties() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"sparkProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ValidateCredentials() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"validateCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ValidateForComputeEnvironments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validateForComputeEnvironments",
		&returns,
	)
	return returns
}


func NewDataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccDatazoneConnection.DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference_Override(d DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccDatazoneConnection.DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetInternalValue(val *DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

