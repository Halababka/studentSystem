<template>
  <div id='home-widget'>
    <div class="top">
      <h1>{{ localeHours }}, {{ userData.first_name }}</h1>
      <img src="@/assets/avatar.jpg" alt="Avatar"/>
    </div>
    <p>У вас {{ CountLesson }}</p>
    <p>Контрольных точек не запланировано</p>
  </div>
</template>

<script setup>
import {computed} from "vue";
import TimetableWidget from "@/components/Widgets/TimetableWidget.vue";
import {userData} from "@/composables/getUserData";

const CountLesson = computed(() => {
  const valueCountLesson = TimetableWidget.methods.getCountLessons();
  if (valueCountLesson) {
    let lessonDeclination = (valueCountLesson > 4) ? "пар" : "пары";
    return valueCountLesson + " " + lessonDeclination;
  } else {
    return "нету пар";
  }
});

const localeHours = computed(() => {
  let localeHours = new Date().getHours();
  if (localeHours > 3 && localeHours < 12) return "Доброе утро";
  else if (localeHours > 11 && localeHours < 19) return "Добрый день";
  else if (localeHours > 18 && localeHours < 24) return "Добрый вечер";
  else if (localeHours > 23 || localeHours < 4) return "Доброй ночи";
  throw Error("localHours error");
});

</script>

<style scoped>
.lds-ripple {
  align-self: center;
}


#home-widget {
  min-height: 220px;
  flex-direction: column;
  justify-content: flex-start;
  gap: 15px;
}

#home-widget .top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 150px;
}

.top img {
  width: 120px;
  border-radius: 15px;
}

#home-widget p, #home-widget h1 {
  font-size: 24px;
  margin: 0;
  z-index: 1;
}
</style>
