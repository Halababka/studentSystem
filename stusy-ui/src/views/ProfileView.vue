<template>
  <div class="mainContent">
    <div class="personalData">
      <div class="photoName">
        <div class="photo">
          <img src="@/assets/avatar.jpg" alt="Фото профиля" draggable="false">
        </div>
        <div class="nameNum">
          <p class="name">{{ userData.first_name }} {{ userData.last_name }}</p>
          <p>Номер студенческого: {{ localUserData.id }}</p>
        </div>
      </div>
      <form v-on:submit.prevent="submit(localUserData)">
        <div class="personalDataForm">
          <label class="input-form">
            <p>Имя</p>
            <input
                v-model="localUserData.first_name"
                type="text"
                placeholder="Введите имя"
                required
            />
          </label>
          <label class="input-form">
            <p>Фамилия</p>
            <input
                v-model="localUserData.last_name"
                type="text"
                placeholder="Введите фамилию"
                required
            />
          </label>
          <label class="input-form">
            <p>Отчество</p>
            <input
                type="text"
                placeholder="Введите отчество"
            />
          </label>
        </div>
        <div class="button-login">
          <input type="submit" value="Сохранить"/>
        </div>
      </form>

    </div>
  </div>
</template>

<script setup>
import {userData} from "@/composables/getUserData";
import {getCookie} from "@/global";
import {ref} from "vue";

let localUserData = ref({
  id: getCookie("ID"),
  first_name: userData.value.first_name,
  last_name: userData.value.last_name,
});
</script>

<style scoped>

.mainContent {
  display: flex;
  justify-content: space-evenly;
  align-items: stretch;
  gap: 30px;
  padding: 57px 29px;
  flex-wrap: wrap;
}

.personalData {
  background: var(--light);
  padding: 27px 32px;
  border-radius: 25px;
  width: 553px;
}

.photoName {
  display: flex;
  gap: 25px;
}

.photo img {
  width: 72px;
  border-radius: 15px;
  user-select: none;
}

.personalDataForm {
  display: flex;
  gap: 15px;
  flex-wrap: wrap;
}

input:disabled,
input[disabled] {
  background: #000 !important;
}

.name {
  margin-bottom: 10px;
  font-size: 17px;
  font-weight: bold;
}

.input-form input {
  background: var(--dark-mate-mate);
  width: auto;
}

.input-form label span {
  font-size: 18px;
}

</style>