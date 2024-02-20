package cloudfrontdistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cloudfrontdistribution/internal"
)

type CloudfrontDistributionDistributionConfigOutputReference interface {
	cdktf.ComplexObject
	Aliases() *[]*string
	SetAliases(val *[]*string)
	AliasesInput() *[]*string
	CacheBehaviors() CloudfrontDistributionDistributionConfigCacheBehaviorsList
	CacheBehaviorsInput() interface{}
	Cnames() *[]*string
	SetCnames(val *[]*string)
	CnamesInput() *[]*string
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
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
	ContinuousDeploymentPolicyId() *string
	SetContinuousDeploymentPolicyId(val *string)
	ContinuousDeploymentPolicyIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomErrorResponses() CloudfrontDistributionDistributionConfigCustomErrorResponsesList
	CustomErrorResponsesInput() interface{}
	CustomOrigin() CloudfrontDistributionDistributionConfigCustomOriginOutputReference
	CustomOriginInput() interface{}
	DefaultCacheBehavior() CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference
	DefaultCacheBehaviorInput() interface{}
	DefaultRootObject() *string
	SetDefaultRootObject(val *string)
	DefaultRootObjectInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	HttpVersion() *string
	SetHttpVersion(val *string)
	HttpVersionInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ipv6Enabled() interface{}
	SetIpv6Enabled(val interface{})
	Ipv6EnabledInput() interface{}
	Logging() CloudfrontDistributionDistributionConfigLoggingOutputReference
	LoggingInput() interface{}
	OriginGroups() CloudfrontDistributionDistributionConfigOriginGroupsOutputReference
	OriginGroupsInput() interface{}
	Origins() CloudfrontDistributionDistributionConfigOriginsList
	OriginsInput() interface{}
	PriceClass() *string
	SetPriceClass(val *string)
	PriceClassInput() *string
	Restrictions() CloudfrontDistributionDistributionConfigRestrictionsOutputReference
	RestrictionsInput() interface{}
	S3Origin() CloudfrontDistributionDistributionConfigS3OriginOutputReference
	S3OriginInput() interface{}
	Staging() interface{}
	SetStaging(val interface{})
	StagingInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ViewerCertificate() CloudfrontDistributionDistributionConfigViewerCertificateOutputReference
	ViewerCertificateInput() interface{}
	WebAclId() *string
	SetWebAclId(val *string)
	WebAclIdInput() *string
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
	PutCacheBehaviors(value interface{})
	PutCustomErrorResponses(value interface{})
	PutCustomOrigin(value *CloudfrontDistributionDistributionConfigCustomOrigin)
	PutDefaultCacheBehavior(value *CloudfrontDistributionDistributionConfigDefaultCacheBehavior)
	PutLogging(value *CloudfrontDistributionDistributionConfigLogging)
	PutOriginGroups(value *CloudfrontDistributionDistributionConfigOriginGroups)
	PutOrigins(value interface{})
	PutRestrictions(value *CloudfrontDistributionDistributionConfigRestrictions)
	PutS3Origin(value *CloudfrontDistributionDistributionConfigS3Origin)
	PutViewerCertificate(value *CloudfrontDistributionDistributionConfigViewerCertificate)
	ResetAliases()
	ResetCacheBehaviors()
	ResetCnames()
	ResetComment()
	ResetContinuousDeploymentPolicyId()
	ResetCustomErrorResponses()
	ResetCustomOrigin()
	ResetDefaultRootObject()
	ResetHttpVersion()
	ResetIpv6Enabled()
	ResetLogging()
	ResetOriginGroups()
	ResetOrigins()
	ResetPriceClass()
	ResetRestrictions()
	ResetS3Origin()
	ResetStaging()
	ResetViewerCertificate()
	ResetWebAclId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontDistributionDistributionConfigOutputReference
type jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) AliasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) CacheBehaviors() CloudfrontDistributionDistributionConfigCacheBehaviorsList {
	var returns CloudfrontDistributionDistributionConfigCacheBehaviorsList
	_jsii_.Get(
		j,
		"cacheBehaviors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) CacheBehaviorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheBehaviorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) Cnames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) CnamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cnamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ContinuousDeploymentPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"continuousDeploymentPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ContinuousDeploymentPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"continuousDeploymentPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) CustomErrorResponses() CloudfrontDistributionDistributionConfigCustomErrorResponsesList {
	var returns CloudfrontDistributionDistributionConfigCustomErrorResponsesList
	_jsii_.Get(
		j,
		"customErrorResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) CustomErrorResponsesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customErrorResponsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) CustomOrigin() CloudfrontDistributionDistributionConfigCustomOriginOutputReference {
	var returns CloudfrontDistributionDistributionConfigCustomOriginOutputReference
	_jsii_.Get(
		j,
		"customOrigin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) CustomOriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customOriginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) DefaultCacheBehavior() CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference {
	var returns CloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference
	_jsii_.Get(
		j,
		"defaultCacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) DefaultCacheBehaviorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultCacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) DefaultRootObject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) DefaultRootObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) HttpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) HttpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) Ipv6Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) Ipv6EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) Logging() CloudfrontDistributionDistributionConfigLoggingOutputReference {
	var returns CloudfrontDistributionDistributionConfigLoggingOutputReference
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) LoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) OriginGroups() CloudfrontDistributionDistributionConfigOriginGroupsOutputReference {
	var returns CloudfrontDistributionDistributionConfigOriginGroupsOutputReference
	_jsii_.Get(
		j,
		"originGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) OriginGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) Origins() CloudfrontDistributionDistributionConfigOriginsList {
	var returns CloudfrontDistributionDistributionConfigOriginsList
	_jsii_.Get(
		j,
		"origins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) OriginsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) PriceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) PriceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) Restrictions() CloudfrontDistributionDistributionConfigRestrictionsOutputReference {
	var returns CloudfrontDistributionDistributionConfigRestrictionsOutputReference
	_jsii_.Get(
		j,
		"restrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) RestrictionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) S3Origin() CloudfrontDistributionDistributionConfigS3OriginOutputReference {
	var returns CloudfrontDistributionDistributionConfigS3OriginOutputReference
	_jsii_.Get(
		j,
		"s3Origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) S3OriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3OriginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) Staging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) StagingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stagingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ViewerCertificate() CloudfrontDistributionDistributionConfigViewerCertificateOutputReference {
	var returns CloudfrontDistributionDistributionConfigViewerCertificateOutputReference
	_jsii_.Get(
		j,
		"viewerCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ViewerCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"viewerCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) WebAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) WebAclIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclIdInput",
		&returns,
	)
	return returns
}


func NewCloudfrontDistributionDistributionConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudfrontDistributionDistributionConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontDistributionDistributionConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference{}

	_jsii_.Create(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionDistributionConfigOutputReference_Override(c CloudfrontDistributionDistributionConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cloudfrontDistribution.CloudfrontDistributionDistributionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference)SetAliases(val *[]*string) {
	if err := j.validateSetAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliases",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference)SetCnames(val *[]*string) {
	if err := j.validateSetCnamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cnames",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference)SetContinuousDeploymentPolicyId(val *string) {
	if err := j.validateSetContinuousDeploymentPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"continuousDeploymentPolicyId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference)SetDefaultRootObject(val *string) {
	if err := j.validateSetDefaultRootObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRootObject",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference)SetHttpVersion(val *string) {
	if err := j.validateSetHttpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpVersion",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference)SetIpv6Enabled(val interface{}) {
	if err := j.validateSetIpv6EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Enabled",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference)SetPriceClass(val *string) {
	if err := j.validateSetPriceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priceClass",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference)SetStaging(val interface{}) {
	if err := j.validateSetStagingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staging",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference)SetWebAclId(val *string) {
	if err := j.validateSetWebAclIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webAclId",
		val,
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) PutCacheBehaviors(value interface{}) {
	if err := c.validatePutCacheBehaviorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCacheBehaviors",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) PutCustomErrorResponses(value interface{}) {
	if err := c.validatePutCustomErrorResponsesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomErrorResponses",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) PutCustomOrigin(value *CloudfrontDistributionDistributionConfigCustomOrigin) {
	if err := c.validatePutCustomOriginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomOrigin",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) PutDefaultCacheBehavior(value *CloudfrontDistributionDistributionConfigDefaultCacheBehavior) {
	if err := c.validatePutDefaultCacheBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDefaultCacheBehavior",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) PutLogging(value *CloudfrontDistributionDistributionConfigLogging) {
	if err := c.validatePutLoggingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLogging",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) PutOriginGroups(value *CloudfrontDistributionDistributionConfigOriginGroups) {
	if err := c.validatePutOriginGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOriginGroups",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) PutOrigins(value interface{}) {
	if err := c.validatePutOriginsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOrigins",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) PutRestrictions(value *CloudfrontDistributionDistributionConfigRestrictions) {
	if err := c.validatePutRestrictionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRestrictions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) PutS3Origin(value *CloudfrontDistributionDistributionConfigS3Origin) {
	if err := c.validatePutS3OriginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putS3Origin",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) PutViewerCertificate(value *CloudfrontDistributionDistributionConfigViewerCertificate) {
	if err := c.validatePutViewerCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putViewerCertificate",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetAliases() {
	_jsii_.InvokeVoid(
		c,
		"resetAliases",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetCacheBehaviors() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheBehaviors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetCnames() {
	_jsii_.InvokeVoid(
		c,
		"resetCnames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		c,
		"resetComment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetContinuousDeploymentPolicyId() {
	_jsii_.InvokeVoid(
		c,
		"resetContinuousDeploymentPolicyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetCustomErrorResponses() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomErrorResponses",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetCustomOrigin() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomOrigin",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetDefaultRootObject() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultRootObject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetHttpVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetIpv6Enabled() {
	_jsii_.InvokeVoid(
		c,
		"resetIpv6Enabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetLogging() {
	_jsii_.InvokeVoid(
		c,
		"resetLogging",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetOriginGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetOrigins() {
	_jsii_.InvokeVoid(
		c,
		"resetOrigins",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetPriceClass() {
	_jsii_.InvokeVoid(
		c,
		"resetPriceClass",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetRestrictions() {
	_jsii_.InvokeVoid(
		c,
		"resetRestrictions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetS3Origin() {
	_jsii_.InvokeVoid(
		c,
		"resetS3Origin",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetStaging() {
	_jsii_.InvokeVoid(
		c,
		"resetStaging",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetViewerCertificate() {
	_jsii_.InvokeVoid(
		c,
		"resetViewerCertificate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ResetWebAclId() {
	_jsii_.InvokeVoid(
		c,
		"resetWebAclId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

