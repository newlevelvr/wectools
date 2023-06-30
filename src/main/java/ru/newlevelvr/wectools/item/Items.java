package ru.newlevelvr.wectools.item;

import net.fabricmc.fabric.api.client.itemgroup.FabricItemGroupBuilder;
import net.fabricmc.fabric.api.item.v1.FabricItemSettings;
import net.minecraft.item.Item;
import net.minecraft.item.ItemGroup;
import net.minecraft.item.ItemStack;
import net.minecraft.util.Identifier;
import net.minecraft.util.registry.Registry;
import ru.newlevelvr.wectools.WorldEditColoredTools;

import java.util.List;

public class Items {

    private static final List<String> COLORS = List.of(
            "gray", "lightgray", "white", "candle", "brown", "copper", "orange", "yellow", "lime",
            "green", "cyan", "lightblue", "blue", "black", "purple", "violet", "pink", "red"
    );
    private static final List<String> TOOLS = List.of(
            "wooden_pickaxe", "wooden_axe", "wooden_shovel", "wooden_hoe", "stick"
    );

    private static Item wand;

    public static Item getWand() {
        return wand;
    }

    public static final ItemGroup ITEM_GROUP = FabricItemGroupBuilder
            .build(new Identifier(WorldEditColoredTools.MOD_ID, "tools"), () -> new ItemStack(getWand()));

    private static Item registerTool(String id) {
        var itemId = new Identifier(WorldEditColoredTools.MOD_ID, id);
        return Registry.register(Registry.ITEM, itemId, new Item(new FabricItemSettings().maxCount(1).group(ITEM_GROUP)));
    }

    private static void registerToolGroup(String prefix) {
        for (String color : COLORS) {
            registerTool("%s_%s".formatted(prefix, color));
        }
    }

    public static void registerTools() {
        TOOLS.forEach(Items::registerToolGroup);
        wand = registerTool("wand");
    }
}
