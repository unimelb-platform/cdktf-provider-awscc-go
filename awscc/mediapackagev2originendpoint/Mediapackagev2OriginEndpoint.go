package mediapackagev2originendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/mediapackagev2originendpoint/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint awscc_mediapackagev2_origin_endpoint}.
type Mediapackagev2OriginEndpoint interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChannelGroupName() *string
	SetChannelGroupName(val *string)
	ChannelGroupNameInput() *string
	ChannelName() *string
	SetChannelName(val *string)
	ChannelNameInput() *string
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
	CreatedAt() *string
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
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HlsManifests() Mediapackagev2OriginEndpointHlsManifestsList
	HlsManifestsInput() interface{}
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LowLatencyHlsManifests() Mediapackagev2OriginEndpointLowLatencyHlsManifestsList
	LowLatencyHlsManifestsInput() interface{}
	ModifiedAt() *string
	// The tree node.
	Node() constructs.Node
	OriginEndpointName() *string
	SetOriginEndpointName(val *string)
	OriginEndpointNameInput() *string
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
	Segment() Mediapackagev2OriginEndpointSegmentOutputReference
	SegmentInput() interface{}
	StartoverWindowSeconds() *float64
	SetStartoverWindowSeconds(val *float64)
	StartoverWindowSecondsInput() *float64
	Tags() Mediapackagev2OriginEndpointTagsList
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
	PutHlsManifests(value interface{})
	PutLowLatencyHlsManifests(value interface{})
	PutSegment(value *Mediapackagev2OriginEndpointSegment)
	PutTags(value interface{})
	ResetChannelGroupName()
	ResetChannelName()
	ResetDescription()
	ResetHlsManifests()
	ResetLowLatencyHlsManifests()
	ResetOriginEndpointName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSegment()
	ResetStartoverWindowSeconds()
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

// The jsii proxy struct for Mediapackagev2OriginEndpoint
type jsiiProxy_Mediapackagev2OriginEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) ChannelGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) ChannelGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) ChannelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) ChannelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) ContainerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) ContainerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) HlsManifests() Mediapackagev2OriginEndpointHlsManifestsList {
	var returns Mediapackagev2OriginEndpointHlsManifestsList
	_jsii_.Get(
		j,
		"hlsManifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) HlsManifestsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsManifestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) LowLatencyHlsManifests() Mediapackagev2OriginEndpointLowLatencyHlsManifestsList {
	var returns Mediapackagev2OriginEndpointLowLatencyHlsManifestsList
	_jsii_.Get(
		j,
		"lowLatencyHlsManifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) LowLatencyHlsManifestsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lowLatencyHlsManifestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) ModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) OriginEndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originEndpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) OriginEndpointNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originEndpointNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) Segment() Mediapackagev2OriginEndpointSegmentOutputReference {
	var returns Mediapackagev2OriginEndpointSegmentOutputReference
	_jsii_.Get(
		j,
		"segment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) SegmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"segmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) StartoverWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startoverWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) StartoverWindowSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startoverWindowSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) Tags() Mediapackagev2OriginEndpointTagsList {
	var returns Mediapackagev2OriginEndpointTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint awscc_mediapackagev2_origin_endpoint} Resource.
func NewMediapackagev2OriginEndpoint(scope constructs.Construct, id *string, config *Mediapackagev2OriginEndpointConfig) Mediapackagev2OriginEndpoint {
	_init_.Initialize()

	if err := validateNewMediapackagev2OriginEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Mediapackagev2OriginEndpoint{}

	_jsii_.Create(
		"awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint awscc_mediapackagev2_origin_endpoint} Resource.
func NewMediapackagev2OriginEndpoint_Override(m Mediapackagev2OriginEndpoint, scope constructs.Construct, id *string, config *Mediapackagev2OriginEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpoint",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint)SetChannelGroupName(val *string) {
	if err := j.validateSetChannelGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelGroupName",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint)SetChannelName(val *string) {
	if err := j.validateSetChannelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelName",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint)SetContainerType(val *string) {
	if err := j.validateSetContainerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerType",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint)SetOriginEndpointName(val *string) {
	if err := j.validateSetOriginEndpointNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originEndpointName",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Mediapackagev2OriginEndpoint)SetStartoverWindowSeconds(val *float64) {
	if err := j.validateSetStartoverWindowSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startoverWindowSeconds",
		val,
	)
}

// Generates CDKTF code for importing a Mediapackagev2OriginEndpoint resource upon running "cdktf plan <stack-name>".
func Mediapackagev2OriginEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMediapackagev2OriginEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpoint",
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
func Mediapackagev2OriginEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediapackagev2OriginEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Mediapackagev2OriginEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediapackagev2OriginEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Mediapackagev2OriginEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediapackagev2OriginEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Mediapackagev2OriginEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.mediapackagev2OriginEndpoint.Mediapackagev2OriginEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) PutHlsManifests(value interface{}) {
	if err := m.validatePutHlsManifestsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putHlsManifests",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) PutLowLatencyHlsManifests(value interface{}) {
	if err := m.validatePutLowLatencyHlsManifestsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLowLatencyHlsManifests",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) PutSegment(value *Mediapackagev2OriginEndpointSegment) {
	if err := m.validatePutSegmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSegment",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) PutTags(value interface{}) {
	if err := m.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTags",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) ResetChannelGroupName() {
	_jsii_.InvokeVoid(
		m,
		"resetChannelGroupName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) ResetChannelName() {
	_jsii_.InvokeVoid(
		m,
		"resetChannelName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) ResetHlsManifests() {
	_jsii_.InvokeVoid(
		m,
		"resetHlsManifests",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) ResetLowLatencyHlsManifests() {
	_jsii_.InvokeVoid(
		m,
		"resetLowLatencyHlsManifests",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) ResetOriginEndpointName() {
	_jsii_.InvokeVoid(
		m,
		"resetOriginEndpointName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) ResetSegment() {
	_jsii_.InvokeVoid(
		m,
		"resetSegment",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) ResetStartoverWindowSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetStartoverWindowSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mediapackagev2OriginEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

