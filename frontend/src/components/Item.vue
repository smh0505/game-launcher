<template>
    <div class="item" @mouseenter="isHover = true" @mouseleave="isHover = false" @click="startGame">
        <img class="item-bg" :src="getItem.image ? getItem.image : 'default.png'" title="Thumbnail" :style="getBackground" draggable="false">
        <div class="item-name">{{ getItem.name }}</div>
        <button v-show="isHover" class="item-edit" @click.stop="showModal = true">
            <span class="material-symbols-outlined">feature_search</span>
        </button>
        <div class="handle" @click.stop><span class="material-symbols-outlined">open_with</span></div>

        <Teleport to="body">
            <Modal :show="showModal" @close="check">
                <template #modal-body>
                    <div class="modal-info">
                        <img class="thumbnail" :src="getItem.image ? getItem.image : 'default.png'" 
                            @click="loadThumbnail" title="Thumbnail" draggable="false">

                        <input class="game-input" type="text" v-model="getItem.name" placeholder="Game Name">

                        <div class="game-row">
                            <input class="game-input readonly" type="text" v-model="getItem.path" placeholder="Game Path" readonly>
                            <a @click="openDir"><span class="material-symbols-outlined">folder_open</span></a>
                        </div>

                        <div class="game-row">
                            <input class="game-input readonly" type="text" v-model="getItem.link" placeholder="Game Link" readonly>
                            <a @click="loadExe"><span class="material-symbols-outlined">edit_square</span></a>
                        </div>

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
        showModal: false,
        items: useItemStore()
    }),
    computed: {
        getItem() { return this.items.items[this.index] },
        getBackground() { return { opacity: this.isHover ? "100%" : "50%" } },
    },
    components: { Modal },
    methods: {
        check() {
            if (this.getItem.name && this.items.checkUnique(this.getItem.name)) {
                this.showModal = false
                if (this.getItem.path.substring(6) !== this.getItem.name) {
                    GoFunc.RenameFolder(this.getItem.name, this.getItem.path, this.getItem.link)
                        .then(x => {
                            this.getItem.path = x.path
                            this.getItem.link = x.link
                        })
                        .then(this.items.saveItems)
                }
            }
        },
        remove() {
            this.showModal = false
            GoFunc.DeleteGame(this.getItem.name, this.getItem.path, this.getItem.image)
                .then(x => { if (x) { this.items.items.splice(this.index, 1) } })
                .then(this.items.saveItems)
        },
        loadThumbnail() {
            GoFunc.LocateThumbnail(this.getItem.image, this.getItem.name)
                .then(x => { if (x) { this.getItem.image = x } })
        },
        loadExe() { GoFunc.LocateExecutive().then(x => { if (x) { this.getItem.link = x }})},
        openDir() { GoFunc.OpenFolder(this.getItem.path) },
        startGame() { GoFunc.Start(this.getItem.link) }
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

        object-fit: cover;
        transition: opacity 0.2s ease;
        user-select: none;
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
        text-shadow: 1px 1px 6px black, -1px -1px 6px black;
        user-select: none;
    }

    .item-edit {
        position: absolute;
        @include button(60px, 60px, 36pt);
        @include right-bottom(12px);
        transition: background 0.2s ease;
    }

    .handle {
        display: flex;
        position: absolute;
        @include left-bottom(0px);

        width: 40px;
        height: 40px;

        background-color: white;
        border-top: 4px solid oklch(30% var(--chroma) var(--hue));
        border-right: 4px solid oklch(30% var(--chroma) var(--hue));
        border-radius: 0px 12px;
        cursor: default;
        user-select: none;
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

        &:focus { box-shadow: 1px 1px 4px black, -1px -1px 4px black; }
        &.readonly { text-overflow: ellipsis; }
    }

    .game-delete {
        @include button(60px, 60px, 36pt);
        position: absolute;
        bottom: 30px;
        left: 40px;
        transition: background 0.2s ease;
    }
}
</style>
