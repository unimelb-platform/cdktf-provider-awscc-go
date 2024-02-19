package s3bucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/s3bucket/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket awscc_s3_bucket}.
type S3Bucket interface {
	cdktf.TerraformResource
	AccelerateConfiguration() S3BucketAccelerateConfigurationOutputReference
	AccelerateConfigurationInput() interface{}
	AccessControl() *string
	SetAccessControl(val *string)
	AccessControlInput() *string
	AnalyticsConfigurations() S3BucketAnalyticsConfigurationsList
	AnalyticsConfigurationsInput() interface{}
	Arn() *string
	BucketEncryption() S3BucketBucketEncryptionOutputReference
	BucketEncryptionInput() interface{}
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CorsConfiguration() S3BucketCorsConfigurationOutputReference
	CorsConfigurationInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DomainName() *string
	DualStackDomainName() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IntelligentTieringConfigurations() S3BucketIntelligentTieringConfigurationsList
	IntelligentTieringConfigurationsInput() interface{}
	InventoryConfigurations() S3BucketInventoryConfigurationsList
	InventoryConfigurationsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LifecycleConfiguration() S3BucketLifecycleConfigurationOutputReference
	LifecycleConfigurationInput() interface{}
	LoggingConfiguration() S3BucketLoggingConfigurationOutputReference
	LoggingConfigurationInput() interface{}
	MetricsConfigurations() S3BucketMetricsConfigurationsList
	MetricsConfigurationsInput() interface{}
	// The tree node.
	Node() constructs.Node
	NotificationConfiguration() S3BucketNotificationConfigurationOutputReference
	NotificationConfigurationInput() interface{}
	ObjectLockConfiguration() S3BucketObjectLockConfigurationOutputReference
	ObjectLockConfigurationInput() interface{}
	ObjectLockEnabled() interface{}
	SetObjectLockEnabled(val interface{})
	ObjectLockEnabledInput() interface{}
	OwnershipControls() S3BucketOwnershipControlsOutputReference
	OwnershipControlsInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicAccessBlockConfiguration() S3BucketPublicAccessBlockConfigurationOutputReference
	PublicAccessBlockConfigurationInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	RegionalDomainName() *string
	ReplicationConfiguration() S3BucketReplicationConfigurationOutputReference
	ReplicationConfigurationInput() interface{}
	Tags() S3BucketTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VersioningConfiguration() S3BucketVersioningConfigurationOutputReference
	VersioningConfigurationInput() interface{}
	WebsiteConfiguration() S3BucketWebsiteConfigurationOutputReference
	WebsiteConfigurationInput() interface{}
	WebsiteUrl() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAccelerateConfiguration(value *S3BucketAccelerateConfiguration)
	PutAnalyticsConfigurations(value interface{})
	PutBucketEncryption(value *S3BucketBucketEncryption)
	PutCorsConfiguration(value *S3BucketCorsConfiguration)
	PutIntelligentTieringConfigurations(value interface{})
	PutInventoryConfigurations(value interface{})
	PutLifecycleConfiguration(value *S3BucketLifecycleConfiguration)
	PutLoggingConfiguration(value *S3BucketLoggingConfiguration)
	PutMetricsConfigurations(value interface{})
	PutNotificationConfiguration(value *S3BucketNotificationConfiguration)
	PutObjectLockConfiguration(value *S3BucketObjectLockConfiguration)
	PutOwnershipControls(value *S3BucketOwnershipControls)
	PutPublicAccessBlockConfiguration(value *S3BucketPublicAccessBlockConfiguration)
	PutReplicationConfiguration(value *S3BucketReplicationConfiguration)
	PutTags(value interface{})
	PutVersioningConfiguration(value *S3BucketVersioningConfiguration)
	PutWebsiteConfiguration(value *S3BucketWebsiteConfiguration)
	ResetAccelerateConfiguration()
	ResetAccessControl()
	ResetAnalyticsConfigurations()
	ResetBucketEncryption()
	ResetBucketName()
	ResetCorsConfiguration()
	ResetIntelligentTieringConfigurations()
	ResetInventoryConfigurations()
	ResetLifecycleConfiguration()
	ResetLoggingConfiguration()
	ResetMetricsConfigurations()
	ResetNotificationConfiguration()
	ResetObjectLockConfiguration()
	ResetObjectLockEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwnershipControls()
	ResetPublicAccessBlockConfiguration()
	ResetReplicationConfiguration()
	ResetTags()
	ResetVersioningConfiguration()
	ResetWebsiteConfiguration()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for S3Bucket
type jsiiProxy_S3Bucket struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_S3Bucket) AccelerateConfiguration() S3BucketAccelerateConfigurationOutputReference {
	var returns S3BucketAccelerateConfigurationOutputReference
	_jsii_.Get(
		j,
		"accelerateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) AccelerateConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accelerateConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) AccessControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) AccessControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) AnalyticsConfigurations() S3BucketAnalyticsConfigurationsList {
	var returns S3BucketAnalyticsConfigurationsList
	_jsii_.Get(
		j,
		"analyticsConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) AnalyticsConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"analyticsConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) BucketEncryption() S3BucketBucketEncryptionOutputReference {
	var returns S3BucketBucketEncryptionOutputReference
	_jsii_.Get(
		j,
		"bucketEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) BucketEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bucketEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) CorsConfiguration() S3BucketCorsConfigurationOutputReference {
	var returns S3BucketCorsConfigurationOutputReference
	_jsii_.Get(
		j,
		"corsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) CorsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"corsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) DualStackDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dualStackDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) IntelligentTieringConfigurations() S3BucketIntelligentTieringConfigurationsList {
	var returns S3BucketIntelligentTieringConfigurationsList
	_jsii_.Get(
		j,
		"intelligentTieringConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) IntelligentTieringConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intelligentTieringConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) InventoryConfigurations() S3BucketInventoryConfigurationsList {
	var returns S3BucketInventoryConfigurationsList
	_jsii_.Get(
		j,
		"inventoryConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) InventoryConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inventoryConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) LifecycleConfiguration() S3BucketLifecycleConfigurationOutputReference {
	var returns S3BucketLifecycleConfigurationOutputReference
	_jsii_.Get(
		j,
		"lifecycleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) LifecycleConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecycleConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) LoggingConfiguration() S3BucketLoggingConfigurationOutputReference {
	var returns S3BucketLoggingConfigurationOutputReference
	_jsii_.Get(
		j,
		"loggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) LoggingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) MetricsConfigurations() S3BucketMetricsConfigurationsList {
	var returns S3BucketMetricsConfigurationsList
	_jsii_.Get(
		j,
		"metricsConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) MetricsConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) NotificationConfiguration() S3BucketNotificationConfigurationOutputReference {
	var returns S3BucketNotificationConfigurationOutputReference
	_jsii_.Get(
		j,
		"notificationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) NotificationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ObjectLockConfiguration() S3BucketObjectLockConfigurationOutputReference {
	var returns S3BucketObjectLockConfigurationOutputReference
	_jsii_.Get(
		j,
		"objectLockConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ObjectLockConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"objectLockConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ObjectLockEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"objectLockEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ObjectLockEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"objectLockEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) OwnershipControls() S3BucketOwnershipControlsOutputReference {
	var returns S3BucketOwnershipControlsOutputReference
	_jsii_.Get(
		j,
		"ownershipControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) OwnershipControlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ownershipControlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) PublicAccessBlockConfiguration() S3BucketPublicAccessBlockConfigurationOutputReference {
	var returns S3BucketPublicAccessBlockConfigurationOutputReference
	_jsii_.Get(
		j,
		"publicAccessBlockConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) PublicAccessBlockConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicAccessBlockConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) RegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ReplicationConfiguration() S3BucketReplicationConfigurationOutputReference {
	var returns S3BucketReplicationConfigurationOutputReference
	_jsii_.Get(
		j,
		"replicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ReplicationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Tags() S3BucketTagsList {
	var returns S3BucketTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) VersioningConfiguration() S3BucketVersioningConfigurationOutputReference {
	var returns S3BucketVersioningConfigurationOutputReference
	_jsii_.Get(
		j,
		"versioningConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) VersioningConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versioningConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) WebsiteConfiguration() S3BucketWebsiteConfigurationOutputReference {
	var returns S3BucketWebsiteConfigurationOutputReference
	_jsii_.Get(
		j,
		"websiteConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) WebsiteConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"websiteConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) WebsiteUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteUrl",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket awscc_s3_bucket} Resource.
func NewS3Bucket(scope constructs.Construct, id *string, config *S3BucketConfig) S3Bucket {
	_init_.Initialize()

	if err := validateNewS3BucketParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3Bucket{}

	_jsii_.Create(
		"awscc.s3Bucket.S3Bucket",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket awscc_s3_bucket} Resource.
func NewS3Bucket_Override(s S3Bucket, scope constructs.Construct, id *string, config *S3BucketConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.s3Bucket.S3Bucket",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_S3Bucket)SetAccessControl(val *string) {
	if err := j.validateSetAccessControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessControl",
		val,
	)
}

func (j *jsiiProxy_S3Bucket)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_S3Bucket)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_S3Bucket)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_S3Bucket)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_S3Bucket)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_S3Bucket)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_S3Bucket)SetObjectLockEnabled(val interface{}) {
	if err := j.validateSetObjectLockEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockEnabled",
		val,
	)
}

func (j *jsiiProxy_S3Bucket)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_S3Bucket)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a S3Bucket resource upon running "cdktf plan <stack-name>".
func S3Bucket_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateS3Bucket_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.s3Bucket.S3Bucket",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func S3Bucket_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3Bucket_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.s3Bucket.S3Bucket",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func S3Bucket_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3Bucket_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.s3Bucket.S3Bucket",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func S3Bucket_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3Bucket_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.s3Bucket.S3Bucket",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func S3Bucket_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.s3Bucket.S3Bucket",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_S3Bucket) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_S3Bucket) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_S3Bucket) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3Bucket) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3Bucket) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3Bucket) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3Bucket) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3Bucket) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3Bucket) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3Bucket) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3Bucket) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3Bucket) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_S3Bucket) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_S3Bucket) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_S3Bucket) PutAccelerateConfiguration(value *S3BucketAccelerateConfiguration) {
	if err := s.validatePutAccelerateConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAccelerateConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutAnalyticsConfigurations(value interface{}) {
	if err := s.validatePutAnalyticsConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAnalyticsConfigurations",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutBucketEncryption(value *S3BucketBucketEncryption) {
	if err := s.validatePutBucketEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBucketEncryption",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutCorsConfiguration(value *S3BucketCorsConfiguration) {
	if err := s.validatePutCorsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCorsConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutIntelligentTieringConfigurations(value interface{}) {
	if err := s.validatePutIntelligentTieringConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIntelligentTieringConfigurations",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutInventoryConfigurations(value interface{}) {
	if err := s.validatePutInventoryConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInventoryConfigurations",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutLifecycleConfiguration(value *S3BucketLifecycleConfiguration) {
	if err := s.validatePutLifecycleConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLifecycleConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutLoggingConfiguration(value *S3BucketLoggingConfiguration) {
	if err := s.validatePutLoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLoggingConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutMetricsConfigurations(value interface{}) {
	if err := s.validatePutMetricsConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMetricsConfigurations",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutNotificationConfiguration(value *S3BucketNotificationConfiguration) {
	if err := s.validatePutNotificationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNotificationConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutObjectLockConfiguration(value *S3BucketObjectLockConfiguration) {
	if err := s.validatePutObjectLockConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putObjectLockConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutOwnershipControls(value *S3BucketOwnershipControls) {
	if err := s.validatePutOwnershipControlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOwnershipControls",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutPublicAccessBlockConfiguration(value *S3BucketPublicAccessBlockConfiguration) {
	if err := s.validatePutPublicAccessBlockConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPublicAccessBlockConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutReplicationConfiguration(value *S3BucketReplicationConfiguration) {
	if err := s.validatePutReplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putReplicationConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutTags(value interface{}) {
	if err := s.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTags",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutVersioningConfiguration(value *S3BucketVersioningConfiguration) {
	if err := s.validatePutVersioningConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVersioningConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutWebsiteConfiguration(value *S3BucketWebsiteConfiguration) {
	if err := s.validatePutWebsiteConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putWebsiteConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3Bucket) ResetAccelerateConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetAccelerateConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetAccessControl() {
	_jsii_.InvokeVoid(
		s,
		"resetAccessControl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetAnalyticsConfigurations() {
	_jsii_.InvokeVoid(
		s,
		"resetAnalyticsConfigurations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetBucketEncryption() {
	_jsii_.InvokeVoid(
		s,
		"resetBucketEncryption",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetBucketName() {
	_jsii_.InvokeVoid(
		s,
		"resetBucketName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetCorsConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetCorsConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetIntelligentTieringConfigurations() {
	_jsii_.InvokeVoid(
		s,
		"resetIntelligentTieringConfigurations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetInventoryConfigurations() {
	_jsii_.InvokeVoid(
		s,
		"resetInventoryConfigurations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetLifecycleConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetLifecycleConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetLoggingConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetLoggingConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetMetricsConfigurations() {
	_jsii_.InvokeVoid(
		s,
		"resetMetricsConfigurations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetNotificationConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetNotificationConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetObjectLockConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectLockConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetObjectLockEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectLockEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetOwnershipControls() {
	_jsii_.InvokeVoid(
		s,
		"resetOwnershipControls",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetPublicAccessBlockConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetPublicAccessBlockConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetReplicationConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetReplicationConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetVersioningConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetVersioningConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetWebsiteConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetWebsiteConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

