package com.engineersbox.example.common.register;

import com.engineersbox.example.ExampleMod;
import net.minecraft.block.Block;
import net.minecraft.block.material.Material;
import net.minecraftforge.fml.RegistryObject;
import net.minecraftforge.registries.DeferredRegister;
import net.minecraftforge.registries.ForgeRegistries;

public class Blocks {

    public static final DeferredRegister<Block> BLOCKS = DeferredRegister.create(ForgeRegistries.BLOCKS, ExampleMod.MODID);

    public static final RegistryObject<Block> COPPER_BLOCK = BLOCKS.register("copper_block", () -> new Block(Block.Properties.create(Material.IRON)));
    public static final RegistryObject<Block> TITANIUM_BLOCK = BLOCKS.register("titanium_block", () -> new Block(Block.Properties.create(Material.IRON)));
	public static final RegistryObject<Block> SOME_TESTER_BLOCK = BLOCKS.register("some_tester_block", () -> new Block(Block.Properties.create(Material.IRON)));
	public static final RegistryObject<Block> RANDOM_TEST_BLOCK = BLOCKS.register("random_test_block", () -> new Block(Block.Properties.create(Material.GLASS)));
	public static final RegistryObject<Block> ANOTHER_TEST_BLOCK = BLOCKS.register("another_test_block", () -> new Block(Block.Properties.create(Material.GLASS)));
}
