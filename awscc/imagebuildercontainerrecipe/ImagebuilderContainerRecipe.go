package imagebuildercontainerrecipe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/imagebuildercontainerrecipe/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_container_recipe awscc_imagebuilder_container_recipe}.
type ImagebuilderContainerRecipe interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Components() ImagebuilderContainerRecipeComponentsList
	ComponentsInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerType() *string
	SetContainerType(val *string)
	ContainerTypeInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DockerfileTemplateData() *string
	SetDockerfileTemplateData(val *string)
	DockerfileTemplateDataInput() *string
	DockerfileTemplateUri() *string
	SetDockerfileTemplateUri(val *string)
	DockerfileTemplateUriInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	ImageOsVersionOverride() *string
	SetImageOsVersionOverride(val *string)
	ImageOsVersionOverrideInput() *string
	InstanceConfiguration() ImagebuilderContainerRecipeInstanceConfigurationOutputReference
	InstanceConfigurationInput() interface{}
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ParentImage() *string
	SetParentImage(val *string)
	ParentImageInput() *string
	PlatformOverride() *string
	SetPlatformOverride(val *string)
	PlatformOverrideInput() *string
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TargetRepository() ImagebuilderContainerRecipeTargetRepositoryOutputReference
	TargetRepositoryInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	WorkingDirectory() *string
	SetWorkingDirectory(val *string)
	WorkingDirectoryInput() *string
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
	PutComponents(value interface{})
	PutInstanceConfiguration(value *ImagebuilderContainerRecipeInstanceConfiguration)
	PutTargetRepository(value *ImagebuilderContainerRecipeTargetRepository)
	ResetComponents()
	ResetContainerType()
	ResetDescription()
	ResetDockerfileTemplateData()
	ResetDockerfileTemplateUri()
	ResetImageOsVersionOverride()
	ResetInstanceConfiguration()
	ResetKmsKeyId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentImage()
	ResetPlatformOverride()
	ResetTags()
	ResetTargetRepository()
	ResetVersion()
	ResetWorkingDirectory()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ImagebuilderContainerRecipe
type jsiiProxy_ImagebuilderContainerRecipe struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) Components() ImagebuilderContainerRecipeComponentsList {
	var returns ImagebuilderContainerRecipeComponentsList
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) ComponentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"componentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) ContainerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) ContainerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) DockerfileTemplateData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfileTemplateData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) DockerfileTemplateDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfileTemplateDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) DockerfileTemplateUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfileTemplateUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) DockerfileTemplateUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfileTemplateUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) ImageOsVersionOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageOsVersionOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) ImageOsVersionOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageOsVersionOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) InstanceConfiguration() ImagebuilderContainerRecipeInstanceConfigurationOutputReference {
	var returns ImagebuilderContainerRecipeInstanceConfigurationOutputReference
	_jsii_.Get(
		j,
		"instanceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) InstanceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) ParentImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) ParentImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) PlatformOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) PlatformOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) TargetRepository() ImagebuilderContainerRecipeTargetRepositoryOutputReference {
	var returns ImagebuilderContainerRecipeTargetRepositoryOutputReference
	_jsii_.Get(
		j,
		"targetRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) TargetRepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) WorkingDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipe) WorkingDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectoryInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_container_recipe awscc_imagebuilder_container_recipe} Resource.
func NewImagebuilderContainerRecipe(scope constructs.Construct, id *string, config *ImagebuilderContainerRecipeConfig) ImagebuilderContainerRecipe {
	_init_.Initialize()

	if err := validateNewImagebuilderContainerRecipeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImagebuilderContainerRecipe{}

	_jsii_.Create(
		"awscc.imagebuilderContainerRecipe.ImagebuilderContainerRecipe",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_container_recipe awscc_imagebuilder_container_recipe} Resource.
func NewImagebuilderContainerRecipe_Override(i ImagebuilderContainerRecipe, scope constructs.Construct, id *string, config *ImagebuilderContainerRecipeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.imagebuilderContainerRecipe.ImagebuilderContainerRecipe",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetContainerType(val *string) {
	if err := j.validateSetContainerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerType",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetDockerfileTemplateData(val *string) {
	if err := j.validateSetDockerfileTemplateDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerfileTemplateData",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetDockerfileTemplateUri(val *string) {
	if err := j.validateSetDockerfileTemplateUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerfileTemplateUri",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetImageOsVersionOverride(val *string) {
	if err := j.validateSetImageOsVersionOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageOsVersionOverride",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetParentImage(val *string) {
	if err := j.validateSetParentImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentImage",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetPlatformOverride(val *string) {
	if err := j.validateSetPlatformOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformOverride",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipe)SetWorkingDirectory(val *string) {
	if err := j.validateSetWorkingDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDirectory",
		val,
	)
}

// Generates CDKTF code for importing a ImagebuilderContainerRecipe resource upon running "cdktf plan <stack-name>".
func ImagebuilderContainerRecipe_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateImagebuilderContainerRecipe_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.imagebuilderContainerRecipe.ImagebuilderContainerRecipe",
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
func ImagebuilderContainerRecipe_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateImagebuilderContainerRecipe_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.imagebuilderContainerRecipe.ImagebuilderContainerRecipe",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ImagebuilderContainerRecipe_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateImagebuilderContainerRecipe_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.imagebuilderContainerRecipe.ImagebuilderContainerRecipe",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ImagebuilderContainerRecipe_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateImagebuilderContainerRecipe_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.imagebuilderContainerRecipe.ImagebuilderContainerRecipe",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ImagebuilderContainerRecipe_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.imagebuilderContainerRecipe.ImagebuilderContainerRecipe",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) PutComponents(value interface{}) {
	if err := i.validatePutComponentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putComponents",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) PutInstanceConfiguration(value *ImagebuilderContainerRecipeInstanceConfiguration) {
	if err := i.validatePutInstanceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putInstanceConfiguration",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) PutTargetRepository(value *ImagebuilderContainerRecipeTargetRepository) {
	if err := i.validatePutTargetRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTargetRepository",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ResetComponents() {
	_jsii_.InvokeVoid(
		i,
		"resetComponents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ResetContainerType() {
	_jsii_.InvokeVoid(
		i,
		"resetContainerType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ResetDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ResetDockerfileTemplateData() {
	_jsii_.InvokeVoid(
		i,
		"resetDockerfileTemplateData",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ResetDockerfileTemplateUri() {
	_jsii_.InvokeVoid(
		i,
		"resetDockerfileTemplateUri",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ResetImageOsVersionOverride() {
	_jsii_.InvokeVoid(
		i,
		"resetImageOsVersionOverride",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ResetInstanceConfiguration() {
	_jsii_.InvokeVoid(
		i,
		"resetInstanceConfiguration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		i,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ResetName() {
	_jsii_.InvokeVoid(
		i,
		"resetName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ResetParentImage() {
	_jsii_.InvokeVoid(
		i,
		"resetParentImage",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ResetPlatformOverride() {
	_jsii_.InvokeVoid(
		i,
		"resetPlatformOverride",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ResetTargetRepository() {
	_jsii_.InvokeVoid(
		i,
		"resetTargetRepository",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ResetVersion() {
	_jsii_.InvokeVoid(
		i,
		"resetVersion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ResetWorkingDirectory() {
	_jsii_.InvokeVoid(
		i,
		"resetWorkingDirectory",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderContainerRecipe) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

