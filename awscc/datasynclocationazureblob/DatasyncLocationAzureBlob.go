package datasynclocationazureblob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/datasynclocationazureblob/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_azure_blob awscc_datasync_location_azure_blob}.
type DatasyncLocationAzureBlob interface {
	cdktf.TerraformResource
	AgentArns() *[]*string
	SetAgentArns(val *[]*string)
	AgentArnsInput() *[]*string
	AzureAccessTier() *string
	SetAzureAccessTier(val *string)
	AzureAccessTierInput() *string
	AzureBlobAuthenticationType() *string
	SetAzureBlobAuthenticationType(val *string)
	AzureBlobAuthenticationTypeInput() *string
	AzureBlobContainerUrl() *string
	SetAzureBlobContainerUrl(val *string)
	AzureBlobContainerUrlInput() *string
	AzureBlobSasConfiguration() DatasyncLocationAzureBlobAzureBlobSasConfigurationOutputReference
	AzureBlobSasConfigurationInput() interface{}
	AzureBlobType() *string
	SetAzureBlobType(val *string)
	AzureBlobTypeInput() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocationArn() *string
	LocationUri() *string
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
	Subdirectory() *string
	SetSubdirectory(val *string)
	SubdirectoryInput() *string
	Tags() DatasyncLocationAzureBlobTagsList
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
	PutAzureBlobSasConfiguration(value *DatasyncLocationAzureBlobAzureBlobSasConfiguration)
	PutTags(value interface{})
	ResetAzureAccessTier()
	ResetAzureBlobAuthenticationType()
	ResetAzureBlobContainerUrl()
	ResetAzureBlobSasConfiguration()
	ResetAzureBlobType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSubdirectory()
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

// The jsii proxy struct for DatasyncLocationAzureBlob
type jsiiProxy_DatasyncLocationAzureBlob struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) AgentArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"agentArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) AgentArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"agentArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) AzureAccessTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureAccessTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) AzureAccessTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureAccessTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) AzureBlobAuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureBlobAuthenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) AzureBlobAuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureBlobAuthenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) AzureBlobContainerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureBlobContainerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) AzureBlobContainerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureBlobContainerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) AzureBlobSasConfiguration() DatasyncLocationAzureBlobAzureBlobSasConfigurationOutputReference {
	var returns DatasyncLocationAzureBlobAzureBlobSasConfigurationOutputReference
	_jsii_.Get(
		j,
		"azureBlobSasConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) AzureBlobSasConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureBlobSasConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) AzureBlobType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureBlobType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) AzureBlobTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureBlobTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) LocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) LocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) Subdirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) SubdirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) Tags() DatasyncLocationAzureBlobTagsList {
	var returns DatasyncLocationAzureBlobTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationAzureBlob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_azure_blob awscc_datasync_location_azure_blob} Resource.
func NewDatasyncLocationAzureBlob(scope constructs.Construct, id *string, config *DatasyncLocationAzureBlobConfig) DatasyncLocationAzureBlob {
	_init_.Initialize()

	if err := validateNewDatasyncLocationAzureBlobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatasyncLocationAzureBlob{}

	_jsii_.Create(
		"awscc.datasyncLocationAzureBlob.DatasyncLocationAzureBlob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_azure_blob awscc_datasync_location_azure_blob} Resource.
func NewDatasyncLocationAzureBlob_Override(d DatasyncLocationAzureBlob, scope constructs.Construct, id *string, config *DatasyncLocationAzureBlobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.datasyncLocationAzureBlob.DatasyncLocationAzureBlob",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatasyncLocationAzureBlob)SetAgentArns(val *[]*string) {
	if err := j.validateSetAgentArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentArns",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationAzureBlob)SetAzureAccessTier(val *string) {
	if err := j.validateSetAzureAccessTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureAccessTier",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationAzureBlob)SetAzureBlobAuthenticationType(val *string) {
	if err := j.validateSetAzureBlobAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureBlobAuthenticationType",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationAzureBlob)SetAzureBlobContainerUrl(val *string) {
	if err := j.validateSetAzureBlobContainerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureBlobContainerUrl",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationAzureBlob)SetAzureBlobType(val *string) {
	if err := j.validateSetAzureBlobTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureBlobType",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationAzureBlob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationAzureBlob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationAzureBlob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationAzureBlob)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationAzureBlob)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationAzureBlob)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationAzureBlob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationAzureBlob)SetSubdirectory(val *string) {
	if err := j.validateSetSubdirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subdirectory",
		val,
	)
}

// Generates CDKTF code for importing a DatasyncLocationAzureBlob resource upon running "cdktf plan <stack-name>".
func DatasyncLocationAzureBlob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatasyncLocationAzureBlob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.datasyncLocationAzureBlob.DatasyncLocationAzureBlob",
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
func DatasyncLocationAzureBlob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatasyncLocationAzureBlob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.datasyncLocationAzureBlob.DatasyncLocationAzureBlob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatasyncLocationAzureBlob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatasyncLocationAzureBlob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.datasyncLocationAzureBlob.DatasyncLocationAzureBlob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatasyncLocationAzureBlob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatasyncLocationAzureBlob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.datasyncLocationAzureBlob.DatasyncLocationAzureBlob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatasyncLocationAzureBlob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.datasyncLocationAzureBlob.DatasyncLocationAzureBlob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatasyncLocationAzureBlob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatasyncLocationAzureBlob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatasyncLocationAzureBlob) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatasyncLocationAzureBlob) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatasyncLocationAzureBlob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatasyncLocationAzureBlob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatasyncLocationAzureBlob) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatasyncLocationAzureBlob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatasyncLocationAzureBlob) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatasyncLocationAzureBlob) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) PutAzureBlobSasConfiguration(value *DatasyncLocationAzureBlobAzureBlobSasConfiguration) {
	if err := d.validatePutAzureBlobSasConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureBlobSasConfiguration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) PutTags(value interface{}) {
	if err := d.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) ResetAzureAccessTier() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureAccessTier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) ResetAzureBlobAuthenticationType() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureBlobAuthenticationType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) ResetAzureBlobContainerUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureBlobContainerUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) ResetAzureBlobSasConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureBlobSasConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) ResetAzureBlobType() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureBlobType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) ResetSubdirectory() {
	_jsii_.InvokeVoid(
		d,
		"resetSubdirectory",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationAzureBlob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

