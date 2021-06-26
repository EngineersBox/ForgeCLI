package template

import (
	"engineersbox/forgecli/registration/common"
	"io/ioutil"
	"strings"
)

type Template string

const (
	SingleModelBlockState          Template = "templates/blockstate/single_model_block_template.fct"
	FullBlockMultiTextureModel     Template = "templates/model/block/full_block_multi_texture.fct"
	FullBlockSingleTextureModel    Template = "templates/model/block/full_block_single_texture.fct"
	BlockItemParent                Template = "templates/model/item/block_item_template.fct"
	SingleLayerItem                Template = "templates/model/item/single_layer_item_template.fct"
	RegistryBasicBlockWithMaterial Template = "templates/registry/block/basic_static_block_with_material.fct"
	RegistryBasicItemWIthGroup     Template = "templates/registry/item/basic_item_with_group.fct"
)

func ImportTemplate(filename Template) string {
	b, err := ioutil.ReadFile(string(filename))
	common.CheckError(err)
	return string(b)
}

func ReplaceInlineFormatters(template string, formatters map[string]string) string {
	for key, replacement := range formatters {
		template = strings.ReplaceAll(template, key, replacement)
	}
	return template
}
