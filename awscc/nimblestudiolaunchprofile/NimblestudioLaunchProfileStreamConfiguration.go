package nimblestudiolaunchprofile


type NimblestudioLaunchProfileStreamConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#clipboard_mode NimblestudioLaunchProfile#clipboard_mode}.
	ClipboardMode *string `field:"required" json:"clipboardMode" yaml:"clipboardMode"`
	// <p>The EC2 instance types that users can select from when launching a streaming session             with this launch profile.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#ec_2_instance_types NimblestudioLaunchProfile#ec_2_instance_types}
	Ec2InstanceTypes *[]*string `field:"required" json:"ec2InstanceTypes" yaml:"ec2InstanceTypes"`
	// <p>The streaming images that users can select from when launching a streaming session             with this launch profile.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#streaming_image_ids NimblestudioLaunchProfile#streaming_image_ids}
	StreamingImageIds *[]*string `field:"required" json:"streamingImageIds" yaml:"streamingImageIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#automatic_termination_mode NimblestudioLaunchProfile#automatic_termination_mode}.
	AutomaticTerminationMode *string `field:"optional" json:"automaticTerminationMode" yaml:"automaticTerminationMode"`
	// <p>The length of time, in minutes, that a streaming session can be active before it is             stopped or terminated.
	//
	// After this point, Nimble Studio automatically terminates or
	//             stops the session. The default length of time is 690 minutes, and the maximum length of
	//             time is 30 days.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#max_session_length_in_minutes NimblestudioLaunchProfile#max_session_length_in_minutes}
	MaxSessionLengthInMinutes *float64 `field:"optional" json:"maxSessionLengthInMinutes" yaml:"maxSessionLengthInMinutes"`
	// <p>Integer that determines if you can start and stop your sessions and how long a session             can stay in the <code>STOPPED</code> state.
	//
	// The default value is 0. The maximum value is
	//             5760.</p>
	//          <p>This field is allowed only when <code>sessionPersistenceMode</code> is
	//                 <code>ACTIVATED</code> and <code>automaticTerminationMode</code> is
	//                 <code>ACTIVATED</code>.</p>
	//          <p>If the value is set to 0, your sessions can?t be <code>STOPPED</code>. If you then
	//             call <code>StopStreamingSession</code>, the session fails. If the time that a session
	//             stays in the <code>READY</code> state exceeds the <code>maxSessionLengthInMinutes</code>
	//             value, the session will automatically be terminated (instead of
	//             <code>STOPPED</code>).</p>
	//          <p>If the value is set to a positive number, the session can be stopped. You can call
	//                 <code>StopStreamingSession</code> to stop sessions in the <code>READY</code> state.
	//             If the time that a session stays in the <code>READY</code> state exceeds the
	//                 <code>maxSessionLengthInMinutes</code> value, the session will automatically be
	//             stopped (instead of terminated).</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#max_stopped_session_length_in_minutes NimblestudioLaunchProfile#max_stopped_session_length_in_minutes}
	MaxStoppedSessionLengthInMinutes *float64 `field:"optional" json:"maxStoppedSessionLengthInMinutes" yaml:"maxStoppedSessionLengthInMinutes"`
	// <p>Configures how streaming sessions are backed up when launched from this launch             profile.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#session_backup NimblestudioLaunchProfile#session_backup}
	SessionBackup *NimblestudioLaunchProfileStreamConfigurationSessionBackup `field:"optional" json:"sessionBackup" yaml:"sessionBackup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#session_persistence_mode NimblestudioLaunchProfile#session_persistence_mode}.
	SessionPersistenceMode *string `field:"optional" json:"sessionPersistenceMode" yaml:"sessionPersistenceMode"`
	// <p>The configuration for a streaming session?s upload storage.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#session_storage NimblestudioLaunchProfile#session_storage}
	SessionStorage *NimblestudioLaunchProfileStreamConfigurationSessionStorage `field:"optional" json:"sessionStorage" yaml:"sessionStorage"`
	// <p>Custom volume configuration for the root volumes that are attached to streaming             sessions.</p>          <p>This parameter is only allowed when <code>sessionPersistenceMode</code> is                 <code>ACTIVATED</code>.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#volume_configuration NimblestudioLaunchProfile#volume_configuration}
	VolumeConfiguration *NimblestudioLaunchProfileStreamConfigurationVolumeConfiguration `field:"optional" json:"volumeConfiguration" yaml:"volumeConfiguration"`
}

