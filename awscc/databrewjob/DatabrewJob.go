package databrewjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/databrewjob/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job awscc_databrew_job}.
type DatabrewJob interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DatabaseOutputs() DatabrewJobDatabaseOutputsList
	DatabaseOutputsInput() interface{}
	DataCatalogOutputs() DatabrewJobDataCatalogOutputsList
	DataCatalogOutputsInput() interface{}
	DatasetName() *string
	SetDatasetName(val *string)
	DatasetNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EncryptionKeyArn() *string
	SetEncryptionKeyArn(val *string)
	EncryptionKeyArnInput() *string
	EncryptionMode() *string
	SetEncryptionMode(val *string)
	EncryptionModeInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	JobSample() DatabrewJobJobSampleOutputReference
	JobSampleInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogSubscription() *string
	SetLogSubscription(val *string)
	LogSubscriptionInput() *string
	MaxCapacity() *float64
	SetMaxCapacity(val *float64)
	MaxCapacityInput() *float64
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OutputLocation() DatabrewJobOutputLocationOutputReference
	OutputLocationInput() interface{}
	Outputs() DatabrewJobOutputsList
	OutputsInput() interface{}
	ProfileConfiguration() DatabrewJobProfileConfigurationOutputReference
	ProfileConfigurationInput() interface{}
	ProjectName() *string
	SetProjectName(val *string)
	ProjectNameInput() *string
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
	Recipe() DatabrewJobRecipeOutputReference
	RecipeInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Tags() DatabrewJobTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	Type() *string
	SetType(val *string)
	TypeInput() *string
	ValidationConfigurations() DatabrewJobValidationConfigurationsList
	ValidationConfigurationsInput() interface{}
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutDatabaseOutputs(value interface{})
	PutDataCatalogOutputs(value interface{})
	PutJobSample(value *DatabrewJobJobSample)
	PutOutputLocation(value *DatabrewJobOutputLocation)
	PutOutputs(value interface{})
	PutProfileConfiguration(value *DatabrewJobProfileConfiguration)
	PutRecipe(value *DatabrewJobRecipe)
	PutTags(value interface{})
	PutValidationConfigurations(value interface{})
	ResetDatabaseOutputs()
	ResetDataCatalogOutputs()
	ResetDatasetName()
	ResetEncryptionKeyArn()
	ResetEncryptionMode()
	ResetJobSample()
	ResetLogSubscription()
	ResetMaxCapacity()
	ResetMaxRetries()
	ResetOutputLocation()
	ResetOutputs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProfileConfiguration()
	ResetProjectName()
	ResetRecipe()
	ResetTags()
	ResetTimeout()
	ResetValidationConfigurations()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DatabrewJob
type jsiiProxy_DatabrewJob struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatabrewJob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) DatabaseOutputs() DatabrewJobDatabaseOutputsList {
	var returns DatabrewJobDatabaseOutputsList
	_jsii_.Get(
		j,
		"databaseOutputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) DatabaseOutputsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseOutputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) DataCatalogOutputs() DatabrewJobDataCatalogOutputsList {
	var returns DatabrewJobDataCatalogOutputsList
	_jsii_.Get(
		j,
		"dataCatalogOutputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) DataCatalogOutputsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataCatalogOutputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) DatasetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) DatasetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) EncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) EncryptionKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) EncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) EncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) JobSample() DatabrewJobJobSampleOutputReference {
	var returns DatabrewJobJobSampleOutputReference
	_jsii_.Get(
		j,
		"jobSample",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) JobSampleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobSampleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) LogSubscription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logSubscription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) LogSubscriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logSubscriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) MaxCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) OutputLocation() DatabrewJobOutputLocationOutputReference {
	var returns DatabrewJobOutputLocationOutputReference
	_jsii_.Get(
		j,
		"outputLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) OutputLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) Outputs() DatabrewJobOutputsList {
	var returns DatabrewJobOutputsList
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) OutputsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) ProfileConfiguration() DatabrewJobProfileConfigurationOutputReference {
	var returns DatabrewJobProfileConfigurationOutputReference
	_jsii_.Get(
		j,
		"profileConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) ProfileConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) ProjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) Recipe() DatabrewJobRecipeOutputReference {
	var returns DatabrewJobRecipeOutputReference
	_jsii_.Get(
		j,
		"recipe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) RecipeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recipeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) Tags() DatabrewJobTagsList {
	var returns DatabrewJobTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) ValidationConfigurations() DatabrewJobValidationConfigurationsList {
	var returns DatabrewJobValidationConfigurationsList
	_jsii_.Get(
		j,
		"validationConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabrewJob) ValidationConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationConfigurationsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job awscc_databrew_job} Resource.
func NewDatabrewJob(scope constructs.Construct, id *string, config *DatabrewJobConfig) DatabrewJob {
	_init_.Initialize()

	if err := validateNewDatabrewJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabrewJob{}

	_jsii_.Create(
		"awscc.databrewJob.DatabrewJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job awscc_databrew_job} Resource.
func NewDatabrewJob_Override(d DatabrewJob, scope constructs.Construct, id *string, config *DatabrewJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.databrewJob.DatabrewJob",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatabrewJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatabrewJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatabrewJob)SetDatasetName(val *string) {
	if err := j.validateSetDatasetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetName",
		val,
	)
}

func (j *jsiiProxy_DatabrewJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatabrewJob)SetEncryptionKeyArn(val *string) {
	if err := j.validateSetEncryptionKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionKeyArn",
		val,
	)
}

func (j *jsiiProxy_DatabrewJob)SetEncryptionMode(val *string) {
	if err := j.validateSetEncryptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionMode",
		val,
	)
}

func (j *jsiiProxy_DatabrewJob)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatabrewJob)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatabrewJob)SetLogSubscription(val *string) {
	if err := j.validateSetLogSubscriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logSubscription",
		val,
	)
}

