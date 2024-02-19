package quicksightdatasource


type QuicksightDataSourcePermissions struct {
	// <p>The IAM action to grant or revoke permissions on.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#actions QuicksightDataSource#actions}
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// <p>The Amazon Resource Name (ARN) of the principal.
	//
	// This can be one of the
	//             following:</p>
	//         <ul>
	//             <li>
	//                 <p>The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)</p>
	//             </li>
	//             <li>
	//                 <p>The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)</p>
	//             </li>
	//             <li>
	//                 <p>The ARN of an AWS account root: This is an IAM ARN rather than a QuickSight
	//                     ARN. Use this option only to share resources (templates) across AWS accounts.
	//                     (This is less common.) </p>
	//             </li>
	//          </ul>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#principal QuicksightDataSource#principal}
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

