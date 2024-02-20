package dataawscccloudfrontcontinuousdeploymentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscccloudfrontcontinuousdeploymentpolicy/internal"
)

type DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference interface {
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
	Enabled() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfig
	SetInternalValue(val *DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfig)
	SingleHeaderPolicyConfig() DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleHeaderPolicyConfigOutputReference
	SingleWeightPolicyConfig() DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleWeightPolicyConfigOutputReference
	StagingDistributionDnsNames() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrafficConfig() DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigTrafficConfigOutputReference
	Type() *string
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

// The jsii proxy struct for DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference
type jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) InternalValue() *DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfig {
	var returns *DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) SingleHeaderPolicyConfig() DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleHeaderPolicyConfigOutputReference {
	var returns DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleHeaderPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"singleHeaderPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) SingleWeightPolicyConfig() DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleWeightPolicyConfigOutputReference {
	var returns DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleWeightPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"singleWeightPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) StagingDistributionDnsNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stagingDistributionDnsNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) TrafficConfig() DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigTrafficConfigOutputReference {
	var returns DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigTrafficConfigOutputReference
	_jsii_.Get(
		j,
		"trafficConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewDataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccCloudfrontContinuousDeploymentPolicy.DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference_Override(d DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccCloudfrontContinuousDeploymentPolicy.DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference)SetInternalValue(val *DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

