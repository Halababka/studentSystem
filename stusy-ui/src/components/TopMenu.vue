<template>
  <nav>
    <ul id="topMenu">
<!--      <div class="search">-->
<!--        <input type="search" placeholder="Введите поисковой запрос">-->
<!--        <a href="#">-->
<!--          <button type="submit" class="topMenu_icon">-->
<!--            <img src="@/assets/search.svg" alt="Search">-->
<!--          </button>-->
<!--        </a>-->
<!--      </div>-->
<!--      <li class="menu-item">-->
      <!--        <a href="#">-->
      <!--          <span class="topMenu_icon">-->
      <!--            <img src="@/assets/bell.svg" alt="Bell">-->
      <!--          </span>-->
      <!--        </a>-->
      <!--      </li>-->
      <DropdownMenu class="menu-item" :title="``">
        <template v-slot:top-title>
          <a id="profile-title">
            <span>{{ name }}</span>
            <div class="avatar-arrow" v-bind:style="[menu?{'transform':'rotate(90deg)'}:{}]">
              <img src="@/assets/arrow.svg"/>
            </div>
          </a>
        </template>
        <template v-slot:item>
          <li class="sub-menu-item">
            <router-link to="/profile">Профиль</router-link>
          </li>
          <li class="sub-menu-item" v-on:click="logout"><a>Выйти</a></li>
        </template>
      </DropdownMenu>
    </ul>
  </nav>
  <div class="warn">
    <p><b>У вас не заполнен профиль.</b> Перейдите в
      <router-link to="profile">профиль</router-link>
      , чтобы заполнить данные.
    </p>

  </div>
</template>

<script setup>
import {logout} from "@/global";
import DropdownMenu from "@/components/DropdownMenu";
import {userData} from "@/composables/getUserData";
import {computed} from "vue";

const name = computed(() => {
  return userData.value.first_name && userData.value.last_name ? `${userData.value.first_name} ${userData.value.last_name[0]}.` : "Профиль";
});

let menu = false;

</script>

<style scoped>
#topMenu {
  display: flex;
  width: 100%;
  gap: 16px;
  flex-direction: row;
  justify-content: flex-end;
  align-items: center;
  list-style: none;
  margin: 0;
  padding: 0;
}

#topMenu li a {
  display: flex;
  color: var(--dark);
  align-items: center;
  justify-content: center;
  text-decoration: none;
  font-size: 18px;
  height: 100%;
  position: relative;
  white-space: nowrap;
}

#profile-title {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  gap: 5px;
  min-height: 70px;
}

.menu-item {
  font-style: normal;
  font-weight: 500;
  font-size: 18px;
  text-align: center;
  float: right;

}

.sub-menu-item {
  background-color: var(--light);
}

.sub-menu-item a {
  padding: 10px 0;
}

.sub-menu-item:last-child {
  border-bottom-right-radius: 10px;
  border-bottom-left-radius: 10px;
}

.sub-menu-item:hover {
  background: var(--light-mate);
}

#topMenu > li:last-child {
  background: var(--light);
  border-radius: 10px;
  padding: 0 36px;
  user-select: none;
}

.topMenu_icon {
  background: var(--light);
  border: none;
  border-radius: 15px;
  padding: 20px;
}

.search {
  position: relative;
  padding: 0;
  width: 50%;
}

.search button {
  position: relative;
  display: inline-block;
  cursor: pointer;
  z-index: 2;
  float: right;
}

input[type="search"] {
  position: absolute;
  height: 100%;
  width: 0;
  border: none;
  outline: none;
  top: 0;
  right: 0;
  z-index: 3;
  border-radius: 15px;
  background: none;
  padding: 1% 72px 1% 0;
  transition: width .4s cubic-bezier(0.000, 0.795, 0.000, 1.000);
  cursor: pointer;
  color: var(--dark);
  font-style: normal;
  font-weight: 400;
  font-size: 18px;
  line-height: 21px;
}

input[type="search"]:focus {
  width: 100%;
  z-index: 1;
  background-color: var(--light);
  padding: 1% 72px 1% 20px;
  cursor: text;
}

.userData {
  background: var(--light);
  padding: 12px;
  border-radius: 20px;
}

.userData p {
  color: #000000;
  font-size: 18px;
  font-weight: bold;
}

.warn {
  border: 1px solid #fe4a49;
  border-radius: 15px;
  padding: 15px;
  background: #fe4a4940;
  width: 80%;
  margin: 15px auto 0;
  max-width: 1050px;
  min-width: 310px;
}

@media screen and (max-width: 1000px) {
  nav {
    width: 100%;
  }

  #topMenu {
    float: right;
    flex-wrap: wrap;
    justify-content: center;
  }

  .search {
    width: 100%;
  }
}
</style>