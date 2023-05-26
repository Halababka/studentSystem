<template>
  <div class="mainContent">
    <div class="coursesMainLeft" id="coursesMain">
      <h1>Архивация и сжатие данных</h1>
      <section class="advertisement">
        <h1>Объявление</h1>
        <article>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore
          et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip
          ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu
          fugiat nulla pariatur.
        </article>
      </section>
      <ol class="accordionMenu-list">
        <li><a href="">Введение</a></li>
        <AccordionMenu title="Модуль 1" class="accordionMenu">

        </AccordionMenu>
        <AccordionMenu title="Модуль 2" class="accordionMenu">
          <ol>
            <li><a href="#">Как сжать звуковые файлы</a></li>
            <li><a href="#">Как сжать графические файлы</a></li>
            <li><a href="#">Какими возможностями обладают архиваторы</a></li>
            <li><a href="#">&#10071; Контрольная точка 2 - Тестирование</a></li>
          </ol>
        </AccordionMenu>
        <AccordionMenu title="Модуль 3" class="accordionMenu">

        </AccordionMenu>
      </ol>
    </div>
    <div class="coursesMainAbout">
      <h1>О курсе</h1>
      <div class="coursesAboutInfo">
        <p>3 зачетных единицы</p>
        <p>5 контрольных точек</p>
      </div>
    </div>
  </div>
</template>

<script>
import AccordionMenu from "@/components/AccordionMenu";

export default {
  name: "MyCourses",
  components: {AccordionMenu},
  data() {
    return {
      menu: false
    };
  },
  methods: {
    menuAction() {
      this.menu = !this.menu;
    }
  }
};
</script>

<style scoped>
h1 {
  font-style: normal;
  font-weight: 400;
  font-size: 36px;
  line-height: 41px;
}

p {
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
}

a {
  text-decoration: none;
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
  color: var(--blue);
}

ol {
  list-style-type: none;
}

#coursesMain {
  display: flex;
  flex: 3;
  flex-direction: column;
  gap: 28px;
  padding: 27px 32px;
  max-width: 70%;
  min-height: 80%;
  background: var(--light);
  border-radius: 25px;
}

.coursesAboutInfo {
  margin-top: 16px;
  display: flex;
  flex-direction: column;
  gap: 11px;
}

.advertisement {
  display: flex;
  flex-direction: column;
  padding: 17px 19px;
  border: 3px solid var(--blue);
  border-radius: 25px;
  margin-bottom: 30px;
}

.advertisement h1 {
  font-style: normal;
  font-weight: 500;
  font-size: 24px;
  line-height: 28px;
}

.advertisement article {
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 160%;
  text-align: justify;
}

.accordionMenu ol {
  display: flex;
  flex-direction: column;
  gap: 30px;
  margin-top: 30px;
}
</style>