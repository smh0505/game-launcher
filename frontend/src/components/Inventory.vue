<template>
    <draggable id="inventory" v-model="items.items" :group="game" handle=".handle" 
        item-key="index" animation="200" @end="resolveIndex">
        <template #item="{_, index}">
            <Item :index="index"></Item>
        </template>
    </draggable>
</template>

<script lang="ts">
import draggable from 'vuedraggable'
import { useItemStore, Game } from '../stores/ItemStore.js'
import Item from './Item.vue'

export default {
    data: () => ({
        items: useItemStore(),
        game: {} as Game
    }),
    components: { Item, draggable },
    methods: {
        resolveIndex() {
            this.items.items.forEach((x, index) => x.id = index)
            this.items.saveItems()
        }
    }
}
</script>

<style lang="scss">
#inventory {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 24px;
    padding: 24px;
    padding-bottom: 110px;

    width: 100%;
    height: max-content;
}
</style>
