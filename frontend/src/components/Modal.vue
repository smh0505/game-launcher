<template>
    <Transition name="modal">
        <div v-if="show" id="modal-mask" @mousedown="$emit('close')">
            <div id="modal-container" @mousedown.stop>
                <slot name="modal-body"></slot>
                <button id="modal-button" @click="$emit('close')">
                    <span class="material-symbols-outlined">close</span>
                </button>
            </div>
        </div>
    </Transition>
</template>

<script lang="ts">
export default {
    props: {
        show: {
            type: Boolean,
            required: true
        }
    }
}
</script>

<style lang="scss">
@import "../style.scss";

#modal-mask {
    position: fixed;
    z-index: 999;
    top: 0px;
    left: 0px;

    width: 100%;
    height: 100%;
    background-color: oklch(0% 0% 0 / 0.5);

    display: flex;
    transition: opacity 0.2s ease;
}

#modal-container {
    position: relative;
    width: 540px;
    height: 480px;
    padding: 30px 40px;

    background-color: white;
    border-radius: 16px;
    box-shadow: 2px 2px 8px black, -2px -2px 8px black;
    user-select: none;
    
    transition: all 0.2s ease;
}

#modal-button {
    position: absolute;
    @include button(60px, 60px, 36pt);
    bottom: 30px;
    right: 40px;
    transition: background 0.2s ease;
}

.modal-enter-from, .modal-leave-to {
    opacity: 0%;
    
    #modal-container { 
        transform: translateY(33%) scale(1.1);
    }
}
</style>
