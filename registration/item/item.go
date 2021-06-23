package item

import "io/fs"

const (
	ItemModel = `
{
  "parent": "item/generated",
  "textures": {
    "layer0": "<0>:item/<1>"
  }
}
	`
	FileModeOct fs.FileMode = 0644
)
