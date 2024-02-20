package cloudtrailtrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cloudtrailtrail/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudtrail_trail awscc_cloudtrail_trail}.
type CloudtrailTrail interface {
	cdktf.TerraformResource
	AdvancedEventSelectors() CloudtrailTrailAdvancedEventSelectorsList
	AdvancedEventSelectorsInput() interface{}
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudwatchLogsLogGroupArn() *string
	SetCloudwatchLogsLogGroupArn(val *string)
	CloudwatchLogsLogGroupArnInput() *string
	CloudwatchLogsRoleArn() *string
	SetCloudwatchLogsRoleArn(val *string)
	CloudwatchLogsRoleArnInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableLogFileValidation() interface{}
	SetEnableLogFileValidation(val interface{})
	EnableLogFileValidationInput() interface{}
	EventSelectors() CloudtrailTrailEventSelectorsList
	EventSelectorsInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IncludeGlobalServiceEvents() interface{}
	SetIncludeGlobalServiceEvents(val interface{})
	IncludeGlobalServiceEventsInput() interface{}
	InsightSelectors() CloudtrailTrailInsightSelectorsList
	InsightSelectorsInput() interface{}
	IsLogging() interface{}
	SetIsLogging(val interface{})
	IsLoggingInput() interface{}
	IsMultiRegionTrail() interface{}
	SetIsMultiRegionTrail(val interface{})
	IsMultiRegionTrailInput() interface{}
	IsOrganizationTrail() interface{}
	SetIsOrganizationTrail(val interface{})
	IsOrganizationTrailInput() interface{}
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	S3BucketName() *string
	SetS3BucketName(val *string)
	S3BucketNameInput() *string
	S3KeyPrefix() *string
	SetS3KeyPrefix(val *string)
	S3KeyPrefixInput() *string
	SnsTopicArn() *string
	SnsTopicName() *string
	SetSnsTopicName(val *string)
	SnsTopicNameInput() *string
	Tags() CloudtrailTrailTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TrailName() *string
	SetTrailName(val *string)
	TrailNameInput() *string
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
	PutAdvancedEventSelectors(value interface{})
	PutEventSelectors(value interface{})
	PutInsightSelectors(value interface{})
	PutTags(value interface{})
	ResetAdvancedEventSelectors()
	ResetCloudwatchLogsLogGroupArn()
	ResetCloudwatchLogsRoleArn()
	ResetEnableLogFileValidation()
	ResetEventSelectors()
	ResetIncludeGlobalServiceEvents()
	ResetInsightSelectors()
	ResetIsMultiRegionTrail()
	ResetIsOrganizationTrail()
	ResetKmsKeyId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetS3KeyPrefix()
	ResetSnsTopicName()
	ResetTags()
	ResetTrailName()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudtrailTrail
type jsiiProxy_CloudtrailTrail struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudtrailTrail) AdvancedEventSelectors() CloudtrailTrailAdvancedEventSelectorsList {
	var returns CloudtrailTrailAdvancedEventSelectorsList
	_jsii_.Get(
		j,
		"advancedEventSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) AdvancedEventSelectorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedEventSelectorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) CloudwatchLogsLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogsLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) CloudwatchLogsLogGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogsLogGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) CloudwatchLogsRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) CloudwatchLogsRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogsRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) EnableLogFileValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogFileValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) EnableLogFileValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogFileValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) EventSelectors() CloudtrailTrailEventSelectorsList {
	var returns CloudtrailTrailEventSelectorsList
	_jsii_.Get(
		j,
		"eventSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) EventSelectorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventSelectorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) IncludeGlobalServiceEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalServiceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) IncludeGlobalServiceEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalServiceEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) InsightSelectors() CloudtrailTrailInsightSelectorsList {
	var returns CloudtrailTrailInsightSelectorsList
	_jsii_.Get(
		j,
		"insightSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) InsightSelectorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insightSelectorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) IsLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) IsLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) IsMultiRegionTrail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMultiRegionTrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) IsMultiRegionTrailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMultiRegionTrailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) IsOrganizationTrail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isOrganizationTrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) IsOrganizationTrailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isOrganizationTrailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) S3BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) S3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) S3KeyPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) SnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) SnsTopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) SnsTopicNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) Tags() CloudtrailTrailTagsList {
	var returns CloudtrailTrailTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) TrailName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trailName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailTrail) TrailNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trailNameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudtrail_trail awscc_cloudtrail_trail} Resource.
func NewCloudtrailTrail(scope constructs.Construct, id *string, config *CloudtrailTrailConfig) CloudtrailTrail {
	_init_.Initialize()

	if err := validateNewCloudtrailTrailParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudtrailTrail{}

	_jsii_.Create(
		"awscc.cloudtrailTrail.CloudtrailTrail",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudtrail_trail awscc_cloudtrail_trail} Resource.
func NewCloudtrailTrail_Override(c CloudtrailTrail, scope constructs.Construct, id *string, config *CloudtrailTrailConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cloudtrailTrail.CloudtrailTrail",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetCloudwatchLogsLogGroupArn(val *string) {
	if err := j.validateSetCloudwatchLogsLogGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchLogsLogGroupArn",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetCloudwatchLogsRoleArn(val *string) {
	if err := j.validateSetCloudwatchLogsRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchLogsRoleArn",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetEnableLogFileValidation(val interface{}) {
	if err := j.validateSetEnableLogFileValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLogFileValidation",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetIncludeGlobalServiceEvents(val interface{}) {
	if err := j.validateSetIncludeGlobalServiceEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeGlobalServiceEvents",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetIsLogging(val interface{}) {
	if err := j.validateSetIsLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isLogging",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetIsMultiRegionTrail(val interface{}) {
	if err := j.validateSetIsMultiRegionTrailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isMultiRegionTrail",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetIsOrganizationTrail(val interface{}) {
	if err := j.validateSetIsOrganizationTrailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isOrganizationTrail",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetS3BucketName(val *string) {
	if err := j.validateSetS3BucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetS3KeyPrefix(val *string) {
	if err := j.validateSetS3KeyPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetSnsTopicName(val *string) {
	if err := j.validateSetSnsTopicNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snsTopicName",
		val,
	)
}

func (j *jsiiProxy_CloudtrailTrail)SetTrailName(val *string) {
	if err := j.validateSetTrailNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trailName",
		val,
	)
}

// Generates CDKTF code for importing a CloudtrailTrail resource upon running "cdktf plan <stack-name>".
func CloudtrailTrail_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCloudtrailTrail_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.cloudtrailTrail.CloudtrailTrail",
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
func CloudtrailTrail_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudtrailTrail_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cloudtrailTrail.CloudtrailTrail",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudtrailTrail_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudtrailTrail_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cloudtrailTrail.CloudtrailTrail",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudtrailTrail_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudtrailTrail_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cloudtrailTrail.CloudtrailTrail",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudtrailTrail_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.cloudtrailTrail.CloudtrailTrail",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudtrailTrail) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CloudtrailTrail) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudtrailTrail) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudtrailTrail) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudtrailTrail) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudtrailTrail) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudtrailTrail) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudtrailTrail) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudtrailTrail) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudtrailTrail) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudtrailTrail) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudtrailTrail) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CloudtrailTrail) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailTrail) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CloudtrailTrail) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudtrailTrail) PutAdvancedEventSelectors(value interface{}) {
	if err := c.validatePutAdvancedEventSelectorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAdvancedEventSelectors",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudtrailTrail) PutEventSelectors(value interface{}) {
	if err := c.validatePutEventSelectorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEventSelectors",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudtrailTrail) PutInsightSelectors(value interface{}) {
	if err := c.validatePutInsightSelectorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInsightSelectors",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudtrailTrail) PutTags(value interface{}) {
	if err := c.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTags",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudtrailTrail) ResetAdvancedEventSelectors() {
	_jsii_.InvokeVoid(
		c,
		"resetAdvancedEventSelectors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrail) ResetCloudwatchLogsLogGroupArn() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudwatchLogsLogGroupArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrail) ResetCloudwatchLogsRoleArn() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudwatchLogsRoleArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrail) ResetEnableLogFileValidation() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableLogFileValidation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrail) ResetEventSelectors() {
	_jsii_.InvokeVoid(
		c,
		"resetEventSelectors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrail) ResetIncludeGlobalServiceEvents() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeGlobalServiceEvents",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrail) ResetInsightSelectors() {
	_jsii_.InvokeVoid(
		c,
		"resetInsightSelectors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrail) ResetIsMultiRegionTrail() {
	_jsii_.InvokeVoid(
		c,
		"resetIsMultiRegionTrail",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrail) ResetIsOrganizationTrail() {
	_jsii_.InvokeVoid(
		c,
		"resetIsOrganizationTrail",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrail) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		c,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrail) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrail) ResetS3KeyPrefix() {
	_jsii_.InvokeVoid(
		c,
		"resetS3KeyPrefix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrail) ResetSnsTopicName() {
	_jsii_.InvokeVoid(
		c,
		"resetSnsTopicName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrail) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrail) ResetTrailName() {
	_jsii_.InvokeVoid(
		c,
		"resetTrailName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailTrail) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailTrail) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailTrail) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailTrail) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

