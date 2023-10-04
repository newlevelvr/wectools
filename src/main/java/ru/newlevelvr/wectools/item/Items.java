package ru.newlevelvr.wectools.item;

import net.fabricmc.fabric.api.item.v1.FabricItemSettings;
import net.fabricmc.fabric.api.itemgroup.v1.FabricItemGroup;
import net.minecraft.item.Item;
import net.minecraft.item.ItemGroup;
import net.minecraft.item.ItemStack;
import net.minecraft.registry.Registries;
import net.minecraft.registry.Registry;
import net.minecraft.text.Text;
import net.minecraft.util.Identifier;
import ru.newlevelvr.wectools.WorldEditColoredTools;

import java.util.ArrayList;
import java.util.List;

public class Items {

    private static final List<String> COLORS = List.of(
            "gray", "lightgray", "white", "candle", "brown", "copper", "orange", "yellow", "lime",
            "green", "cyan", "lightblue", "blue", "black", "purple", "violet", "pink", "red"
    );
    private static final List<String> TOOLS = List.of(
            "wooden_pickaxe", "wooden_axe", "wooden_shovel", "wooden_hoe", "stick"
    );

    private static final List<Item> ITEMS = new ArrayList<>();

    private static Item wand;

    public static Item getWand() {
        return wand;
    }

    public static final ItemGroup ITEM_GROUP = FabricItemGroup.builder()
            .displayName(Text.translatable("itemGroup.wectools.tools"))
            .icon(() -> new ItemStack(getWand()))
            .entries((context, entries) -> ITEMS.forEach(entries::add))
            .build();

    private static Item registerTool(String id) {
        var itemId = new Identifier(WorldEditColoredTools.MOD_ID, id);
        var item = Registry.register(Registries.ITEM, itemId, new Item(new FabricItemSettings().maxCount(1)));
        ITEMS.add(item);
        return item;
    }

    private static void registerToolGroup(String prefix) {
        for (String color : COLORS) {
            registerTool("%s_%s".formatted(prefix, color));
        }
    }

    public static void registerTools() {
        TOOLS.forEach(Items::registerToolGroup);
        wand = registerTool("wand");
        Registry.register(Registries.ITEM_GROUP, new Identifier(WorldEditColoredTools.MOD_ID, "tools"), ITEM_GROUP);
    }
}
