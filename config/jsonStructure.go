package config

type RegistryFiles struct {
	Block      string `json:"block"`
	Item       string `json:"item"`
	TileEntity string `json:"tile_entity"`
	Fluid      string `json:"fluid"`
}

type ForgeCLIConfig struct {
	ModName      string        `json:"mod_name"`
	ResourcesDir string        `json:"resources_dir"`
	RegFiles     RegistryFiles `json:"registry_files"`
}
