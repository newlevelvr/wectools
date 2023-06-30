package ru.newlevelvr.wectools;

import net.fabricmc.api.ModInitializer;
import ru.newlevelvr.wectools.item.Items;

public class WorldEditColoredTools implements ModInitializer {

    public static final String MOD_ID = "wectools";

    @Override
    public void onInitialize() {
        Items.registerTools();
    }
}
