<template>
    <div id="inventory">
        <Item v-for="(item, index) in items.items" 
            :key="index" 
            :name="item.name" 
            :path="item.path" 
            :link="item.link" 
            :image="item.image"></Item>
    </div>
</template>

<script lang="ts">
import { useItemStore } from '../stores/ItemStore.js'
import Item from './Item.vue'

export default {
    data: () => ({
        items: useItemStore()
    }),
    components: { Item }
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
