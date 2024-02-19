package fsxdatarepositoryassociation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/fsxdatarepositoryassociation/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fsx_data_repository_association awscc_fsx_data_repository_association}.
type FsxDataRepositoryAssociation interface {
	cdktf.TerraformResource
	AssociationId() *string
	BatchImportMetaDataOnCreate() interface{}
	SetBatchImportMetaDataOnCreate(val interface{})
	BatchImportMetaDataOnCreateInput() interface{}
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
	DataRepositoryPath() *string
	SetDataRepositoryPath(val *string)
	DataRepositoryPathInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FileSystemId() *string
	SetFileSystemId(val *string)
	FileSystemIdInput() *string
	FileSystemPath() *string
	SetFileSystemPath(val *string)
	FileSystemPathInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	ImportedFileChunkSize() *float64
	SetImportedFileChunkSize(val *float64)
	ImportedFileChunkSizeInput() *float64
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
	ResourceArn() *string
	S3() FsxDataRepositoryAssociationS3OutputReference
	S3Input() interface{}
	Tags() FsxDataRepositoryAssociationTagsList
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutS3(value *FsxDataRepositoryAssociationS3)
	PutTags(value interface{})
	ResetBatchImportMetaDataOnCreate()
	ResetImportedFileChunkSize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetS3()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for FsxDataRepositoryAssociation
type jsiiProxy_FsxDataRepositoryAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) AssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) BatchImportMetaDataOnCreate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchImportMetaDataOnCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) BatchImportMetaDataOnCreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchImportMetaDataOnCreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) DataRepositoryPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataRepositoryPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) DataRepositoryPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataRepositoryPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) FileSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) FileSystemPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) FileSystemPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) ImportedFileChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importedFileChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) ImportedFileChunkSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importedFileChunkSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) S3() FsxDataRepositoryAssociationS3OutputReference {
	var returns FsxDataRepositoryAssociationS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) S3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Tags() FsxDataRepositoryAssociationTagsList {
	var returns FsxDataRepositoryAssociationTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fsx_data_repository_association awscc_fsx_data_repository_association} Resource.
func NewFsxDataRepositoryAssociation(scope constructs.Construct, id *string, config *FsxDataRepositoryAssociationConfig) FsxDataRepositoryAssociation {
	_init_.Initialize()

	if err := validateNewFsxDataRepositoryAssociationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxDataRepositoryAssociation{}

	_jsii_.Create(
		"awscc.fsxDataRepositoryAssociation.FsxDataRepositoryAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fsx_data_repository_association awscc_fsx_data_repository_association} Resource.
func NewFsxDataRepositoryAssociation_Override(f FsxDataRepositoryAssociation, scope constructs.Construct, id *string, config *FsxDataRepositoryAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.fsxDataRepositoryAssociation.FsxDataRepositoryAssociation",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation)SetBatchImportMetaDataOnCreate(val interface{}) {
	if err := j.validateSetBatchImportMetaDataOnCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchImportMetaDataOnCreate",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation)SetDataRepositoryPath(val *string) {
	if err := j.validateSetDataRepositoryPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataRepositoryPath",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation)SetFileSystemId(val *string) {
	if err := j.validateSetFileSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemId",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation)SetFileSystemPath(val *string) {
	if err := j.validateSetFileSystemPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemPath",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation)SetImportedFileChunkSize(val *float64) {
	if err := j.validateSetImportedFileChunkSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importedFileChunkSize",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a FsxDataRepositoryAssociation resource upon running "cdktf plan <stack-name>".
func FsxDataRepositoryAssociation_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateFsxDataRepositoryAssociation_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.fsxDataRepositoryAssociation.FsxDataRepositoryAssociation",
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
func FsxDataRepositoryAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFsxDataRepositoryAssociation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.fsxDataRepositoryAssociation.FsxDataRepositoryAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FsxDataRepositoryAssociation_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFsxDataRepositoryAssociation_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.fsxDataRepositoryAssociation.FsxDataRepositoryAssociation",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FsxDataRepositoryAssociation_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFsxDataRepositoryAssociation_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.fsxDataRepositoryAssociation.FsxDataRepositoryAssociation",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FsxDataRepositoryAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.fsxDataRepositoryAssociation.FsxDataRepositoryAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) PutS3(value *FsxDataRepositoryAssociationS3) {
	if err := f.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putS3",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) PutTags(value interface{}) {
	if err := f.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTags",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) ResetBatchImportMetaDataOnCreate() {
	_jsii_.InvokeVoid(
		f,
		"resetBatchImportMetaDataOnCreate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) ResetImportedFileChunkSize() {
	_jsii_.InvokeVoid(
		f,
		"resetImportedFileChunkSize",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) ResetS3() {
	_jsii_.InvokeVoid(
		f,
		"resetS3",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

