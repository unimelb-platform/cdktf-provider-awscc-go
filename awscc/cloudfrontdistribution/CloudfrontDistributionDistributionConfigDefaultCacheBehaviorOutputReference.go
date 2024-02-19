package cloudfrontdistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cloudfrontdistribution/internal"
)

type CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference interface {
	cdktf.ComplexObject
	AllowedMethods() *[]*string
	SetAllowedMethods(val *[]*string)
	AllowedMethodsInput() *[]*string
	CachedMethods() *[]*string
	SetCachedMethods(val *[]*string)
	CachedMethodsInput() *[]*string
	CachePolicyId() *string
	SetCachePolicyId(val *string)
	CachePolicyIdInput() *string
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
	Compress() interface{}
	SetCompress(val interface{})
	CompressInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultTtl() *float64
	SetDefaultTtl(val *float64)
	DefaultTtlInput() *float64
	FieldLevelEncryptionId() *string
	SetFieldLevelEncryptionId(val *string)
	FieldLevelEncryptionIdInput() *string
	ForwardedValues() CloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValuesOutputReference
	ForwardedValuesInput() interface{}
	// Experimental.
	Fqn() *string
	FunctionAssociations() CloudfrontDistributionDistributionConfigDefaultCacheBehaviorFunctionAssociationsList
	FunctionAssociationsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LambdaFunctionAssociations() CloudfrontDistributionDistributionConfigDefaultCacheBehaviorLambdaFunctionAssociationsList
	LambdaFunctionAssociationsInput() interface{}
	MaxTtl() *float64
	SetMaxTtl(val *float64)
	MaxTtlInput() *float64
	MinTtl() *float64
	SetMinTtl(val *float64)
	MinTtlInput() *float64
	OriginRequestPolicyId() *string
	SetOriginRequestPolicyId(val *string)
	OriginRequestPolicyIdInput() *string
	RealtimeLogConfigArn() *string
	SetRealtimeLogConfigArn(val *string)
	RealtimeLogConfigArnInput() *string
	ResponseHeadersPolicyId() *string
	SetResponseHeadersPolicyId(val *string)
	ResponseHeadersPolicyIdInput() *string
	SmoothStreaming() interface{}
	SetSmoothStreaming(val interface{})
	SmoothStreamingInput() interface{}
	TargetOriginId() *string
	SetTargetOriginId(val *string)
	TargetOriginIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustedKeyGroups() *[]*string
	SetTrustedKeyGroups(val *[]*string)
	TrustedKeyGroupsInput() *[]*string
	TrustedSigners() *[]*string
	SetTrustedSigners(val *[]*string)
	TrustedSignersInput() *[]*string
	ViewerProtocolPolicy() *string
	SetViewerProtocolPolicy(val *string)
	ViewerProtocolPolicyInput() *string
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
	PutForwardedValues(value *CloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValues)
	PutFunctionAssociations(value interface{})
	PutLambdaFunctionAssociations(value interface{})
	ResetAllowedMethods()
	ResetCachedMethods()
	ResetCachePolicyId()
	ResetCompress()
	ResetDefaultTtl()
	ResetFieldLevelEncryptionId()
	ResetForwardedValues()
	ResetFunctionAssociations()
	ResetLambdaFunctionAssociations()
	ResetMaxTtl()
	ResetMinTtl()
	ResetOriginRequestPolicyId()
	ResetRealtimeLogConfigArn()
	ResetResponseHeadersPolicyId()
	ResetSmoothStreaming()
	ResetTrustedKeyGroups()
	ResetTrustedSigners()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference
type jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) AllowedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) AllowedMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) CachedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cachedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) CachedMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cachedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) CachePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) CachePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) Compress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) CompressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) DefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) FieldLevelEncryptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldLevelEncryptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) FieldLevelEncryptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldLevelEncryptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ForwardedValues() CloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValuesOutputReference {
	var returns CloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValuesOutputReference
	_jsii_.Get(
		j,
		"forwardedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ForwardedValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardedValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) FunctionAssociations() CloudfrontDistributionDistributionConfigDefaultCacheBehaviorFunctionAssociationsList {
	var returns CloudfrontDistributionDistributionConfigDefaultCacheBehaviorFunctionAssociationsList
	_jsii_.Get(
		j,
		"functionAssociations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) FunctionAssociationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionAssociationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) LambdaFunctionAssociations() CloudfrontDistributionDistributionConfigDefaultCacheBehaviorLambdaFunctionAssociationsList {
	var returns CloudfrontDistributionDistributionConfigDefaultCacheBehaviorLambdaFunctionAssociationsList
	_jsii_.Get(
		j,
		"lambdaFunctionAssociations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) LambdaFunctionAssociationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionAssociationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) MaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) MaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) MinTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) MinTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) OriginRequestPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) OriginRequestPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) RealtimeLogConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeLogConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) RealtimeLogConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeLogConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResponseHeadersPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResponseHeadersPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) SmoothStreaming() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smoothStreaming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) SmoothStreamingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smoothStreamingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) TargetOriginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetOriginId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) TargetOriginIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetOriginIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) TrustedKeyGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedKeyGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) TrustedKeyGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedKeyGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) TrustedSigners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedSigners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) TrustedSignersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedSignersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ViewerProtocolPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewerProtocolPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ViewerProtocolPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewerProtocolPolicyInput",
		&returns,
	)
	return returns
}


func NewCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference{}

	_jsii_.Create(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference_Override(c CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetAllowedMethods(val *[]*string) {
	if err := j.validateSetAllowedMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedMethods",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetCachedMethods(val *[]*string) {
	if err := j.validateSetCachedMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cachedMethods",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetCachePolicyId(val *string) {
	if err := j.validateSetCachePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cachePolicyId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetCompress(val interface{}) {
	if err := j.validateSetCompressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compress",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetDefaultTtl(val *float64) {
	if err := j.validateSetDefaultTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTtl",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetFieldLevelEncryptionId(val *string) {
	if err := j.validateSetFieldLevelEncryptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldLevelEncryptionId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetMaxTtl(val *float64) {
	if err := j.validateSetMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTtl",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetMinTtl(val *float64) {
	if err := j.validateSetMinTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTtl",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetOriginRequestPolicyId(val *string) {
	if err := j.validateSetOriginRequestPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originRequestPolicyId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetRealtimeLogConfigArn(val *string) {
	if err := j.validateSetRealtimeLogConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realtimeLogConfigArn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetResponseHeadersPolicyId(val *string) {
	if err := j.validateSetResponseHeadersPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseHeadersPolicyId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetSmoothStreaming(val interface{}) {
	if err := j.validateSetSmoothStreamingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smoothStreaming",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetTargetOriginId(val *string) {
	if err := j.validateSetTargetOriginIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetOriginId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetTrustedKeyGroups(val *[]*string) {
	if err := j.validateSetTrustedKeyGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedKeyGroups",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetTrustedSigners(val *[]*string) {
	if err := j.validateSetTrustedSignersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedSigners",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference)SetViewerProtocolPolicy(val *string) {
	if err := j.validateSetViewerProtocolPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewerProtocolPolicy",
		val,
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) PutForwardedValues(value *CloudfrontDistributionDistributionConfigDefaultCacheBehaviorForwardedValues) {
	if err := c.validatePutForwardedValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putForwardedValues",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) PutFunctionAssociations(value interface{}) {
	if err := c.validatePutFunctionAssociationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFunctionAssociations",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) PutLambdaFunctionAssociations(value interface{}) {
	if err := c.validatePutLambdaFunctionAssociationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLambdaFunctionAssociations",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResetAllowedMethods() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedMethods",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResetCachedMethods() {
	_jsii_.InvokeVoid(
		c,
		"resetCachedMethods",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResetCachePolicyId() {
	_jsii_.InvokeVoid(
		c,
		"resetCachePolicyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResetCompress() {
	_jsii_.InvokeVoid(
		c,
		"resetCompress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResetFieldLevelEncryptionId() {
	_jsii_.InvokeVoid(
		c,
		"resetFieldLevelEncryptionId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResetForwardedValues() {
	_jsii_.InvokeVoid(
		c,
		"resetForwardedValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResetFunctionAssociations() {
	_jsii_.InvokeVoid(
		c,
		"resetFunctionAssociations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResetLambdaFunctionAssociations() {
	_jsii_.InvokeVoid(
		c,
		"resetLambdaFunctionAssociations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResetMinTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetMinTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResetOriginRequestPolicyId() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginRequestPolicyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResetRealtimeLogConfigArn() {
	_jsii_.InvokeVoid(
		c,
		"resetRealtimeLogConfigArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResetResponseHeadersPolicyId() {
	_jsii_.InvokeVoid(
		c,
		"resetResponseHeadersPolicyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResetSmoothStreaming() {
	_jsii_.InvokeVoid(
		c,
		"resetSmoothStreaming",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResetTrustedKeyGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetTrustedKeyGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ResetTrustedSigners() {
	_jsii_.InvokeVoid(
		c,
		"resetTrustedSigners",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

