package datazoneenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/datazoneenvironment/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_environment awscc_datazone_environment}.
type DatazoneEnvironment interface {
	cdktf.TerraformResource
	AwsAccountId() *string
	AwsAccountRegion() *string
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
	CreatedAt() *string
	CreatedBy() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DomainId() *string
	DomainIdentifier() *string
	SetDomainIdentifier(val *string)
	DomainIdentifierInput() *string
	EnvironmentAccountIdentifier() *string
	SetEnvironmentAccountIdentifier(val *string)
	EnvironmentAccountIdentifierInput() *string
	EnvironmentAccountRegion() *string
	SetEnvironmentAccountRegion(val *string)
	EnvironmentAccountRegionInput() *string
	EnvironmentBlueprintId() *string
	EnvironmentId() *string
	EnvironmentProfileId() *string
	EnvironmentProfileIdentifier() *string
	SetEnvironmentProfileIdentifier(val *string)
	EnvironmentProfileIdentifierInput() *string
	EnvironmentRoleArn() *string
	SetEnvironmentRoleArn(val *string)
	EnvironmentRoleArnInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GlossaryTerms() *[]*string
	SetGlossaryTerms(val *[]*string)
	GlossaryTermsInput() *[]*string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ProjectId() *string
	ProjectIdentifier() *string
	SetProjectIdentifier(val *string)
	ProjectIdentifierInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProviderName() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatedAt() *string
	UserParameters() DatazoneEnvironmentUserParametersList
	UserParametersInput() interface{}
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
	PutUserParameters(value interface{})
	ResetDescription()
	ResetEnvironmentAccountIdentifier()
	ResetEnvironmentAccountRegion()
	ResetEnvironmentProfileIdentifier()
	ResetEnvironmentRoleArn()
	ResetGlossaryTerms()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetUserParameters()
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

// The jsii proxy struct for DatazoneEnvironment
type jsiiProxy_DatazoneEnvironment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatazoneEnvironment) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) AwsAccountRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) DomainIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) DomainIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) EnvironmentAccountIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentAccountIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) EnvironmentAccountIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentAccountIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) EnvironmentAccountRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentAccountRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) EnvironmentAccountRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentAccountRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) EnvironmentBlueprintId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentBlueprintId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) EnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) EnvironmentProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) EnvironmentProfileIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentProfileIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) EnvironmentProfileIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentProfileIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) EnvironmentRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) EnvironmentRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) GlossaryTerms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"glossaryTerms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) GlossaryTermsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"glossaryTermsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) ProjectIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) ProjectIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) UserParameters() DatazoneEnvironmentUserParametersList {
	var returns DatazoneEnvironmentUserParametersList
	_jsii_.Get(
		j,
		"userParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironment) UserParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userParametersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_environment awscc_datazone_environment} Resource.
func NewDatazoneEnvironment(scope constructs.Construct, id *string, config *DatazoneEnvironmentConfig) DatazoneEnvironment {
	_init_.Initialize()

	if err := validateNewDatazoneEnvironmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatazoneEnvironment{}

	_jsii_.Create(
		"awscc.datazoneEnvironment.DatazoneEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_environment awscc_datazone_environment} Resource.
func NewDatazoneEnvironment_Override(d DatazoneEnvironment, scope constructs.Construct, id *string, config *DatazoneEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.datazoneEnvironment.DatazoneEnvironment",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatazoneEnvironment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironment)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironment)SetDomainIdentifier(val *string) {
	if err := j.validateSetDomainIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainIdentifier",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironment)SetEnvironmentAccountIdentifier(val *string) {
	if err := j.validateSetEnvironmentAccountIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentAccountIdentifier",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironment)SetEnvironmentAccountRegion(val *string) {
	if err := j.validateSetEnvironmentAccountRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentAccountRegion",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironment)SetEnvironmentProfileIdentifier(val *string) {
	if err := j.validateSetEnvironmentProfileIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentProfileIdentifier",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironment)SetEnvironmentRoleArn(val *string) {
	if err := j.validateSetEnvironmentRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentRoleArn",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironment)SetGlossaryTerms(val *[]*string) {
	if err := j.validateSetGlossaryTermsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"glossaryTerms",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironment)SetProjectIdentifier(val *string) {
	if err := j.validateSetProjectIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectIdentifier",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a DatazoneEnvironment resource upon running "cdktf plan <stack-name>".
func DatazoneEnvironment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatazoneEnvironment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.datazoneEnvironment.DatazoneEnvironment",
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
func DatazoneEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatazoneEnvironment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.datazoneEnvironment.DatazoneEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatazoneEnvironment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatazoneEnvironment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.datazoneEnvironment.DatazoneEnvironment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatazoneEnvironment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatazoneEnvironment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.datazoneEnvironment.DatazoneEnvironment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatazoneEnvironment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.datazoneEnvironment.DatazoneEnvironment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatazoneEnvironment) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatazoneEnvironment) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatazoneEnvironment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatazoneEnvironment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatazoneEnvironment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatazoneEnvironment) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatazoneEnvironment) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatazoneEnvironment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatazoneEnvironment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatazoneEnvironment) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatazoneEnvironment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatazoneEnvironment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneEnvironment) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatazoneEnvironment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatazoneEnvironment) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatazoneEnvironment) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatazoneEnvironment) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatazoneEnvironment) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatazoneEnvironment) PutUserParameters(value interface{}) {
	if err := d.validatePutUserParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUserParameters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneEnvironment) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneEnvironment) ResetEnvironmentAccountIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetEnvironmentAccountIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneEnvironment) ResetEnvironmentAccountRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetEnvironmentAccountRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneEnvironment) ResetEnvironmentProfileIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetEnvironmentProfileIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneEnvironment) ResetEnvironmentRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetEnvironmentRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneEnvironment) ResetGlossaryTerms() {
	_jsii_.InvokeVoid(
		d,
		"resetGlossaryTerms",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneEnvironment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneEnvironment) ResetUserParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetUserParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneEnvironment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneEnvironment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneEnvironment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneEnvironment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneEnvironment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneEnvironment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

