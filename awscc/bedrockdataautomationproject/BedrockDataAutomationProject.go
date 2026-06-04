package bedrockdataautomationproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/bedrockdataautomationproject/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_data_automation_project awscc_bedrock_data_automation_project}.
type BedrockDataAutomationProject interface {
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
	CreationTime() *string
	CustomOutputConfiguration() BedrockDataAutomationProjectCustomOutputConfigurationOutputReference
	CustomOutputConfigurationInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	KmsEncryptionContext() *map[string]*string
	SetKmsEncryptionContext(val *map[string]*string)
	KmsEncryptionContextInput() *map[string]*string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	LastModifiedTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OverrideConfiguration() BedrockDataAutomationProjectOverrideConfigurationOutputReference
	OverrideConfigurationInput() interface{}
	ProjectArn() *string
	ProjectDescription() *string
	SetProjectDescription(val *string)
	ProjectDescriptionInput() *string
	ProjectName() *string
	SetProjectName(val *string)
	ProjectNameInput() *string
	ProjectStage() *string
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
	StandardOutputConfiguration() BedrockDataAutomationProjectStandardOutputConfigurationOutputReference
	StandardOutputConfigurationInput() interface{}
	Status() *string
	Tags() BedrockDataAutomationProjectTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutCustomOutputConfiguration(value *BedrockDataAutomationProjectCustomOutputConfiguration)
	PutOverrideConfiguration(value *BedrockDataAutomationProjectOverrideConfiguration)
	PutStandardOutputConfiguration(value *BedrockDataAutomationProjectStandardOutputConfiguration)
	PutTags(value interface{})
	ResetCustomOutputConfiguration()
	ResetKmsEncryptionContext()
	ResetKmsKeyId()
	ResetOverrideConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProjectDescription()
	ResetStandardOutputConfiguration()
	ResetTags()
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

// The jsii proxy struct for BedrockDataAutomationProject
type jsiiProxy_BedrockDataAutomationProject struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BedrockDataAutomationProject) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) CustomOutputConfiguration() BedrockDataAutomationProjectCustomOutputConfigurationOutputReference {
	var returns BedrockDataAutomationProjectCustomOutputConfigurationOutputReference
	_jsii_.Get(
		j,
		"customOutputConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) CustomOutputConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customOutputConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) KmsEncryptionContext() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"kmsEncryptionContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) KmsEncryptionContextInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"kmsEncryptionContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) OverrideConfiguration() BedrockDataAutomationProjectOverrideConfigurationOutputReference {
	var returns BedrockDataAutomationProjectOverrideConfigurationOutputReference
	_jsii_.Get(
		j,
		"overrideConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) OverrideConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) ProjectArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) ProjectDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) ProjectDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) ProjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) ProjectStage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) StandardOutputConfiguration() BedrockDataAutomationProjectStandardOutputConfigurationOutputReference {
	var returns BedrockDataAutomationProjectStandardOutputConfigurationOutputReference
	_jsii_.Get(
		j,
		"standardOutputConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) StandardOutputConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"standardOutputConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) Tags() BedrockDataAutomationProjectTagsList {
	var returns BedrockDataAutomationProjectTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProject) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_data_automation_project awscc_bedrock_data_automation_project} Resource.
func NewBedrockDataAutomationProject(scope constructs.Construct, id *string, config *BedrockDataAutomationProjectConfig) BedrockDataAutomationProject {
	_init_.Initialize()

	if err := validateNewBedrockDataAutomationProjectParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockDataAutomationProject{}

	_jsii_.Create(
		"awscc.bedrockDataAutomationProject.BedrockDataAutomationProject",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_data_automation_project awscc_bedrock_data_automation_project} Resource.
func NewBedrockDataAutomationProject_Override(b BedrockDataAutomationProject, scope constructs.Construct, id *string, config *BedrockDataAutomationProjectConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.bedrockDataAutomationProject.BedrockDataAutomationProject",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProject)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProject)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProject)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProject)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProject)SetKmsEncryptionContext(val *map[string]*string) {
	if err := j.validateSetKmsEncryptionContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsEncryptionContext",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProject)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProject)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProject)SetProjectDescription(val *string) {
	if err := j.validateSetProjectDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectDescription",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProject)SetProjectName(val *string) {
	if err := j.validateSetProjectNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectName",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProject)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProject)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a BedrockDataAutomationProject resource upon running "cdktf plan <stack-name>".
func BedrockDataAutomationProject_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBedrockDataAutomationProject_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.bedrockDataAutomationProject.BedrockDataAutomationProject",
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
func BedrockDataAutomationProject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockDataAutomationProject_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.bedrockDataAutomationProject.BedrockDataAutomationProject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockDataAutomationProject_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockDataAutomationProject_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.bedrockDataAutomationProject.BedrockDataAutomationProject",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockDataAutomationProject_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockDataAutomationProject_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.bedrockDataAutomationProject.BedrockDataAutomationProject",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BedrockDataAutomationProject_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.bedrockDataAutomationProject.BedrockDataAutomationProject",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProject) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProject) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProject) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProject) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProject) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProject) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProject) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProject) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProject) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProject) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProject) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProject) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) PutCustomOutputConfiguration(value *BedrockDataAutomationProjectCustomOutputConfiguration) {
	if err := b.validatePutCustomOutputConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCustomOutputConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) PutOverrideConfiguration(value *BedrockDataAutomationProjectOverrideConfiguration) {
	if err := b.validatePutOverrideConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putOverrideConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) PutStandardOutputConfiguration(value *BedrockDataAutomationProjectStandardOutputConfiguration) {
	if err := b.validatePutStandardOutputConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStandardOutputConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) PutTags(value interface{}) {
	if err := b.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTags",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) ResetCustomOutputConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetCustomOutputConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) ResetKmsEncryptionContext() {
	_jsii_.InvokeVoid(
		b,
		"resetKmsEncryptionContext",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		b,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) ResetOverrideConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) ResetProjectDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetProjectDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) ResetStandardOutputConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetStandardOutputConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProject) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProject) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProject) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProject) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProject) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

