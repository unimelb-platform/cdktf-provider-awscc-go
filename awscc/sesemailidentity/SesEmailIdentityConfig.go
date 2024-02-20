package sesemailidentity

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SesEmailIdentityConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The email address or domain to verify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_email_identity#email_identity SesEmailIdentity#email_identity}
	EmailIdentity *string `field:"required" json:"emailIdentity" yaml:"emailIdentity"`
	// Used to associate a configuration set with an email identity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_email_identity#configuration_set_attributes SesEmailIdentity#configuration_set_attributes}
	ConfigurationSetAttributes *SesEmailIdentityConfigurationSetAttributes `field:"optional" json:"configurationSetAttributes" yaml:"configurationSetAttributes"`
	// Used to enable or disable DKIM authentication for an email identity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_email_identity#dkim_attributes SesEmailIdentity#dkim_attributes}
	DkimAttributes *SesEmailIdentityDkimAttributes `field:"optional" json:"dkimAttributes" yaml:"dkimAttributes"`
	// If your request includes this object, Amazon SES configures the identity to use Bring Your Own DKIM (BYODKIM) for DKIM authentication purposes, or, configures the key length to be used for Easy DKIM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_email_identity#dkim_signing_attributes SesEmailIdentity#dkim_signing_attributes}
	DkimSigningAttributes *SesEmailIdentityDkimSigningAttributes `field:"optional" json:"dkimSigningAttributes" yaml:"dkimSigningAttributes"`
	// Used to enable or disable feedback forwarding for an identity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_email_identity#feedback_attributes SesEmailIdentity#feedback_attributes}
	FeedbackAttributes *SesEmailIdentityFeedbackAttributes `field:"optional" json:"feedbackAttributes" yaml:"feedbackAttributes"`
	// Used to enable or disable the custom Mail-From domain configuration for an email identity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_email_identity#mail_from_attributes SesEmailIdentity#mail_from_attributes}
	MailFromAttributes *SesEmailIdentityMailFromAttributes `field:"optional" json:"mailFromAttributes" yaml:"mailFromAttributes"`
}

