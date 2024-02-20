package secretsmanagersecret


type SecretsmanagerSecretGenerateSecretString struct {
	// A string that excludes characters in the generated password.
	//
	// By default, all characters from the included sets can be used. The string can be a minimum length of 0 characters and a maximum length of 7168 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#exclude_characters SecretsmanagerSecret#exclude_characters}
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// Specifies the generated password should not include lowercase letters.
	//
	// By default, ecrets Manager disables this parameter, and the generated password can include lowercase False, and the generated password can include lowercase letters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#exclude_lowercase SecretsmanagerSecret#exclude_lowercase}
	ExcludeLowercase interface{} `field:"optional" json:"excludeLowercase" yaml:"excludeLowercase"`
	// Specifies that the generated password should exclude digits.
	//
	// By default, Secrets Manager does not enable the parameter, False, and the generated password can include digits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#exclude_numbers SecretsmanagerSecret#exclude_numbers}
	ExcludeNumbers interface{} `field:"optional" json:"excludeNumbers" yaml:"excludeNumbers"`
	// Specifies that the generated password should not include punctuation characters.
	//
	// The default if you do not include this switch parameter is that punctuation characters can be included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#exclude_punctuation SecretsmanagerSecret#exclude_punctuation}
	ExcludePunctuation interface{} `field:"optional" json:"excludePunctuation" yaml:"excludePunctuation"`
	// Specifies that the generated password should not include uppercase letters.
	//
	// The default behavior is False, and the generated password can include uppercase letters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#exclude_uppercase SecretsmanagerSecret#exclude_uppercase}
	ExcludeUppercase interface{} `field:"optional" json:"excludeUppercase" yaml:"excludeUppercase"`
	// The JSON key name used to add the generated password to the JSON structure specified by the SecretStringTemplate parameter.
	//
	// If you specify this parameter, then you must also specify SecretStringTemplate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#generate_string_key SecretsmanagerSecret#generate_string_key}
	GenerateStringKey *string `field:"optional" json:"generateStringKey" yaml:"generateStringKey"`
	// Specifies that the generated password can include the space character.
	//
	// By default, Secrets Manager disables this parameter, and the generated password doesn't include space
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#include_space SecretsmanagerSecret#include_space}
	IncludeSpace interface{} `field:"optional" json:"includeSpace" yaml:"includeSpace"`
	// The desired length of the generated password.
	//
	// The default value if you do not include this parameter is 32 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#password_length SecretsmanagerSecret#password_length}
	PasswordLength *float64 `field:"optional" json:"passwordLength" yaml:"passwordLength"`
	// Specifies whether the generated password must include at least one of every allowed character type.
	//
	// By default, Secrets Manager enables this parameter, and the generated password includes at least one of every character type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#require_each_included_type SecretsmanagerSecret#require_each_included_type}
	RequireEachIncludedType interface{} `field:"optional" json:"requireEachIncludedType" yaml:"requireEachIncludedType"`
	// A properly structured JSON string that the generated password can be added to.
	//
	// If you specify this parameter, then you must also specify GenerateStringKey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#secret_string_template SecretsmanagerSecret#secret_string_template}
	SecretStringTemplate *string `field:"optional" json:"secretStringTemplate" yaml:"secretStringTemplate"`
}

