<template>
    <div id="background">
        <WindowButton></WindowButton>
        <div id="image"></div>
        <div id="foreground">
            <!--<input type="text" id="search-bar" v-model="search">-->
            <div id="content">
                <Inventory></Inventory>
            </div>
            <button id="install" @click="install()">
                <span class="material-symbols-outlined">post_add</span>
            </button>
        </div>
    </div>
</template>

<script lang="ts">
import WindowButton from './components/WindowButton.vue'
import Inventory from './components/Inventory.vue'
import { useSettingStore } from './stores/SettingStore.js'
import { useItemStore } from './stores/ItemStore.js'
import { InstallGame } from '../wailsjs/go/main/App.js'

export default {
    data: () => ({
        setting: useSettingStore(),
        items: useItemStore(),
        search: ""
    }),
    components: { WindowButton, Inventory },
    methods: {
        install() {
            InstallGame().then(x => {
                if (x) { this.items.addItem(x.path, x.name, x.link) }
                this.items.saveItems()
            })
        }
    },
    mounted() {
        this.setting.loadSetting()
        this.items.loadItems()

        window.addEventListener("beforeunload", this.setting.saveSetting)
    }
}
</script>

<style lang="scss">
@import "./style.scss";

#background {
    --wails-draggable: drag;

    width: 100vw;
    height: 100vh;
    background-color: oklch(70% var(--chroma) var(--hue) / 0.5);

    #image {
        position: absolute;
        width: 100%;
        height: 100%;

        background-image: var(--background);
        background-size: contain;
        background-position: var(--position);
        background-repeat: no-repeat;
        opacity: var(--opacity);
    }

    #foreground {
        position: absolute;
        width: 100%;
        height: 100%;

        input[type=text] {
            --wails-draggable: none;

            width: calc(100% - 150px);
            height: 40px;
        }

        #search-bar {
            padding-inline: 6px;
            border: 4px solid oklch(30% var(--chroma) var(--hue));
            background-color: oklch(50% var(--chroma) var(--hue) / 0.5);
            font: 16pt "Pretendard-Regular", sans-serif;
        }

        #content {
            position: absolute;
            bottom: 0px;
            width: 100%;
            height: calc(100% - 40px);
            overflow-y: auto;
        }

        #install {
            position: absolute;
            @include button(60px, 60px, 36pt);
            @include left-bottom(24px);
            transition: background 0.2s ease;
        }
    }
}
</style>
