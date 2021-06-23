package block

const (
	BlockState = `
{
  "variants": {
    "": {
      "model": "<0>:block/<1>"
    }
  }
}
	`
	BlockModel = `
{
  "parent": "block/cube_all",
  "textures": {
    "all": "<0>:block/<1>"
  }
}
	`
	BlockItemModel = `
{
  "parent": "<0>:block/<1>"
}
	`
	RegistryObject = `public static final RegistryObject<Block> <0> = BLOCKS.register("<1>", () -> new Block(Block.Properties.create(Material.<2>)));`
)
