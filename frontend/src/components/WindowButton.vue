<template>
    <div id="wb-container">
        <a @mouseenter="isHover = true" @mouseleave="isHover = false">
            <span class="material-symbols-outlined">settings</span>
        </a>
        <a @click="minimize()"><span class="material-symbols-outlined">minimize</span></a>
        <a id="wb-close" @click="close()"><span class="material-symbols-outlined">close</span></a>
        <Transition>
            <div v-show="isHover" id="wb-setting" @mouseenter="isHover = true" @mouseleave="isHover = false">
                <input type="range" v-model="setting.chroma" min="0" max="0.4" step="0.001" @input="setting.setColor()">
                <label>Chroma: {{ setting.chroma }}</label>
                <input type="range" v-model="setting.hue" min="0" max="360" step="1" @input="setting.setColor()">
                <label>Hue: {{ setting.hue }}</label>
                <input type="range" v-model="setting.opacity" min="0" max="1" step="0.001" @input="setting.setColor()">
                <label>Bg. Opacity: {{ setting.opacity }}</label>
                <button @click="loadImage()">Load Image</button>
            </div>
        </Transition>
    </div>
</template>

<script lang="ts">
import { useSettingStore } from '../stores/SettingStore.js'
import { Quit, WindowMinimise } from '../../wailsjs/runtime/runtime.js'
import { OpenImageFile } from '../../wailsjs/go/main/App.js'

export default {
    data: () => ({
        isHover: false,
        setting: useSettingStore()
    }),
    methods: {
        minimize() { WindowMinimise() },
        close() { Quit() },
        async loadImage() {
            await OpenImageFile().then(x => { if (x) { this.setting.background = x }})
        }
    }
}
</script>

<style lang="scss">
#wb-container {
    display: grid;
    grid-template-columns: repeat(3, 1fr);

    position: absolute;
    right: 0px;
    z-index: 99;

    width: 150px;
    height: 40px;
    user-select: none;

    a {
        display: flex;
        width: 100%;
        height: 100%;
        transition: background 0.2s ease-out, color 0.2s ease-out;

        &:hover {
            background-color: oklch(40% var(--chroma) var(--hue) / 0.5);
            color: white;
            text-shadow: 0px 0px 12px black;
        }

        &#wb-close:hover {
            color: red;
        }
    }

    #wb-setting {
        --wails-draggable: none;

        display: flex;
        flex-direction: column;

        position: absolute;
        top: 40px;
        right: 100px;

        width: 200px;
        height: 200px;
        background-color: oklch(20% var(--chroma) var(--hue));
        color: white;

        button {
            margin-top: 8px;
            padding: 8px 12px;

            border: none;
            border-radius: 8px;
            background-color: transparent;
            color: white;
            transition: background 0.2s ease-out;

            &:hover {
                background-color: oklch(40% var(--chroma) var(--hue) / 0.5);
            }
        }
    }
}

.v-enter-from, .v-leave-to { opacity: 0; }
.v-enter-active, .v-leave-active { transition: opacity 0.2s ease-out; }
</style>
