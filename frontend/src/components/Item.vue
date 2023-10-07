<template>
    <div class="item">
        <img :src="image" title="thumbnail">
        <div>{{ name }}</div>
    </div>
</template>

<script lang="ts">
export default {
    props: {
        image: String,
        name: String,
        path: String,
        link: String
    }
}
</script>

<style lang="scss">
.item {
    height: 200px;
    border-radius: 12px;
    border: 4px solid oklch(30% var(--chroma) var(--hue));
    background-color: oklch(50% var(--chroma) var(--hue) / 0.5);
    font: 16pt "Pretendard-Regular", sans-serif; 

    img {
        
    }
}
</style>
