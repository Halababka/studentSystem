<template>
  <li class="dropdown" @click="isOpen = !isOpen"
      v-bind:style="[isOpen?{'border-bottom-right-radius': '0','border-bottom-left-radius': '0',}:{}]">
    <slot name="top-title"/>
    {{ title }}
    <div v-bind:style="[isOpen?{'transform':'rotate(90deg)', 'transition':'.3s'}:{}]">
      <slot name="bottom-title"/>
    </div>
    <transition name="slide-fade">
      <div class="sub-menu" v-if="isOpen" @click="hideDropdownMenu">
        <slot name="item"/>
      </div>
    </transition>
  </li>
</template>

<script>
export default {
  name: "DropdownMenu",
  props: ["title"],
  data() {
    return {
      isOpen: false
    };
  }
};
</script>

<style scoped>
slot {
  width: 100%;
}

.dropdown {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  gap: 5px;
  position: relative;
  font-size: 18px;
  cursor: pointer;
}

.sub-menu {
  position: absolute;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: stretch;
  cursor: pointer;
  user-select: none;
  z-index: 2;
  width: 100%;
  top: calc(100%);
  border-bottom-right-radius: 10px;
  border-bottom-left-radius: 10px;
  box-shadow: #0000003b 6px 8px 9px 0px;
}
</style>