package quicksighttheme


type QuicksightThemeConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_theme#data_color_palette QuicksightTheme#data_color_palette}.
	DataColorPalette *QuicksightThemeConfigurationDataColorPalette `field:"optional" json:"dataColorPalette" yaml:"dataColorPalette"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_theme#sheet QuicksightTheme#sheet}.
	Sheet *QuicksightThemeConfigurationSheet `field:"optional" json:"sheet" yaml:"sheet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_theme#typography QuicksightTheme#typography}.
	Typography *QuicksightThemeConfigurationTypography `field:"optional" json:"typography" yaml:"typography"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_theme#ui_color_palette QuicksightTheme#ui_color_palette}.
	UiColorPalette *QuicksightThemeConfigurationUiColorPalette `field:"optional" json:"uiColorPalette" yaml:"uiColorPalette"`
}

