package dataawscccloudfrontdistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscccloudfrontdistribution/internal"
)

type DataAwsccCloudfrontDistributionDistributionConfigOutputReference interface {
	cdktf.ComplexObject
	Aliases() *[]*string
	CacheBehaviors() DataAwsccCloudfrontDistributionDistributionConfigCacheBehaviorsList
	Cnames() *[]*string
	Comment() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomErrorResponses() DataAwsccCloudfrontDistributionDistributionConfigCustomErrorResponsesList
	CustomOrigin() DataAwsccCloudfrontDistributionDistributionConfigCustomOriginOutputReference
	DefaultCacheBehavior() DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference
	DefaultRootObject() *string
	Enabled() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	HttpVersion() *string
	InternalValue() *DataAwsccCloudfrontDistributionDistributionConfig
	SetInternalValue(val *DataAwsccCloudfrontDistributionDistributionConfig)
	Ipv6Enabled() cdktf.IResolvable
	Logging() DataAwsccCloudfrontDistributionDistributionConfigLoggingOutputReference
	OriginGroups() DataAwsccCloudfrontDistributionDistributionConfigOriginGroupsOutputReference
	Origins() DataAwsccCloudfrontDistributionDistributionConfigOriginsList
	PriceClass() *string
	Restrictions() DataAwsccCloudfrontDistributionDistributionConfigRestrictionsOutputReference
	S3Origin() DataAwsccCloudfrontDistributionDistributionConfigS3OriginOutputReference
	Staging() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ViewerCertificate() DataAwsccCloudfrontDistributionDistributionConfigViewerCertificateOutputReference
	WebAclId() *string
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

// The jsii proxy struct for DataAwsccCloudfrontDistributionDistributionConfigOutputReference
type jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) CacheBehaviors() DataAwsccCloudfrontDistributionDistributionConfigCacheBehaviorsList {
	var returns DataAwsccCloudfrontDistributionDistributionConfigCacheBehaviorsList
	_jsii_.Get(
		j,
		"cacheBehaviors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) Cnames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) ContinuousDeploymentPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"continuousDeploymentPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) CustomErrorResponses() DataAwsccCloudfrontDistributionDistributionConfigCustomErrorResponsesList {
	var returns DataAwsccCloudfrontDistributionDistributionConfigCustomErrorResponsesList
	_jsii_.Get(
		j,
		"customErrorResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) CustomOrigin() DataAwsccCloudfrontDistributionDistributionConfigCustomOriginOutputReference {
	var returns DataAwsccCloudfrontDistributionDistributionConfigCustomOriginOutputReference
	_jsii_.Get(
		j,
		"customOrigin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) DefaultCacheBehavior() DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference {
	var returns DataAwsccCloudfrontDistributionDistributionConfigDefaultCacheBehaviorOutputReference
	_jsii_.Get(
		j,
		"defaultCacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) DefaultRootObject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) HttpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) InternalValue() *DataAwsccCloudfrontDistributionDistributionConfig {
	var returns *DataAwsccCloudfrontDistributionDistributionConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) Ipv6Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ipv6Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) Logging() DataAwsccCloudfrontDistributionDistributionConfigLoggingOutputReference {
	var returns DataAwsccCloudfrontDistributionDistributionConfigLoggingOutputReference
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) OriginGroups() DataAwsccCloudfrontDistributionDistributionConfigOriginGroupsOutputReference {
	var returns DataAwsccCloudfrontDistributionDistributionConfigOriginGroupsOutputReference
	_jsii_.Get(
		j,
		"originGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) Origins() DataAwsccCloudfrontDistributionDistributionConfigOriginsList {
	var returns DataAwsccCloudfrontDistributionDistributionConfigOriginsList
	_jsii_.Get(
		j,
		"origins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) PriceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) Restrictions() DataAwsccCloudfrontDistributionDistributionConfigRestrictionsOutputReference {
	var returns DataAwsccCloudfrontDistributionDistributionConfigRestrictionsOutputReference
	_jsii_.Get(
		j,
		"restrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) S3Origin() DataAwsccCloudfrontDistributionDistributionConfigS3OriginOutputReference {
	var returns DataAwsccCloudfrontDistributionDistributionConfigS3OriginOutputReference
	_jsii_.Get(
		j,
		"s3Origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) Staging() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"staging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) ViewerCertificate() DataAwsccCloudfrontDistributionDistributionConfigViewerCertificateOutputReference {
	var returns DataAwsccCloudfrontDistributionDistributionConfigViewerCertificateOutputReference
	_jsii_.Get(
		j,
		"viewerCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) WebAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclId",
		&returns,
	)
	return returns
}


func NewDataAwsccCloudfrontDistributionDistributionConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccCloudfrontDistributionDistributionConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccCloudfrontDistributionDistributionConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccCloudfrontDistribution.DataAwsccCloudfrontDistributionDistributionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccCloudfrontDistributionDistributionConfigOutputReference_Override(d DataAwsccCloudfrontDistributionDistributionConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccCloudfrontDistribution.DataAwsccCloudfrontDistributionDistributionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference)SetInternalValue(val *DataAwsccCloudfrontDistributionDistributionConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccCloudfrontDistributionDistributionConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

