<template>
    <draggable id="inventory" :style="getGap" v-model="items.items" :group="game" handle=".handle" 
        item-key="index" animation="200" @end="resolveIndex">
        <template #item="{_, index}">
            <Item :index="index"></Item>
        </template>
    </draggable>
</template>

<script lang="ts">
import draggable from 'vuedraggable'
import { useSettingStore } from '../stores/SettingStore.js'
import { useItemStore, Game } from '../stores/ItemStore.js'
import Item from './Item.vue'

export default {
    data: () => ({
        setting: useSettingStore(),
        items: useItemStore(),
        game: {} as Game
    }),
    components: { Item, draggable },
    computed: {
        getGap() { return { gap: this.setting.isListview ? "8px" : "24px" } }
    },
    methods: {
        resolveIndex() {
            this.setting.isSaving = true
            this.items.items.forEach((x, index) => x.id = index)
            this.items.saveItems()
        }
    }
}
</script>

<style lang="scss">
#inventory {
    display: flex;
    flex-direction: row;
    justify-content: left;
    align-items: start;
    flex-wrap: wrap;

    padding: 24px;
    padding-bottom: 110px;

    width: 100%;
    height: max-content;
}
</style>