func (j *jsiiProxy_DatabrewJob)SetMaxCapacity(val *float64) {
	if err := j.validateSetMaxCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_DatabrewJob)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_DatabrewJob)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DatabrewJob)SetProjectName(val *string) {
	if err := j.validateSetProjectNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectName",
		val,
	)
}

func (j *jsiiProxy_DatabrewJob)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatabrewJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatabrewJob)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_DatabrewJob)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_DatabrewJob)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a DatabrewJob resource upon running "cdktf plan <stack-name>".
func DatabrewJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatabrewJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.databrewJob.DatabrewJob",
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
func DatabrewJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabrewJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.databrewJob.DatabrewJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabrewJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabrewJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.databrewJob.DatabrewJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabrewJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabrewJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.databrewJob.DatabrewJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatabrewJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.databrewJob.DatabrewJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabrewJob) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatabrewJob) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatabrewJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabrewJob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabrewJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabrewJob) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabrewJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabrewJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabrewJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabrewJob) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabrewJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabrewJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabrewJob) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatabrewJob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabrewJob) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabrewJob) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatabrewJob) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabrewJob) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatabrewJob) PutDatabaseOutputs(value interface{}) {
	if err := d.validatePutDatabaseOutputsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDatabaseOutputs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewJob) PutDataCatalogOutputs(value interface{}) {
	if err := d.validatePutDataCatalogOutputsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDataCatalogOutputs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewJob) PutJobSample(value *DatabrewJobJobSample) {
	if err := d.validatePutJobSampleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJobSample",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewJob) PutOutputLocation(value *DatabrewJobOutputLocation) {
	if err := d.validatePutOutputLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOutputLocation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewJob) PutOutputs(value interface{}) {
	if err := d.validatePutOutputsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOutputs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewJob) PutProfileConfiguration(value *DatabrewJobProfileConfiguration) {
	if err := d.validatePutProfileConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProfileConfiguration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewJob) PutRecipe(value *DatabrewJobRecipe) {
	if err := d.validatePutRecipeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRecipe",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewJob) PutTags(value interface{}) {
	if err := d.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewJob) PutValidationConfigurations(value interface{}) {
	if err := d.validatePutValidationConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putValidationConfigurations",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabrewJob) ResetDatabaseOutputs() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseOutputs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) ResetDataCatalogOutputs() {
	_jsii_.InvokeVoid(
		d,
		"resetDataCatalogOutputs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) ResetDatasetName() {
	_jsii_.InvokeVoid(
		d,
		"resetDatasetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) ResetEncryptionKeyArn() {
	_jsii_.InvokeVoid(
		d,
		"resetEncryptionKeyArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) ResetEncryptionMode() {
	_jsii_.InvokeVoid(
		d,
		"resetEncryptionMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) ResetJobSample() {
	_jsii_.InvokeVoid(
		d,
		"resetJobSample",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) ResetLogSubscription() {
	_jsii_.InvokeVoid(
		d,
		"resetLogSubscription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) ResetMaxCapacity() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxCapacity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) ResetOutputLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetOutputLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) ResetOutputs() {
	_jsii_.InvokeVoid(
		d,
		"resetOutputs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) ResetProfileConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetProfileConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) ResetProjectName() {
	_jsii_.InvokeVoid(
		d,
		"resetProjectName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) ResetRecipe() {
	_jsii_.InvokeVoid(
		d,
		"resetRecipe",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) ResetTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) ResetValidationConfigurations() {
	_jsii_.InvokeVoid(
		d,
		"resetValidationConfigurations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabrewJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabrewJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabrewJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabrewJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabrewJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabrewJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

