package dataawscccloudfrontdistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscccloudfrontdistribution/internal"
)

type DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference interface {
	cdktf.ComplexObject
	AllowedMethods() *[]*string
	CachedMethods() *[]*string
	CachePolicyId() *string
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
	Compress() cdktf.IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultTtl() *float64
	FieldLevelEncryptionId() *string
	ForwardedValues() DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValuesOutputReference
	// Experimental.
	Fqn() *string
	FunctionAssociations() DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorFunctionAssociationsList
	InternalValue() *DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehavior
	SetInternalValue(val *DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehavior)
	LambdaFunctionAssociations() DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorLambdaFunctionAssociationsList
	MaxTtl() *float64
	MinTtl() *float64
	OriginRequestPolicyId() *string
	RealtimeLogConfigArn() *string
	ResponseHeadersPolicyId() *string
	SmoothStreaming() cdktf.IResolvable
	TargetOriginId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustedKeyGroups() *[]*string
	TrustedSigners() *[]*string
	ViewerProtocolPolicy() *string
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

// The jsii proxy struct for DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference
type jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) AllowedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) CachedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cachedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) CachePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) Compress() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"compress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) FieldLevelEncryptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldLevelEncryptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ForwardedValues() DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValuesOutputReference {
	var returns DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValuesOutputReference
	_jsii_.Get(
		j,
		"forwardedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) FunctionAssociations() DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorFunctionAssociationsList {
	var returns DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorFunctionAssociationsList
	_jsii_.Get(
		j,
		"functionAssociations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) InternalValue() *DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehavior {
	var returns *DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehavior
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) LambdaFunctionAssociations() DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorLambdaFunctionAssociationsList {
	var returns DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorLambdaFunctionAssociationsList
	_jsii_.Get(
		j,
		"lambdaFunctionAssociations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) MaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) MinTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) OriginRequestPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) RealtimeLogConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeLogConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResponseHeadersPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) SmoothStreaming() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"smoothStreaming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) TargetOriginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetOriginId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) TrustedKeyGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedKeyGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) TrustedSigners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedSigners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ViewerProtocolPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewerProtocolPolicy",
		&returns,
	)
	return returns
}


func NewDataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccCloudfrontDistribution.DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference_Override(d DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccCloudfrontDistribution.DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetInternalValue(val *DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehavior) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

