package dataawsccamplifyapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccamplifyapp/internal"
)

type DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference interface {
	cdktf.ComplexObject
	AutoBranchCreationPatterns() *[]*string
	BasicAuthConfig() DataAwsccAmplifyAppAutoBranchCreationConfigBasicAuthConfigOutputReference
	BuildSpec() *string
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
	EnableAutoBranchCreation() cdktf.IResolvable
	EnableAutoBuild() cdktf.IResolvable
	EnablePerformanceMode() cdktf.IResolvable
	EnablePullRequestPreview() cdktf.IResolvable
	EnvironmentVariables() DataAwsccAmplifyAppAutoBranchCreationConfigEnvironmentVariablesList
	// Experimental.
	Fqn() *string
	Framework() *string
	InternalValue() *DataAwsccAmplifyAppAutoBranchCreationConfig
	SetInternalValue(val *DataAwsccAmplifyAppAutoBranchCreationConfig)
	PullRequestEnvironmentName() *string
	Stage() *string
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

// The jsii proxy struct for DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference
type jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) AutoBranchCreationPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoBranchCreationPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) BasicAuthConfig() DataAwsccAmplifyAppAutoBranchCreationConfigBasicAuthConfigOutputReference {
	var returns DataAwsccAmplifyAppAutoBranchCreationConfigBasicAuthConfigOutputReference
	_jsii_.Get(
		j,
		"basicAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) BuildSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) EnableAutoBranchCreation() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableAutoBranchCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) EnableAutoBuild() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableAutoBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) EnablePerformanceMode() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enablePerformanceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) EnablePullRequestPreview() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enablePullRequestPreview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) EnvironmentVariables() DataAwsccAmplifyAppAutoBranchCreationConfigEnvironmentVariablesList {
	var returns DataAwsccAmplifyAppAutoBranchCreationConfigEnvironmentVariablesList
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) Framework() *string {
	var returns *string
	_jsii_.Get(
		j,
		"framework",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) InternalValue() *DataAwsccAmplifyAppAutoBranchCreationConfig {
	var returns *DataAwsccAmplifyAppAutoBranchCreationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) PullRequestEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pullRequestEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccAmplifyAppAutoBranchCreationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccAmplifyAppAutoBranchCreationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccAmplifyApp.DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccAmplifyAppAutoBranchCreationConfigOutputReference_Override(d DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccAmplifyApp.DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference)SetInternalValue(val *DataAwsccAmplifyAppAutoBranchCreationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccAmplifyAppAutoBranchCreationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

