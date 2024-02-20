package cloudwatchcompositealarm

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"awscc.cloudwatchCompositeAlarm.CloudwatchCompositeAlarm",
		reflect.TypeOf((*CloudwatchCompositeAlarm)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionsEnabled", GoGetter: "ActionsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "actionsEnabledInput", GoGetter: "ActionsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "actionsSuppressor", GoGetter: "ActionsSuppressor"},
			_jsii_.MemberProperty{JsiiProperty: "actionsSuppressorExtensionPeriod", GoGetter: "ActionsSuppressorExtensionPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "actionsSuppressorExtensionPeriodInput", GoGetter: "ActionsSuppressorExtensionPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "actionsSuppressorInput", GoGetter: "ActionsSuppressorInput"},
			_jsii_.MemberProperty{JsiiProperty: "actionsSuppressorWaitPeriod", GoGetter: "ActionsSuppressorWaitPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "actionsSuppressorWaitPeriodInput", GoGetter: "ActionsSuppressorWaitPeriodInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alarmActions", GoGetter: "AlarmActions"},
			_jsii_.MemberProperty{JsiiProperty: "alarmActionsInput", GoGetter: "AlarmActionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "alarmDescription", GoGetter: "AlarmDescription"},
			_jsii_.MemberProperty{JsiiProperty: "alarmDescriptionInput", GoGetter: "AlarmDescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "alarmName", GoGetter: "AlarmName"},
			_jsii_.MemberProperty{JsiiProperty: "alarmNameInput", GoGetter: "AlarmNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "alarmRule", GoGetter: "AlarmRule"},
			_jsii_.MemberProperty{JsiiProperty: "alarmRuleInput", GoGetter: "AlarmRuleInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "insufficientDataActions", GoGetter: "InsufficientDataActions"},
			_jsii_.MemberProperty{JsiiProperty: "insufficientDataActionsInput", GoGetter: "InsufficientDataActionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "okActions", GoGetter: "OkActions"},
			_jsii_.MemberProperty{JsiiProperty: "okActionsInput", GoGetter: "OkActionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetActionsEnabled", GoMethod: "ResetActionsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetActionsSuppressor", GoMethod: "ResetActionsSuppressor"},
			_jsii_.MemberMethod{JsiiMethod: "resetActionsSuppressorExtensionPeriod", GoMethod: "ResetActionsSuppressorExtensionPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetActionsSuppressorWaitPeriod", GoMethod: "ResetActionsSuppressorWaitPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlarmActions", GoMethod: "ResetAlarmActions"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlarmDescription", GoMethod: "ResetAlarmDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlarmName", GoMethod: "ResetAlarmName"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsufficientDataActions", GoMethod: "ResetInsufficientDataActions"},
			_jsii_.MemberMethod{JsiiMethod: "resetOkActions", GoMethod: "ResetOkActions"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CloudwatchCompositeAlarm{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscc.cloudwatchCompositeAlarm.CloudwatchCompositeAlarmConfig",
		reflect.TypeOf((*CloudwatchCompositeAlarmConfig)(nil)).Elem(),
	)
}
