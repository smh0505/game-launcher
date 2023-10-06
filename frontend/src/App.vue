<template>
    <div id="background">
        <WindowButton></WindowButton>
        <div id="image" :style="getBackground"></div>
        <div id="foreground">
            <input type="text" v-model="search">
        </div>
    </div>
</template>

<script lang="ts">
import WindowButton from './components/WindowButton.vue'
import { useSettingStore } from './stores/SettingStore.js'

export default {
    data: () => ({
        setting: useSettingStore(),
        search: ""
    }),
    components: { WindowButton },
    computed: {
        getBackground() {
            return {
                "background-image": `url(${this.setting.background})`
            }
        }
    },
    mounted() {
        this.setting.loadSetting()

        window.addEventListener("beforeunload", this.setting.saveSetting) 
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
        background-position: center;
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
    }
}
</style>
