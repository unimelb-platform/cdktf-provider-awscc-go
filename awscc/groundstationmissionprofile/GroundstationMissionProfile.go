package groundstationmissionprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/groundstationmissionprofile/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/groundstation_mission_profile awscc_groundstation_mission_profile}.
type GroundstationMissionProfile interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContactPostPassDurationSeconds() *float64
	SetContactPostPassDurationSeconds(val *float64)
	ContactPostPassDurationSecondsInput() *float64
	ContactPrePassDurationSeconds() *float64
	SetContactPrePassDurationSeconds(val *float64)
	ContactPrePassDurationSecondsInput() *float64
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DataflowEdges() GroundstationMissionProfileDataflowEdgesList
	DataflowEdgesInput() interface{}
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
	MinimumViableContactDurationSeconds() *float64
	SetMinimumViableContactDurationSeconds(val *float64)
	MinimumViableContactDurationSecondsInput() *float64
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
	Region() *string
	StreamsKmsKey() GroundstationMissionProfileStreamsKmsKeyOutputReference
	StreamsKmsKeyInput() interface{}
	StreamsKmsRole() *string
	SetStreamsKmsRole(val *string)
	StreamsKmsRoleInput() *string
	Tags() GroundstationMissionProfileTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TrackingConfigArn() *string
	SetTrackingConfigArn(val *string)
	TrackingConfigArnInput() *string
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
	PutDataflowEdges(value interface{})
	PutStreamsKmsKey(value *GroundstationMissionProfileStreamsKmsKey)
	PutTags(value interface{})
	ResetContactPostPassDurationSeconds()
	ResetContactPrePassDurationSeconds()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStreamsKmsKey()
	ResetStreamsKmsRole()
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

// The jsii proxy struct for GroundstationMissionProfile
type jsiiProxy_GroundstationMissionProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GroundstationMissionProfile) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) ContactPostPassDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"contactPostPassDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) ContactPostPassDurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"contactPostPassDurationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) ContactPrePassDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"contactPrePassDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) ContactPrePassDurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"contactPrePassDurationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) DataflowEdges() GroundstationMissionProfileDataflowEdgesList {
	var returns GroundstationMissionProfileDataflowEdgesList
	_jsii_.Get(
		j,
		"dataflowEdges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) DataflowEdgesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataflowEdgesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) MinimumViableContactDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumViableContactDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) MinimumViableContactDurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumViableContactDurationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) StreamsKmsKey() GroundstationMissionProfileStreamsKmsKeyOutputReference {
	var returns GroundstationMissionProfileStreamsKmsKeyOutputReference
	_jsii_.Get(
		j,
		"streamsKmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) StreamsKmsKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamsKmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) StreamsKmsRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamsKmsRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) StreamsKmsRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamsKmsRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) Tags() GroundstationMissionProfileTagsList {
	var returns GroundstationMissionProfileTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) TrackingConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackingConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroundstationMissionProfile) TrackingConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackingConfigArnInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/groundstation_mission_profile awscc_groundstation_mission_profile} Resource.
func NewGroundstationMissionProfile(scope constructs.Construct, id *string, config *GroundstationMissionProfileConfig) GroundstationMissionProfile {
	_init_.Initialize()

	if err := validateNewGroundstationMissionProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GroundstationMissionProfile{}

	_jsii_.Create(
		"awscc.groundstationMissionProfile.GroundstationMissionProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/groundstation_mission_profile awscc_groundstation_mission_profile} Resource.
func NewGroundstationMissionProfile_Override(g GroundstationMissionProfile, scope constructs.Construct, id *string, config *GroundstationMissionProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.groundstationMissionProfile.GroundstationMissionProfile",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GroundstationMissionProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GroundstationMissionProfile)SetContactPostPassDurationSeconds(val *float64) {
	if err := j.validateSetContactPostPassDurationSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactPostPassDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_GroundstationMissionProfile)SetContactPrePassDurationSeconds(val *float64) {
	if err := j.validateSetContactPrePassDurationSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactPrePassDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_GroundstationMissionProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GroundstationMissionProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GroundstationMissionProfile)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GroundstationMissionProfile)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GroundstationMissionProfile)SetMinimumViableContactDurationSeconds(val *float64) {
	if err := j.validateSetMinimumViableContactDurationSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumViableContactDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_GroundstationMissionProfile)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GroundstationMissionProfile)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GroundstationMissionProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GroundstationMissionProfile)SetStreamsKmsRole(val *string) {
	if err := j.validateSetStreamsKmsRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamsKmsRole",
		val,
	)
}

func (j *jsiiProxy_GroundstationMissionProfile)SetTrackingConfigArn(val *string) {
	if err := j.validateSetTrackingConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackingConfigArn",
		val,
	)
}

// Generates CDKTF code for importing a GroundstationMissionProfile resource upon running "cdktf plan <stack-name>".
func GroundstationMissionProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGroundstationMissionProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.groundstationMissionProfile.GroundstationMissionProfile",
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
func GroundstationMissionProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroundstationMissionProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.groundstationMissionProfile.GroundstationMissionProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GroundstationMissionProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroundstationMissionProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.groundstationMissionProfile.GroundstationMissionProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GroundstationMissionProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroundstationMissionProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.groundstationMissionProfile.GroundstationMissionProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GroundstationMissionProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.groundstationMissionProfile.GroundstationMissionProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GroundstationMissionProfile) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GroundstationMissionProfile) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GroundstationMissionProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationMissionProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationMissionProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationMissionProfile) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationMissionProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationMissionProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationMissionProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationMissionProfile) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationMissionProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationMissionProfile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationMissionProfile) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GroundstationMissionProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationMissionProfile) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GroundstationMissionProfile) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GroundstationMissionProfile) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GroundstationMissionProfile) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GroundstationMissionProfile) PutDataflowEdges(value interface{}) {
	if err := g.validatePutDataflowEdgesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataflowEdges",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationMissionProfile) PutStreamsKmsKey(value *GroundstationMissionProfileStreamsKmsKey) {
	if err := g.validatePutStreamsKmsKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStreamsKmsKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationMissionProfile) PutTags(value interface{}) {
	if err := g.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTags",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroundstationMissionProfile) ResetContactPostPassDurationSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetContactPostPassDurationSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationMissionProfile) ResetContactPrePassDurationSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetContactPrePassDurationSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationMissionProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationMissionProfile) ResetStreamsKmsKey() {
	_jsii_.InvokeVoid(
		g,
		"resetStreamsKmsKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationMissionProfile) ResetStreamsKmsRole() {
	_jsii_.InvokeVoid(
		g,
		"resetStreamsKmsRole",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationMissionProfile) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroundstationMissionProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationMissionProfile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationMissionProfile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationMissionProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationMissionProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroundstationMissionProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

