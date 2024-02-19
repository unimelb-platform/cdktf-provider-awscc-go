package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/kinesisfirehosedeliverystream/internal"
)

type KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference interface {
	cdktf.ComplexObject
	CloudwatchLoggingOptions() KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCloudwatchLoggingOptionsOutputReference
	CloudwatchLoggingOptionsInput() interface{}
	ClusterJdbcurl() *string
	SetClusterJdbcurl(val *string)
	ClusterJdbcurlInput() *string
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
	CopyCommand() KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCopyCommandOutputReference
	CopyCommandInput() *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCopyCommand
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	ProcessingConfiguration() KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationOutputReference
	ProcessingConfigurationInput() interface{}
	RetryOptions() KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationRetryOptionsOutputReference
	RetryOptionsInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	S3BackupConfiguration() KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationOutputReference
	S3BackupConfigurationInput() interface{}
	S3BackupMode() *string
	SetS3BackupMode(val *string)
	S3BackupModeInput() *string
	S3Configuration() KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationOutputReference
	S3ConfigurationInput() *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3Configuration
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	PutCloudwatchLoggingOptions(value *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCloudwatchLoggingOptions)
	PutCopyCommand(value *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCopyCommand)
	PutProcessingConfiguration(value *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfiguration)
	PutRetryOptions(value *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationRetryOptions)
	PutS3BackupConfiguration(value *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfiguration)
	PutS3Configuration(value *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3Configuration)
	ResetCloudwatchLoggingOptions()
	ResetProcessingConfiguration()
	ResetRetryOptions()
	ResetS3BackupConfiguration()
	ResetS3BackupMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference
type jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) CloudwatchLoggingOptions() KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCloudwatchLoggingOptionsOutputReference {
	var returns KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCloudwatchLoggingOptionsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) CloudwatchLoggingOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) ClusterJdbcurl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterJdbcurl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) ClusterJdbcurlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterJdbcurlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) CopyCommand() KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCopyCommandOutputReference {
	var returns KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCopyCommandOutputReference
	_jsii_.Get(
		j,
		"copyCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) CopyCommandInput() *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCopyCommand {
	var returns *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCopyCommand
	_jsii_.Get(
		j,
		"copyCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) ProcessingConfiguration() KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfigurationOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) ProcessingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) RetryOptions() KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationRetryOptionsOutputReference {
	var returns KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationRetryOptionsOutputReference
	_jsii_.Get(
		j,
		"retryOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) RetryOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) S3BackupConfiguration() KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3BackupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) S3BackupConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3BackupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) S3BackupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) S3BackupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) S3Configuration() KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3ConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) S3ConfigurationInput() *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3Configuration {
	var returns *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3Configuration
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewKinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference_Override(k KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference)SetClusterJdbcurl(val *string) {
	if err := j.validateSetClusterJdbcurlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterJdbcurl",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference)SetS3BackupMode(val *string) {
	if err := j.validateSetS3BackupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BackupMode",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) PutCloudwatchLoggingOptions(value *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCloudwatchLoggingOptions) {
	if err := k.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) PutCopyCommand(value *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCopyCommand) {
	if err := k.validatePutCopyCommandParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCopyCommand",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) PutProcessingConfiguration(value *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationProcessingConfiguration) {
	if err := k.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) PutRetryOptions(value *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationRetryOptions) {
	if err := k.validatePutRetryOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putRetryOptions",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) PutS3BackupConfiguration(value *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3BackupConfiguration) {
	if err := k.validatePutS3BackupConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putS3BackupConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) PutS3Configuration(value *KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationS3Configuration) {
	if err := k.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		k,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) ResetRetryOptions() {
	_jsii_.InvokeVoid(
		k,
		"resetRetryOptions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) ResetS3BackupConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetS3BackupConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) ResetS3BackupMode() {
	_jsii_.InvokeVoid(
		k,
		"resetS3BackupMode",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

