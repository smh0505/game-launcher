<template>
    <div class="item">
        <img :src="image ? image : 'default.png'">
        <div class="item-name">{{ name }}</div>
    </div>
</template>

<script lang="ts">
export default {
    props: {
        image: String,
        name: String,
        path: String,
        link: String
    },
}
</script>

<style lang="scss">
.item {
    position: relative;
    height: 200px;
    overflow: hidden;

    border-radius: 12px;
    border: 4px solid oklch(30% var(--chroma) var(--hue));
    // background-color: oklch(50% var(--chroma) var(--hue) / 0.5);
    font: 16pt "Pretendard-Regular", sans-serif; 

    img {
        width: 100%;
        height: 100%;
        object-fit: cover;
        opacity: 50%;
        border-radius: 8px;
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
    }
}
</style>
