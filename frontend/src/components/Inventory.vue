<template>
    <div id="inventory">
        <Item v-for="i in 10" :key="i"></Item>
    </div>
</template>

<script lang="ts">
import Item from './Item.vue';

export default {
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
