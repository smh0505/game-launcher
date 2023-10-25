<template>
    <div class="item" :style="getListview" @mouseenter="isHover = true" @mouseleave="isHover = false" @click="startGame">
        <div :class="setting.isListview ? 'item-list-view' : 'item-gall-view'">
            <div v-if="setting.isListview" class="item-text-bg" :style="getBackground"></div>
            <img class="item-bg" :src="getItem.image ? getItem.image : 'default.png'" title="Thumbnail" :style="getBackground" draggable="false">
            <div v-if="getItem.code" class="item-caption">{{ getItem.code }}</div>
            <div :style="getTopPos" class="item-name">{{ getItem.name }}</div>
            <button v-show="isHover" class="item-edit" @click.stop="showData = true">
                <span class="material-symbols-outlined">feature_search</span>
            </button>
            <a class="item-star" @click.stop="items.setIfPlayed(index)">
                <span class="material-symbols-outlined">{{ isChecked[getItem.isPlayed] }}</span>
                <span class="material-symbols-outlined">radio_button_unchecked</span>
            </a>
            <div class="handle" @click.stop><span class="material-symbols-outlined">open_with</span></div>
        </div>

        <Teleport to="body">
            <Modal :show="showData" @close="check">
                <template #modal-body>
                    <div class="modal-info">
                        <img class="thumbnail" :src="getItem.image ? getItem.image : 'default.png'" 
                            @click="loadThumbnail" title="Thumbnail" draggable="false">

                        <input class="game-input" type="text" v-model="getItem.name" placeholder="Game Name">
                        <input class="game-input" type="text" v-model="getItem.code" placeholder="Game Caption (ex. Product Number)">

                        <div class="game-row">
                            <input class="game-input readonly" type="text" v-model="getItem.path" placeholder="Game Path" readonly>
                            <a @click="openDir"><span class="material-symbols-outlined">folder_open</span></a>
                        </div>

                        <div class="game-row">
                            <input class="game-input readonly" type="text" v-model="getItem.link" placeholder="Game Link" readonly>
                            <a @click="loadExe"><span class="material-symbols-outlined">edit_square</span></a>
                        </div>

                        <button class="game-delete" @click="swapModal()">
                            <span class="material-symbols-outlined">delete_forever</span>
                        </button>
                    </div>
                </template>
            </Modal>
            <Modal :show="showDelete" @close="swapModal()">
                <template #modal-body>
                    <div class="modal-info">
                        <h1>Warning</h1>
                        <p>
                            You are about to delete <strong>{{ getItem.name }}</strong>.<br>
                            Do you want to continue?
                        </p>
                        <button class="game-delete" @click="remove()">
                            <span class="material-symbols-outlined">delete_forever</span>
                        </button>
                    </div>
                </template>
            </Modal>
        </Teleport>
    </div>
</template>

<script lang="ts">
import Modal from './Modal.vue'
import { useSettingStore } from '../stores/SettingStore.js'
import { useItemStore } from '../stores/ItemStore.js'
import * as GoFunc from '../../wailsjs/go/main/App.js'

