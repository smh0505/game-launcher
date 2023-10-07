<template>
    <div id="background">
        <WindowButton></WindowButton>
        <div id="image" :style="getBackground"></div>
        <div id="foreground">
            <input type="text" v-model="search">
            <div id="content">
                <Inventory></Inventory>
            </div>
            <button @click="install()">
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
    computed: {
        getBackground() {
            return {
                "background-image": `url(${this.setting.background})`
            }
        }
    },
    methods: {
        install() {
            InstallGame().then(x => {
                if (x) { this.items.addItem(x.path, x.name, x.link) }
            })
        }
    },
    mounted() {
        this.setting.loadSetting()
        this.items.loadItems()

        window.addEventListener("beforeunload", () => {
            this.setting.saveSetting()
            this.items.saveItems()
        })
    }
}
</script>

<style lang="scss">
#background {
    --wails-draggable: drag;

    width: 100vw;
    height: 100vh;
    background-color: oklch(70% var(--chroma) var(--hue) / 0.5);
    background-size: contain;
    background-position: center;
    background-repeat: no-repeat;
    background-blend-mode: multiply;

    #image {
        position: absolute;
        width: 100%;
        height: 100%;

        background-size: contain;
        background-position: bottom;
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

        #content {
            position: relative;
            width: 100%;
            height: calc(100% - 40px);
            overflow-y: auto;
        }

        button {
            position: absolute;
            left: 24px;
            bottom: 24px;
            
            width: 60px;
            height: 60px;

            background-color: oklch(55% var(--chroma) var(--hue) / 0.5);
            color: white;
            border: none;
            border-radius: 12px;
            transition: background 0.2s ease-out, color 0.2s ease-out;

            &:hover {
                background-color: oklch(20% var(--chroma) var(--hue) / 0.5);
                color: black;
            }

            span { font-size: 36pt; }
        }
    }
}
</style>
