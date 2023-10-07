<template>
    <div class="item" @mouseenter="isHover = true" @mouseleave="isHover = false" @click="startGame">
        <div class="item-bg" :style="getBackground"></div>
        <div class="item-name">{{ getItem.name }}</div>
        <button v-show="isHover" class="button item-edit" @click.stop="showModal = true">
            <span class="material-symbols-outlined">feature_search</span>
        </button>
        <Teleport to="body">
            <Modal :show="showModal" @close="check">
                <template #modal-body>
                    <div class="modal-info">
                        <img class="thumbnail" :src="getItem.image ? getItem.image : 'default.png'" title="Thumbnail" draggable="false">
                        <div class="game-row">
                            <input class="game-input" type="text" v-model="getItem.image" placeholder="Game Thumbnail" style="text-overflow: ellipsis;" readonly>
                            <a @click="loadThumbnail"><span class="material-symbols-outlined">edit_square</span></a>
                        </div>
                        <input class="game-input" type="text" v-model="getItem.name" placeholder="Game Name">
                        <div class="game-row">
                            <input class="game-input" type="text" v-model="getItem.path" placeholder="Game Path" style="text-overflow: ellipsis;" readonly>
                            <a @click="openDir"><span class="material-symbols-outlined">folder_open</span></a>
                        </div>
                        <div class="game-row">
                            <input class="game-input" type="text" v-model="getItem.link" placeholder="Game Link" style="text-overflow: ellipsis;" readonly>
                            <a @click="loadExe"><span class="material-symbols-outlined">edit_square</span></a>
                        </div>
                    </div>
                </template>
            </Modal>
        </Teleport>
    </div>
</template>

<script lang="ts">
import Modal from './Modal.vue'
import { useItemStore } from '../stores/ItemStore.js'
import { LocateThumbnail, LocateExecutive, Start, OpenFolder } from '../../wailsjs/go/main/App.js'

export default {
    props: {
        index: {
            type: Number,
            required: true
        }
    },
    data: () => ({
        isHover: false,
        showModal: false,
        items: useItemStore()
    }),
    computed: {
        getItem() { return this.items.items[this.index] },
        getBackground() {
            return {
                "background-image": `url(${this.getItem.image ? this.getItem.image : "default.png"})`,
                opacity: this.isHover ? "100%" : "50%"
            };
        },
    },
    components: { Modal },
    methods: {
        check() {
            if (this.getItem.name && this.getItem.path) {
                this.showModal = false
                this.items.saveItems()
            }
        },
        loadThumbnail() { LocateThumbnail().then(x => { if (x) { this.getItem.image = x }})},
        loadExe() { LocateExecutive().then(x => { if (x) { this.getItem.link = x }})},
        openDir() { OpenFolder(this.getItem.path) },
        startGame() { Start(this.getItem.link) }
    }
}
</script>

<style lang="scss">
@import "../style.scss";

.item {
    --wails-draggable: none;

    position: relative;
    height: 200px;
    overflow: hidden;

    border-radius: 12px;
    border: 4px solid oklch(30% var(--chroma) var(--hue));
    // background-color: oklch(50% var(--chroma) var(--hue) / 0.5);
    font: 16pt "Pretendard-Regular", sans-serif; 

    .item-bg {
        width: 100%;
        height: 100%;

        background-size: cover;
        background-position: center;
        background-repeat: no-repeat;
        transition: opacity 0.2s ease;
    }

    .item-name {
        position: absolute;
        top: 0px;
        left: 0px;
        padding: 4px;

        width: 300px;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;

        color: white;
        text-shadow: 4px 4px 12px black, -4px -4px 12px black;
        user-select: none;
    }

    .item-edit {
        @include right-bottom(12px);
    }
}

.modal-info {
    display: flex;
    flex-direction: column;
    gap: 8px;

    .thumbnail {
        width: 100%;
        height: 160px;
        object-fit: cover;

        border: 2px solid black;
        border-radius: 8px;
    }

    .game-row {
        display: flex;
        width: 100%;
        gap: 8px;

        a {
            display: flex;
            padding: 4px;
            border-radius: 8px;
            transition: background 0.2s ease;

            &:hover {
                background-color: gray;
            }
        }
    }

    .game-input {
        width: 100%;
        height: 32px;
        font: 16pt "Pretendard-Regular", sans-serif;
        padding-inline: 6px;

        border: 2px solid black;
        border-radius: 8px;
        transition: box-shadow 0.2s ease;
        cursor: default !important;

        &:focus {
            box-shadow: 1px 1px 4px black, -1px -1px 4px black;
        }
    }
}
</style>
