package curreportdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/curreportdefinition/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cur_report_definition awscc_cur_report_definition}.
type CurReportDefinition interface {
	cdktf.TerraformResource
	AdditionalArtifacts() *[]*string
	SetAdditionalArtifacts(val *[]*string)
	AdditionalArtifactsInput() *[]*string
	AdditionalSchemaElements() *[]*string
	SetAdditionalSchemaElements(val *[]*string)
	AdditionalSchemaElementsInput() *[]*string
	BillingViewArn() *string
	SetBillingViewArn(val *string)
	BillingViewArnInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Compression() *string
	SetCompression(val *string)
	CompressionInput() *string
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
	RefreshClosedReports() interface{}
	SetRefreshClosedReports(val interface{})
	RefreshClosedReportsInput() interface{}
	ReportName() *string
	SetReportName(val *string)
	ReportNameInput() *string
	ReportVersioning() *string
	SetReportVersioning(val *string)
	ReportVersioningInput() *string
	S3Bucket() *string
	SetS3Bucket(val *string)
	S3BucketInput() *string
	S3Prefix() *string
	SetS3Prefix(val *string)
	S3PrefixInput() *string
	S3Region() *string
	SetS3Region(val *string)
	S3RegionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeUnit() *string
	SetTimeUnit(val *string)
	TimeUnitInput() *string
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
	ResetAdditionalArtifacts()
	ResetAdditionalSchemaElements()
	ResetBillingViewArn()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for CurReportDefinition
type jsiiProxy_CurReportDefinition struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CurReportDefinition) AdditionalArtifacts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) AdditionalArtifactsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) AdditionalSchemaElements() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalSchemaElements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) AdditionalSchemaElementsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalSchemaElementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) BillingViewArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingViewArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) BillingViewArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingViewArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) RefreshClosedReports() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"refreshClosedReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) RefreshClosedReportsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"refreshClosedReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) ReportName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) ReportNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) ReportVersioning() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportVersioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) ReportVersioningInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportVersioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) S3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) S3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) S3Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) S3PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3PrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) S3Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) S3RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3RegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) TimeUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurReportDefinition) TimeUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeUnitInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cur_report_definition awscc_cur_report_definition} Resource.
func NewCurReportDefinition(scope constructs.Construct, id *string, config *CurReportDefinitionConfig) CurReportDefinition {
	_init_.Initialize()

	if err := validateNewCurReportDefinitionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CurReportDefinition{}

	_jsii_.Create(
		"awscc.curReportDefinition.CurReportDefinition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cur_report_definition awscc_cur_report_definition} Resource.
func NewCurReportDefinition_Override(c CurReportDefinition, scope constructs.Construct, id *string, config *CurReportDefinitionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.curReportDefinition.CurReportDefinition",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetAdditionalArtifacts(val *[]*string) {
	if err := j.validateSetAdditionalArtifactsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalArtifacts",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetAdditionalSchemaElements(val *[]*string) {
	if err := j.validateSetAdditionalSchemaElementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalSchemaElements",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetBillingViewArn(val *string) {
	if err := j.validateSetBillingViewArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingViewArn",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetRefreshClosedReports(val interface{}) {
	if err := j.validateSetRefreshClosedReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshClosedReports",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetReportName(val *string) {
	if err := j.validateSetReportNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportName",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetReportVersioning(val *string) {
	if err := j.validateSetReportVersioningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportVersioning",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetS3Bucket(val *string) {
	if err := j.validateSetS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Bucket",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetS3Prefix(val *string) {
	if err := j.validateSetS3PrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Prefix",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetS3Region(val *string) {
	if err := j.validateSetS3RegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Region",
		val,
	)
}

func (j *jsiiProxy_CurReportDefinition)SetTimeUnit(val *string) {
	if err := j.validateSetTimeUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeUnit",
		val,
	)
}

// Generates CDKTF code for importing a CurReportDefinition resource upon running "cdktf plan <stack-name>".
func CurReportDefinition_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCurReportDefinition_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.curReportDefinition.CurReportDefinition",
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
func CurReportDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCurReportDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.curReportDefinition.CurReportDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CurReportDefinition_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCurReportDefinition_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.curReportDefinition.CurReportDefinition",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CurReportDefinition_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCurReportDefinition_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.curReportDefinition.CurReportDefinition",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CurReportDefinition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.curReportDefinition.CurReportDefinition",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CurReportDefinition) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CurReportDefinition) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CurReportDefinition) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CurReportDefinition) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CurReportDefinition) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CurReportDefinition) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CurReportDefinition) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CurReportDefinition) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CurReportDefinition) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CurReportDefinition) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CurReportDefinition) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CurReportDefinition) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurReportDefinition) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CurReportDefinition) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CurReportDefinition) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CurReportDefinition) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CurReportDefinition) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CurReportDefinition) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CurReportDefinition) ResetAdditionalArtifacts() {
	_jsii_.InvokeVoid(
		c,
		"resetAdditionalArtifacts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurReportDefinition) ResetAdditionalSchemaElements() {
	_jsii_.InvokeVoid(
		c,
		"resetAdditionalSchemaElements",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurReportDefinition) ResetBillingViewArn() {
	_jsii_.InvokeVoid(
		c,
		"resetBillingViewArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurReportDefinition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurReportDefinition) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurReportDefinition) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurReportDefinition) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurReportDefinition) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurReportDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurReportDefinition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