export default {
    props: {
        index: {
            type: Number,
            required: true
        }
    },
    data: () => ({
        isHover: false,
        showData: false,
        showDelete: false,
        items: useItemStore(),
        setting: useSettingStore(),
        isChecked: [
            "radio_button_unchecked",
            "radio_button_partial",
            "radio_button_checked"
        ]
    }),
    computed: {
        getItem() { return this.items.items[this.index] },
        getBackground() {
            return {
                opacity: this.isHover ? "100%" : "50%",
                transition: "opacity 0.2s ease"
            }
        },
        getTopPos() { return { top: this.getItem.code ? "16px" : "0px" } },
        getListview() {
            return {
                width: this.setting.isListview ? "100%" : "309px",
                height: this.setting.isListview ? "60px" : "200px"
            }
        }
    },
    components: { Modal },
    methods: {
        check() {
            if (this.getItem.name && this.items.checkUnique(this.getItem.name)) {
                this.showData = false
                const last = this.getItem.path.substring(this.getItem.path.lastIndexOf('\\'))
                if (last !== this.getItem.name) {
                    this.setting.isSaving = true
                    GoFunc.RenameFolder(this.getItem.name, this.getItem.path, this.getItem.link)
                        .then(x => {
                            this.getItem.path = x.path
                            this.getItem.link = x.link
                        })
                        .then(this.items.saveItems)
                }
            }
        },
        swapModal() {
            this.showData = !this.showData
            this.showDelete = !this.showDelete
        },
        remove() {
            this.showDelete = false
            GoFunc.DeleteGame(this.getItem.path, this.getItem.image)
                .then(() => this.items.items.splice(this.index, 1))
                .then(this.items.saveItems)
        },
        loadThumbnail() {
            GoFunc.LocateThumbnail(this.getItem.image, this.getItem.name)
                .then(x => { if (x) { this.getItem.image = x } })
        },
        loadExe() { GoFunc.LocateExecutive(this.getItem.path).then(x => { if (x) { this.getItem.link = x }})},
        openDir() { GoFunc.OpenFolder(this.getItem.path) },
        startGame() { GoFunc.Start(this.getItem.link) }
    }
}
</script>

<style lang="scss">
@import "../style.scss";

@mixin item-label($left, $width, $color) {
    position: absolute;
    left: $left;
    padding: 4px;

    width: $width;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;

    color: $color;
    text-shadow: 1px 1px 6px black, -1px -1px 6px black;
    user-select: none;
}

.item {
    --wails-draggable: none;

    display: block;
    position: relative;
    overflow: hidden;

    border-radius: 12px;
    border: 4px solid oklch(30% var(--chroma) var(--hue));
    font: 16pt "Pretendard-Regular", sans-serif; 

    .item-gall-view, .item-list-view {
        width: 100%;
        height: 100%;
    }

    .item-bg {
        object-fit: cover;
        user-select: none;
    }

    .item-caption {
        top: 0px;
        font: 12pt "Pretendard-Regular", sans-serif;
    }

    .item-edit {
        position: absolute;
        transition: background 0.2s ease;
    }

    .handle {
        display: flex;
        position: absolute;
        @include left-bottom(0px);

        background-color: white;
        border-right: 4px solid oklch(30% var(--chroma) var(--hue));
        cursor: default;
        user-select: none;
    }

    .item-star {
        display: flex;
        position: absolute;
        user-select: none;

        span { 
            position: absolute;
            font-size: 36px;
            text-shadow: 0px 0px 8px black;

            &:first-child {
                color: yellow;
            }
        }
    }

    .item-gall-view {
        .item-bg { width: 100%; height: 100%; }
        .item-caption { @include item-label(0px, 300px, yellow); }
        .item-name { @include item-label(0px, 300px, white); }

        .item-edit {
            @include button(60px, 60px, 36pt);
            @include right-bottom(12px);
        }

        .handle {
            width: 40px;
            height: 40px;

            border-top: 4px solid oklch(30% var(--chroma) var(--hue));
            border-radius: 0px 12px;
        }

        .item-star {
            bottom: 0px;
            left: 40px;

            width: 40px;
            height: 40px;
        }
    }

    .item-list-view {
        .item-text-bg {
            position: absolute;
            right: 0px;
            width: calc(100% - 248px);
            height: 100%;
            background-color: white;
        }

        .item-bg {
            position: absolute;
            left: 48px;
            width: 200px;
            height: 100%;
        }

        .item-caption { @include item-label(249px, calc(100% - 250px), yellow); }
        .item-name { @include item-label(249px, calc(100% - 250px), white); }

        .item-edit {
            @include button(48px, 36px, 20pt);
            @include right-bottom(8px);
        }

        .handle {
            width: 48px;
            height: 100%;
        }

        .item-star {
            bottom: 6px;
            left: 54px;

            width: 40px;
            height: 40px;
        }
    }
}

.modal-info {
    .game-delete {
        @include button(60px, 60px, 36pt);
        position: absolute;
        bottom: 30px;
        left: 40px;
        transition: background 0.2s ease;
    }
}
</style>
