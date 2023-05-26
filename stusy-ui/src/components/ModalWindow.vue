<template>
  <div class="modal">
    <slot>
    </slot>
  </div>
</template>

<script>
export default {
  name: "ModalWindow"
};
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.modal {
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #000000a6;
  z-index: 100;
}

</style>
