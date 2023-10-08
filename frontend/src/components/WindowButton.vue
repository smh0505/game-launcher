<template>
    <div id="wb-container">
        <a @click="showModal = true"><span class="material-symbols-outlined">settings</span></a>
        <a @click="minimize()"><span class="material-symbols-outlined">minimize</span></a>
        <a id="wb-close" @click="close()"><span class="material-symbols-outlined">close</span></a>

        <Teleport to="body">
            <Modal :show="showModal" @close="showModal = false">
                <template #modal-body>
                    <div class="modal-info">
                        <label>Chroma: {{ setting.chroma }}</label>
                        <input type="range" v-model="setting.chroma" min="0" max="0.4" step="0.001" @input="setting.setColor">

                        <label>Hue: {{ setting.hue }}</label>
                        <input type="range" v-model="setting.hue" min="0" max="360" step="1" @input="setting.setColor">
                        
                        <label>Background Opacity: {{ setting.opacity }}</label>
                        <input type="range" v-model="setting.opacity" min="0" max="1" step="0.002" @input="setting.setColor">
                        
                        <div id="bg-merge">
                            <label>Merge Background</label>
                            <label>{{ setting.position.toUpperCase() }}</label>
                            <div id="merge-grid">
                                <button v-for="i in 9" @click="setPosition(i - 1)">
                                    <span class="material-symbols-outlined">{{ icons[i - 1] }}</span>
                                </button>
                            </div>
                        </div>
                        
                        <button id="bg-load" @click="loadImage"><span>Load Background Image</span></button>
                    </div>
                </template>
            </Modal>
        </Teleport>
    </div>
</template>

<script lang="ts">
import { useSettingStore } from '../stores/SettingStore.js'
import { Quit, WindowMinimise } from '../../wailsjs/runtime/runtime.js'
import { OpenImageFile } from '../../wailsjs/go/main/App.js'
import Modal from './Modal.vue'

const positions = [
    'left top', 'top', 'right top',
    'left', 'center', 'right',
    'left bottom', 'bottom', 'right bottom'
]

export default {
    data: () => ({
        showModal: false,
        setting: useSettingStore(),
        icons: [
            'north_west', 'north', 'north_east',
            'west', 'my_location', 'east',
            'south_west', 'south', 'south_east'
        ]
    }),
    methods: {
        minimize() { WindowMinimise(); },
        close() { Quit(); },
        setPosition(i: number) {
            this.setting.position = positions[i]
            this.setting.setColor()
        },
        loadImage() {
            OpenImageFile(this.setting.background)
                .then(x => { if (x) { this.setting.background = x; } })
                .then(this.setting.setColor)
                .then()
        }
    },
    components: { Modal }
}
</script>

<style lang="scss">
@import "../style.scss";

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
        transition: background 0.2s ease, color 0.2s ease;

        &:hover {
            background-color: oklch(40% var(--chroma) var(--hue) / 0.5);
            color: white;
            text-shadow: 0px 0px 12px black;
        }

        &#wb-close:hover { color: red; }
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
            transition: background 0.2s ease;

            &:hover { background-color: oklch(40% var(--chroma) var(--hue) / 0.5); }
        }
    }
}

.modal-info {
    input[type=range] { width: 100%; }
    label { font: 16pt "Pretendard-Regular", sans-serif; }

    #bg-merge {
        display: flex;
        flex-direction: column;
        gap: 8px;

        position: absolute;
        @include left-bottom(40px);

        #merge-grid {
            display: grid;
            grid-template-columns: repeat(3, 1fr);
            grid-template-rows: repeat(3, 1fr);

            position: relative;
            gap: 10px;

            button { 
                @include button(44px, 44px, 24pt);
                transition: background 0.2s ease; 
            }
        }
    }

    #bg-load {
        position: absolute;
        top: 230px;
        left: 240px;
        @include button(calc(100% - 280px), 40px, 14pt);
        transition: background 0.2s ease;
    }
}

.v-enter-from, .v-leave-to { opacity: 0; }
.v-enter-active, .v-leave-active { transition: opacity 0.2s ease; }
</style>
