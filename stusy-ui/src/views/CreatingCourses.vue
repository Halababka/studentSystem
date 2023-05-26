<template>
  <div class="mainContent">
    <div class="coursesMainLeft">
      <div class="redactorContentCreating-topRow">
        <h1>Название:</h1>
        <div class="button-save">
          <img src="@/assets/diskette.svg">
          <input type="submit" value="Сохранить"/>
        </div>
      </div>
      <label class="input-form">
        <input type="text" alt="coursesName"/>
      </label>
      <div class="redactorContentCreating-headerOfSection">
        <h1>Разделы</h1>
        <ul class="createNew">
          <dropdown-menu class="button-save" :title="`Создать новый`">
            <template v-slot:top-title>
              <img src="@/assets/added.svg">
            </template>
            <template v-slot:item>
              <li class="sub-menu-item"><a>C подпунктами</a></li>
              <li class="sub-menu-item"><a>Ссылка</a></li>
              <li class="sub-menu-item"><a>Объявление</a></li>
            </template>
          </dropdown-menu>
        </ul>
      </div>
      <ul class="accordionMenu-list">
        <li>
          <img class="moving-icon" src="@/assets/moving.svg">
          <a href="">Введение</a>
        </li>
        <li>
          <AccordionMenu title="Модуль 1" addedIco="true" v-bind:role="this.role" class="accordionMenu"></AccordionMenu>
        </li>
        <li>
          <AccordionMenu title="Модуль 2" addedIco="true" v-bind:role="this.role" class="accordionMenu">
            <ol>
              <li><a href="#">Как сжать звуковые файлы</a></li>
              <li><a href="#">Как сжать графические файлы</a></li>
              <li><a href="#">Какими возможностями обладают архиваторы</a></li>
              <li><a href="#">&#10071; Контрольная точка 2 - Тестирование</a></li>
            </ol>
          </AccordionMenu>
        </li>
        <li>
          <AccordionMenu title="Модуль 3" addedIco="true" v-bind:role="this.role" class="accordionMenu"></AccordionMenu>
        </li>
      </ul>
    </div>
    <div class="coursesMainAbout">
      <h1>О курсе</h1>
      <label class="input-form">
        <p>Количество з.ед:</p>
        <input type="text" alt="coursesName"/>
      </label>
    </div>
  </div>
</template>

<script>
import AccordionMenu from "@/components/AccordionMenu";
import DropdownMenu from "@/components/DropdownMenu";

export default {
  name: "CreatedCourses",
  components: {DropdownMenu, AccordionMenu},
  data() {
    return {
      role: "teacher"
    };
  }
};
</script>

<style scoped>
p {
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
  margin-bottom: 8px;
}

ul {
  list-style-type: none;
}

/*li {*/
/*  display: flex;*/
/*  flex-direction: row;*/
/*}*/

h1 {
  font-style: normal;
  font-weight: 400;
  font-size: 36px;
  line-height: 41px;
}

a {
  cursor: pointer;
  text-decoration: none;
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
  color: var(--blue);
}

.button-save {
  color: white;
}

.input-form, .input-form input {
  margin: 0;
}

.moving-icon {
  margin-right: 28px;
  cursor: pointer !important;
}


.accordionMenu ol {
  display: flex;
  flex-direction: column;
  gap: 30px;
  margin-top: 30px;
  margin-left: 51px;
}

/*.sub-menu {*/
/*  position: absolute;*/
/*  display: flex;*/
/*  flex-direction: column;*/
/*  right: 0;*/
/*  left: 0;*/
/*  top: 55px;*/
/*  border-radius: 0 0 15px 15px;*/
/*  background-color: var(--blue);*/
/*}*/

.sub-menu-item {
  background-color: var(--blue);
}

/*.active {*/
/*  padding: 17px 15px 16px;*/
/*  border-radius: 15px 15px 0 0;*/
/*  border-bottom: 1px solid #1E42BB;*/
/*}*/

/*.sub-menu li {*/
/*  cursor: pointer;*/
/*  padding: 7px 0 7px 45px;*/
/*}*/

.sub-menu-item a {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 10px 0;
  text-decoration: none;
  font-style: normal;
  font-weight: 400;
  font-size: 18px;
  line-height: 21px;
  color: var(--light);
}

.sub-menu-item:last-child {
  border-bottom-right-radius: 10px;
  border-bottom-left-radius: 10px;
}


.createNew > li:first-child {
  position: relative;
}
</style>