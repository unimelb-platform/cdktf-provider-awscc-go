package quicksighttheme


type QuicksightThemeConfigurationSheet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_theme#tile QuicksightTheme#tile}.
	Tile *QuicksightThemeConfigurationSheetTile `field:"optional" json:"tile" yaml:"tile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_theme#tile_layout QuicksightTheme#tile_layout}.
	TileLayout *QuicksightThemeConfigurationSheetTileLayout `field:"optional" json:"tileLayout" yaml:"tileLayout"`
}

