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
            <Teleport to="body">
                <Modal :show="showAlert" @close="close">
                    <template #modal-body>
                        <div class="modal-info">
                            <h1>Internal Error</h1>
                            <p>{{ errors[error] }}</p>
                            <p v-if="error === 1">{{ dir }}</p>
                            <input v-if="error === 1" class="game-input" type="text" v-model="pswd" placeholder="Insert the password here.">
                            <button v-if="error === 1" class="game-delete" @click="retry">
                                <span class="material-symbols-outlined">resume</span>
                            </button>
                        </div>
                    </template>
                </Modal>
            </Teleport>
        </div>
    </div>
</template>

<script lang="ts">
import WindowButton from './components/WindowButton.vue'
import Inventory from './components/Inventory.vue'
import Sort from './components/Sort.vue'
import Modal from './components/Modal.vue'
import { useSettingStore } from './stores/SettingStore.js'
import { useItemStore } from './stores/ItemStore.js'
import { TestGame } from '../wailsjs/go/main/App.js'

export default {
    data: () => ({
        setting: useSettingStore(),
        items: useItemStore(),
        dir: "",
        pswd: "",
        error: 0,
        errors: [
            "No file selected.",
            "Password required for next file:"
        ],
        showAlert: false
    }),
    components: { WindowButton, Inventory, Sort, Modal },
    methods: {
        install() {
            this.setting.isSaving = true
            TestGame(this.dir, this.pswd).then(res => {
                if (!res.success) {
                    this.dir = res.dir
                    this.setting.isSaving = false
                    this.showAlert = true
                    this.error = res.error
                } else {
                    this.dir = ""
                    this.pswd = ""
                    this.items.addItem(res.path, res.name, res.link)
                    this.items.saveItems()
                }
            })
        },
        retry() {
            this.showAlert = false
            this.install()
        },
        close() {
            this.dir = ""
            this.pswd = ""
            this.showAlert = false
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

.modal-info {
    h1 { 
        font: bold 32pt "Pretendard-Regular", sans-serif; 
        text-shadow: 2px 2px 4px red, -2px -2px 4px red;
    }
    p { font: 16pt "Pretendard-Regular", sans-serif; }
}

.notice-leave-to { opacity: 0; }
.notice-leave-active { transition: opacity 2s ease; }
</style>
