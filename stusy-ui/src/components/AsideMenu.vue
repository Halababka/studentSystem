<template>
  <aside id="menuVertical">
    <h1>
      <router-link to="/home" v-on:click="[status='home']">Stusy</router-link>
    </h1>
    <ol>
      <li v-for="link in links" v-bind:key="link.id">
        <router-link v-bind:to="link.to" active-class="active">
          <div class="LeftMenu_icon" v-bind:id="link.iconId"> 
          </div>
          <span>{{ link.text }}</span>
        </router-link>
      </li>
    </ol>
  </aside>
</template>

<script setup>
const links = [
  {
    to: '/home',
    iconId: 'home-ico',
    text: 'Домашняя страница',
  },
  {
    to: '/mycourses',
    iconId: 'book-ico',
    text: 'Мои курсы',
  },
  {
    to: '#',
    iconId: 'test-ico',
    text: 'Меню для тестов',
  },
  {
    to: '/courses',
    iconId: 'test-ico',
    text: 'Курсы',
  },
  {
    to: '/creatingCourses',
    iconId: 'test-ico',
    text: 'Создание курса',
  },
  {
    to: '/creatingTest',
    iconId: 'test-ico',
    text: 'Создание теста',
  },
  {
    to: '/asdasd',
    iconId: 'test-ico',
    text: 'Не найдено',
  },
];
</script>

<style scoped>
#menuVertical {
  display: block;
  text-align: center;
  padding: 39px 24px;
  width: 300px;
  min-height: 100%;
  box-sizing: border-box;
  border-radius: 0 25px 0 0;
  background: var(--light);
}


#menuVertical > h1 {
  margin-bottom: 30px;
}

#menuVertical > h1 > a {
  font-weight: 500;
  font-size: 28px;
  color: var(--blue);
  margin-top: 0;
  text-decoration: none;
}

#menuVertical ol {
  display: block;
  margin: 0;
  padding: 0;
  list-style: none;
  position: relative;
}

#menuVertical ol li {
  display: list-item;
  box-sizing: border-box;
  width: 252px;
  height: 54px;
}

.active {
  background: var(--blue);
}

.active .LeftMenu_icon {
  filter: invert();
}

.active > span {
  color: var(--light)
}


#menuVertical ol li a {
  display: flex;
  color: var(--dark);
  text-decoration: none;
  font-size: 18px;
  padding: 16px 21px;
  border-radius: 15px;
  margin-bottom: 15px;
}

.LeftMenu_icon {
  margin: 0 13px 0 0;
}

#home-ico {
  background: url("@/assets/home.svg");
  width: 21px;
}
#book-ico {
  background: url("@/assets/book.svg");
  width: 19px;
}
#calendar-ico{
  background: url("@/assets/calendar.svg");
  width: 19px;
}
#test-ico{
  background: url("@/assets/test.svg");
  width: 21px;
}
</style>