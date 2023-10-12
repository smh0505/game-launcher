<template>
    <div id="sort-container">
        <a @click="items.sortItems('name', isReversed)">Name</a>
        <a @click="items.sortItems('code', isReversed)">Code</a>
        <a class="sort-right" @click="isReversed = !isReversed">
            <span class="material-symbols-outlined">
                {{ isReversed ? "check_box" : "check_box_outline_blank" }}
            </span>
        </a>
    </div>
</template>

<script lang="ts">
import { useItemStore } from '../stores/ItemStore.js'

export default {
    data: () => ({
        isReversed: false,
        items: useItemStore()
    })
}
</script>

<style lang="scss">
#sort-container {
    display: flex;
    position: absolute;
    justify-content: left;
    z-index: 99;

    width: calc(100% - 150px);
    height: 40px;

    border-right: 4px solid oklch(40% var(--chroma) var(--hue));

    a {
        display: flex;
        align-items: center;
        padding-inline: 12px;

        width: max-content;
        height: 100%;

        color: black;
        font: 12pt "Pretendard-Regular", sans-serif;
        transition: background 0.2s ease, color 0.2s ease;
        user-select: none;

        &:hover {
            background-color: white;
            color: black;
        }

        &.sort-right {
            position: absolute;
            right: 0px;
        }
    }
}
</style>
