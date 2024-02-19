package cleanroomsanalysistemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/cleanroomsanalysistemplate/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_analysis_template awscc_cleanrooms_analysis_template}.
type CleanroomsAnalysisTemplate interface {
	cdktf.TerraformResource
	AnalysisParameters() CleanroomsAnalysisTemplateAnalysisParametersList
	AnalysisParametersInput() interface{}
	AnalysisTemplateIdentifier() *string
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CollaborationArn() *string
	CollaborationIdentifier() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Format() *string
	SetFormat(val *string)
	FormatInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MembershipArn() *string
	MembershipIdentifier() *string
	SetMembershipIdentifier(val *string)
	MembershipIdentifierInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	Schema() CleanroomsAnalysisTemplateSchemaOutputReference
	Source() CleanroomsAnalysisTemplateSourceOutputReference
	SourceInput() interface{}
	Tags() CleanroomsAnalysisTemplateTagsList
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
	PutAnalysisParameters(value interface{})
	PutSource(value *CleanroomsAnalysisTemplateSource)
	PutTags(value interface{})
	ResetAnalysisParameters()
	ResetDescription()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for CleanroomsAnalysisTemplate
type jsiiProxy_CleanroomsAnalysisTemplate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) AnalysisParameters() CleanroomsAnalysisTemplateAnalysisParametersList {
	var returns CleanroomsAnalysisTemplateAnalysisParametersList
	_jsii_.Get(
		j,
		"analysisParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) AnalysisParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"analysisParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) AnalysisTemplateIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysisTemplateIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) CollaborationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaborationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) CollaborationIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaborationIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) MembershipArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"membershipArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) MembershipIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"membershipIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) MembershipIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"membershipIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) Schema() CleanroomsAnalysisTemplateSchemaOutputReference {
	var returns CleanroomsAnalysisTemplateSchemaOutputReference
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) Source() CleanroomsAnalysisTemplateSourceOutputReference {
	var returns CleanroomsAnalysisTemplateSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) SourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) Tags() CleanroomsAnalysisTemplateTagsList {
	var returns CleanroomsAnalysisTemplateTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_analysis_template awscc_cleanrooms_analysis_template} Resource.
func NewCleanroomsAnalysisTemplate(scope constructs.Construct, id *string, config *CleanroomsAnalysisTemplateConfig) CleanroomsAnalysisTemplate {
	_init_.Initialize()

	if err := validateNewCleanroomsAnalysisTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanroomsAnalysisTemplate{}

	_jsii_.Create(
		"awscc.cleanroomsAnalysisTemplate.CleanroomsAnalysisTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_analysis_template awscc_cleanrooms_analysis_template} Resource.
func NewCleanroomsAnalysisTemplate_Override(c CleanroomsAnalysisTemplate, scope constructs.Construct, id *string, config *CleanroomsAnalysisTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.cleanroomsAnalysisTemplate.CleanroomsAnalysisTemplate",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate)SetMembershipIdentifier(val *string) {
	if err := j.validateSetMembershipIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"membershipIdentifier",
		val,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CleanroomsAnalysisTemplate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a CleanroomsAnalysisTemplate resource upon running "cdktf plan <stack-name>".
func CleanroomsAnalysisTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCleanroomsAnalysisTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.cleanroomsAnalysisTemplate.CleanroomsAnalysisTemplate",
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
func CleanroomsAnalysisTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCleanroomsAnalysisTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cleanroomsAnalysisTemplate.CleanroomsAnalysisTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CleanroomsAnalysisTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCleanroomsAnalysisTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cleanroomsAnalysisTemplate.CleanroomsAnalysisTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CleanroomsAnalysisTemplate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCleanroomsAnalysisTemplate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.cleanroomsAnalysisTemplate.CleanroomsAnalysisTemplate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CleanroomsAnalysisTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.cleanroomsAnalysisTemplate.CleanroomsAnalysisTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplate) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplate) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CleanroomsAnalysisTemplate) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) PutAnalysisParameters(value interface{}) {
	if err := c.validatePutAnalysisParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAnalysisParameters",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) PutSource(value *CleanroomsAnalysisTemplateSource) {
	if err := c.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSource",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) PutTags(value interface{}) {
	if err := c.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTags",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) ResetAnalysisParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetAnalysisParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsAnalysisTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

