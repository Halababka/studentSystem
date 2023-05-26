<template>
  <div>
    <div class="tab__header">
      <a href="#"
         class="tab__link">
        <span class="moving-icon" v-if="role === `teacher`">
          <img src="@/assets/moving.svg">
        </span>
        <div class="mainTitle" @click.prevent="active = !active">
        <span class="avatar-arrow" v-bind:style="[!active?{'transform':'rotate(-90deg)'}:{}]">
          <img src="@/assets/arrow.svg"/>
        </span>
          <strong>{{ title }}</strong>
        </div>
        <div class="createNewViewSection" v-if="addedIco === `true` && role === `teacher`">
          <dropdown-menu class="button-save" :title="`Вид раздела`">
            <template v-slot:top-title>
              <img src="@/assets/addedV2.svg"/>
            </template>
            <template v-slot:item>
              <li class="sub-menu-item"><a>Тест</a></li>
              <li class="sub-menu-item" @click="isModalOpen = true"><a>Ссылка</a></li>
              <li class="sub-menu-item"><a>Файлы</a></li>
              <li class="sub-menu-item"><a>Текст</a></li>
            </template>
          </dropdown-menu>
        </div>
        <span class="redactor-icon-pen" v-if="role === `teacher`">
          <img src="@/assets/pen.svg"/>
        </span>
        <span class="trash-icon" v-if="role === `teacher`">
          <img src="@/assets/trash.svg"/>
        </span>
        <strong>{{ titleAfter }}</strong>
      </a>
    </div>
    <transition name="fade">
      <div class="tab__content" v-show="active">
        <slot/>
      </div>
    </transition>
    <ModalWindow v-if="isModalOpen">
      <div class="modal-window">
        <img src="@/assets/close.svg" class="close-button" @click="isModalOpen = false">
        <h1>Создание ссылки:</h1>
        <label class="input-form">
          <label>Название</label>
          <input type="text" alt="coursesName"/>
        </label>
        <label class="input-form">
          <label>Ссылка</label>
          <input type="text" alt="coursesLink"/>
        </label>
      </div>
    </ModalWindow>
  </div>
</template>

<script>
import ModalWindow from "@/components/ModalWindow";
import DropdownMenu from "@/components/DropdownMenu";

export default {
  name: "AccordionMenu",
  components: {DropdownMenu, ModalWindow},
  props: [
    "title",
    "titleAfter",
    "role",
    "addedIco"
  ],
  data() {
    return {
      active: false,
      isModalOpen: false
    };
  }
};
</script>

<style scoped>
ul {
  list-style-type: none;
}

.button-save {
  color: white;
}

.moving-icon {
  margin-right: 28px;
  cursor: pointer !important;
}

.avatar-arrow {
  filter: invert(22%) sepia(61%) saturate(6241%) hue-rotate(228deg) brightness(102%) contrast(92%);
  margin-right: 13px;
  transition: .3s;
}

.viewSection-arrow {
  filter: invert(100%) sepia(0%) saturate(7492%) hue-rotate(22deg) brightness(106%) contrast(100%);
  margin-right: 16px;
}

.redactor-icon-pen {
  margin-right: 12px;
}

.trash-icon {
  margin-left: 5px;
  margin-right: 14px;
}

.mainTitle {
  display: flex;
  align-items: center;
  margin-right: 9px;
}

.tab__link {
  text-decoration: none;
  color: var(--dark);
}

.tab__content {
  margin-left: 22px;
}

.tab__header {
  display: flex;
  align-items: center;
}

.tab__header a {
  display: flex;
  flex-direction: row;
  align-items: center;
}

.tab__header a strong {
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
}

.createNewViewSection {
  margin-right: 12px;
}

.sub-menu-item {
  background-color: var(--blue);
  min-width: 100px;
  color: white;
}

.sub-menu-item a {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 5px 0;
  text-decoration: none;
  font-style: normal;
  font-weight: 400;
  font-size: 18px;
  line-height: 21px;
}

.sub-menu > li:last-child {
  border-radius: 0 0 10px 10px;
}

.sub-menu li:hover {
  background: var(--blue-mate);
}

.fade-enter-active {
  animation: scale-up-ver-top .2s cubic-bezier(0.390, 0.575, 0.565, 1.000) both;
}

.sub-menu-accordionMenu {
  box-shadow: #0000003b 6px 8px 9px 0px;
}

@keyframes scale-up-ver-top {
  0% {
    transform: scaleY(0.4);
    transform-origin: 100% 0;
  }
  100% {
    transform: scaleY(1);
    transform-origin: 100% 0;
  }
}
</style>