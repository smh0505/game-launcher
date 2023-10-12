<template>
    <div id="background">
        <WindowButton></WindowButton>
        <Sort></Sort>
        <img id="bg-image" :src="setting.background ? setting.background : ''" title="background" draggable="false">
        <div id="foreground">
            <div id="content">
                <Inventory></Inventory>
            </div>
            <button id="install" @click="install()">
                <span class="material-symbols-outlined">post_add</span>
            </button>
            <Transition name="notice">
                <div v-if="setting.isSaving" id="notice-saving">Saving...</div>
            </Transition>
        </div>
    </div>
</template>

<script lang="ts">
import WindowButton from './components/WindowButton.vue'
import Inventory from './components/Inventory.vue'
import Sort from './components/Sort.vue'
import { useSettingStore } from './stores/SettingStore.js'
import { useItemStore } from './stores/ItemStore.js'
import { InstallGame } from '../wailsjs/go/main/App.js'

export default {
    data: () => ({
        setting: useSettingStore(),
        items: useItemStore(),
        search: ""
    }),
    components: { WindowButton, Inventory, Sort },
    methods: {
        install() {
            this.setting.isSaving = true
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

    #bg-image {
        position: absolute;
        bottom: 0px;
        width: 100%;
        height: calc(100% - 40px);

        object-fit: contain;
        object-position: var(--position);
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
            position: absolute;
            bottom: 0px;
            width: 100%;
            height: calc(100% - 40px);
            overflow-y: auto;

            &::-webkit-scrollbar { display: none; }
        }

        #install {
            position: absolute;
            @include button(60px, 60px, 36pt);
            @include left-bottom(24px);
            transition: background 0.2s ease;
        }

        #notice-saving {
            display: flex;
            position: absolute;
            @include right-bottom(24px);

            width: 100px;
            height: 50px;

            font: 14pt "Pretendard-Regular", sans-serif;
            border-radius: 12px;
            background-color: oklch(55% var(--chroma) var(--hue) / 0.7);
            color: white;
            box-shadow: 0px 0px 12px black;
        }
    }
}

.notice-leave-to { opacity: 0; }
.notice-leave-active { transition: opacity 2s ease; }
</style>
