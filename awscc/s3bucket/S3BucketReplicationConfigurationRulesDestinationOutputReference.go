package s3bucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/s3bucket/internal"
)

type S3BucketReplicationConfigurationRulesDestinationOutputReference interface {
	cdktf.ComplexObject
	AccessControlTranslation() S3BucketReplicationConfigurationRulesDestinationAccessControlTranslationOutputReference
	AccessControlTranslationInput() interface{}
	Account() *string
	SetAccount(val *string)
	AccountInput() *string
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
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
	EncryptionConfiguration() S3BucketReplicationConfigurationRulesDestinationEncryptionConfigurationOutputReference
	EncryptionConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *S3BucketReplicationConfigurationRulesDestination
	SetInternalValue(val *S3BucketReplicationConfigurationRulesDestination)
	Metrics() S3BucketReplicationConfigurationRulesDestinationMetricsOutputReference
	MetricsInput() interface{}
	ReplicationTime() S3BucketReplicationConfigurationRulesDestinationReplicationTimeOutputReference
	ReplicationTimeInput() interface{}
	StorageClass() *string
	SetStorageClass(val *string)
	StorageClassInput() *string
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
	PutAccessControlTranslation(value *S3BucketReplicationConfigurationRulesDestinationAccessControlTranslation)
	PutEncryptionConfiguration(value *S3BucketReplicationConfigurationRulesDestinationEncryptionConfiguration)
	PutMetrics(value *S3BucketReplicationConfigurationRulesDestinationMetrics)
	PutReplicationTime(value *S3BucketReplicationConfigurationRulesDestinationReplicationTime)
	ResetAccessControlTranslation()
	ResetAccount()
	ResetEncryptionConfiguration()
	ResetMetrics()
	ResetReplicationTime()
	ResetStorageClass()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3BucketReplicationConfigurationRulesDestinationOutputReference
type jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) AccessControlTranslation() S3BucketReplicationConfigurationRulesDestinationAccessControlTranslationOutputReference {
	var returns S3BucketReplicationConfigurationRulesDestinationAccessControlTranslationOutputReference
	_jsii_.Get(
		j,
		"accessControlTranslation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) AccessControlTranslationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlTranslationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) AccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) EncryptionConfiguration() S3BucketReplicationConfigurationRulesDestinationEncryptionConfigurationOutputReference {
	var returns S3BucketReplicationConfigurationRulesDestinationEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) EncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) InternalValue() *S3BucketReplicationConfigurationRulesDestination {
	var returns *S3BucketReplicationConfigurationRulesDestination
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) Metrics() S3BucketReplicationConfigurationRulesDestinationMetricsOutputReference {
	var returns S3BucketReplicationConfigurationRulesDestinationMetricsOutputReference
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) MetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) ReplicationTime() S3BucketReplicationConfigurationRulesDestinationReplicationTimeOutputReference {
	var returns S3BucketReplicationConfigurationRulesDestinationReplicationTimeOutputReference
	_jsii_.Get(
		j,
		"replicationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) ReplicationTimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3BucketReplicationConfigurationRulesDestinationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) S3BucketReplicationConfigurationRulesDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewS3BucketReplicationConfigurationRulesDestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference{}

	_jsii_.Create(
		"awscc.s3Bucket.S3BucketReplicationConfigurationRulesDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3BucketReplicationConfigurationRulesDestinationOutputReference_Override(s S3BucketReplicationConfigurationRulesDestinationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.s3Bucket.S3BucketReplicationConfigurationRulesDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference)SetAccount(val *string) {
	if err := j.validateSetAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"account",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference)SetInternalValue(val *S3BucketReplicationConfigurationRulesDestination) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference)SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) PutAccessControlTranslation(value *S3BucketReplicationConfigurationRulesDestinationAccessControlTranslation) {
	if err := s.validatePutAccessControlTranslationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAccessControlTranslation",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) PutEncryptionConfiguration(value *S3BucketReplicationConfigurationRulesDestinationEncryptionConfiguration) {
	if err := s.validatePutEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) PutMetrics(value *S3BucketReplicationConfigurationRulesDestinationMetrics) {
	if err := s.validatePutMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) PutReplicationTime(value *S3BucketReplicationConfigurationRulesDestinationReplicationTime) {
	if err := s.validatePutReplicationTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putReplicationTime",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) ResetAccessControlTranslation() {
	_jsii_.InvokeVoid(
		s,
		"resetAccessControlTranslation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) ResetAccount() {
	_jsii_.InvokeVoid(
		s,
		"resetAccount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) ResetEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetEncryptionConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) ResetMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) ResetReplicationTime() {
	_jsii_.InvokeVoid(
		s,
		"resetReplicationTime",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) ResetStorageClass() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageClass",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketReplicationConfigurationRulesDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

