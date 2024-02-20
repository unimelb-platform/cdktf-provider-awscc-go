package lightsaildatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lightsaildatabase/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_database awscc_lightsail_database}.
type LightsailDatabase interface {
	cdktf.TerraformResource
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	BackupRetention() interface{}
	SetBackupRetention(val interface{})
	BackupRetentionInput() interface{}
	CaCertificateIdentifier() *string
	SetCaCertificateIdentifier(val *string)
	CaCertificateIdentifierInput() *string
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
	DatabaseArn() *string
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
	MasterDatabaseName() *string
	SetMasterDatabaseName(val *string)
	MasterDatabaseNameInput() *string
	MasterUsername() *string
	SetMasterUsername(val *string)
	MasterUsernameInput() *string
	MasterUserPassword() *string
	SetMasterUserPassword(val *string)
	MasterUserPasswordInput() *string
	// The tree node.
	Node() constructs.Node
	PreferredBackupWindow() *string
	SetPreferredBackupWindow(val *string)
	PreferredBackupWindowInput() *string
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	PreferredMaintenanceWindowInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	PubliclyAccessibleInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	RelationalDatabaseBlueprintId() *string
	SetRelationalDatabaseBlueprintId(val *string)
	RelationalDatabaseBlueprintIdInput() *string
	RelationalDatabaseBundleId() *string
	SetRelationalDatabaseBundleId(val *string)
	RelationalDatabaseBundleIdInput() *string
	RelationalDatabaseName() *string
	SetRelationalDatabaseName(val *string)
	RelationalDatabaseNameInput() *string
	RelationalDatabaseParameters() LightsailDatabaseRelationalDatabaseParametersList
	RelationalDatabaseParametersInput() interface{}
	RotateMasterUserPassword() interface{}
	SetRotateMasterUserPassword(val interface{})
	RotateMasterUserPasswordInput() interface{}
	Tags() LightsailDatabaseTagsList
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
	PutRelationalDatabaseParameters(value interface{})
	PutTags(value interface{})
	ResetAvailabilityZone()
	ResetBackupRetention()
	ResetCaCertificateIdentifier()
	ResetMasterUserPassword()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreferredBackupWindow()
	ResetPreferredMaintenanceWindow()
	ResetPubliclyAccessible()
	ResetRelationalDatabaseParameters()
	ResetRotateMasterUserPassword()
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

// The jsii proxy struct for LightsailDatabase
type jsiiProxy_LightsailDatabase struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LightsailDatabase) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) BackupRetention() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) BackupRetentionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) CaCertificateIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) CaCertificateIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) DatabaseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) MasterDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) MasterDatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterDatabaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) MasterUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) MasterUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) MasterUserPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) PreferredBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) PreferredBackupWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) PreferredMaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) PubliclyAccessibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) RelationalDatabaseBlueprintId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relationalDatabaseBlueprintId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) RelationalDatabaseBlueprintIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relationalDatabaseBlueprintIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) RelationalDatabaseBundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relationalDatabaseBundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) RelationalDatabaseBundleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relationalDatabaseBundleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) RelationalDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relationalDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) RelationalDatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relationalDatabaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) RelationalDatabaseParameters() LightsailDatabaseRelationalDatabaseParametersList {
	var returns LightsailDatabaseRelationalDatabaseParametersList
	_jsii_.Get(
		j,
		"relationalDatabaseParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) RelationalDatabaseParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relationalDatabaseParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) RotateMasterUserPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rotateMasterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) RotateMasterUserPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rotateMasterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) Tags() LightsailDatabaseTagsList {
	var returns LightsailDatabaseTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabase) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_database awscc_lightsail_database} Resource.
func NewLightsailDatabase(scope constructs.Construct, id *string, config *LightsailDatabaseConfig) LightsailDatabase {
	_init_.Initialize()

	if err := validateNewLightsailDatabaseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LightsailDatabase{}

	_jsii_.Create(
		"awscc.lightsailDatabase.LightsailDatabase",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_database awscc_lightsail_database} Resource.
func NewLightsailDatabase_Override(l LightsailDatabase, scope constructs.Construct, id *string, config *LightsailDatabaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lightsailDatabase.LightsailDatabase",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetBackupRetention(val interface{}) {
	if err := j.validateSetBackupRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetention",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetCaCertificateIdentifier(val *string) {
	if err := j.validateSetCaCertificateIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificateIdentifier",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetMasterDatabaseName(val *string) {
	if err := j.validateSetMasterDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterDatabaseName",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetMasterUsername(val *string) {
	if err := j.validateSetMasterUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetMasterUserPassword(val *string) {
	if err := j.validateSetMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetPreferredBackupWindow(val *string) {
	if err := j.validateSetPreferredBackupWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredBackupWindow",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetPreferredMaintenanceWindow(val *string) {
	if err := j.validateSetPreferredMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetRelationalDatabaseBlueprintId(val *string) {
	if err := j.validateSetRelationalDatabaseBlueprintIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"relationalDatabaseBlueprintId",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetRelationalDatabaseBundleId(val *string) {
	if err := j.validateSetRelationalDatabaseBundleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"relationalDatabaseBundleId",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetRelationalDatabaseName(val *string) {
	if err := j.validateSetRelationalDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"relationalDatabaseName",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabase)SetRotateMasterUserPassword(val interface{}) {
	if err := j.validateSetRotateMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotateMasterUserPassword",
		val,
	)
}

// Generates CDKTF code for importing a LightsailDatabase resource upon running "cdktf plan <stack-name>".
func LightsailDatabase_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLightsailDatabase_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.lightsailDatabase.LightsailDatabase",
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
func LightsailDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLightsailDatabase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lightsailDatabase.LightsailDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LightsailDatabase_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLightsailDatabase_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lightsailDatabase.LightsailDatabase",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LightsailDatabase_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLightsailDatabase_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.lightsailDatabase.LightsailDatabase",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LightsailDatabase_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.lightsailDatabase.LightsailDatabase",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LightsailDatabase) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LightsailDatabase) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LightsailDatabase) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabase) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabase) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabase) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabase) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabase) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabase) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabase) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabase) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabase) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LightsailDatabase) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabase) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LightsailDatabase) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LightsailDatabase) PutRelationalDatabaseParameters(value interface{}) {
	if err := l.validatePutRelationalDatabaseParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRelationalDatabaseParameters",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LightsailDatabase) PutTags(value interface{}) {
	if err := l.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTags",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LightsailDatabase) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		l,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabase) ResetBackupRetention() {
	_jsii_.InvokeVoid(
		l,
		"resetBackupRetention",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabase) ResetCaCertificateIdentifier() {
	_jsii_.InvokeVoid(
		l,
		"resetCaCertificateIdentifier",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabase) ResetMasterUserPassword() {
	_jsii_.InvokeVoid(
		l,
		"resetMasterUserPassword",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabase) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabase) ResetPreferredBackupWindow() {
	_jsii_.InvokeVoid(
		l,
		"resetPreferredBackupWindow",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabase) ResetPreferredMaintenanceWindow() {
	_jsii_.InvokeVoid(
		l,
		"resetPreferredMaintenanceWindow",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabase) ResetPubliclyAccessible() {
	_jsii_.InvokeVoid(
		l,
		"resetPubliclyAccessible",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabase) ResetRelationalDatabaseParameters() {
	_jsii_.InvokeVoid(
		l,
		"resetRelationalDatabaseParameters",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabase) ResetRotateMasterUserPassword() {
	_jsii_.InvokeVoid(
		l,
		"resetRotateMasterUserPassword",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabase) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabase) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabase) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabase) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

