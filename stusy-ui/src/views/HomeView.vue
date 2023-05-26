<template>
  <div class="mainContent">
    <HomeWidget class="block TwoTwo"/>
    <TimetableWidget class="TwoFor" v-if="checkCountLesson()"/>
    <CoursesWidget class="block TwoTwo"/>
    <!--    <WeatherWidget class="block TwoTwo"/>-->

  </div>
</template>
<script setup>
import {getCookie} from "@/global";
import HomeWidget from "@/components/Widgets/HomeWidget.vue";
import TimetableWidget from "@/components/Widgets/TimetableWidget.vue";
// import WeatherWidget from "@/components/Widgets/WeatherWidget.vue";
import CoursesWidget from "@/components/Widgets/CoursesWidget.vue";
import {onMounted} from "vue";

function checkCountLesson() {
  let value = TimetableWidget.methods.getCountLessons()
  return !!value;
}

onMounted(() => {
  if (!getCookie("TOKEN")) {
    this.$router.push("/auth");
  }
})

</script>

<style>

.block {
  position: relative;
  overflow: hidden;
  background: var(--light);
  border-radius: 25px;
  display: flex;
  flex-wrap: wrap;
  align-content: space-around;
  justify-content: space-between;
  padding: 15px;
  margin: 5px;

  transition: .2s;
  transition-timing-function: cubic-bezier(0.08, 0.55, 0.82, 0.41)
}

.block:hover {
  transform: scale(1.018);
}

.blockHeader {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.block > .block {
  padding: 0;
}

.OneOne {
  height: 145px;
  width: 145px;
  background: #000;
}

.TwoTwo {
  height: 310px;
  width: 310px;
}

.TwoFor {
  height: 310px;
  max-height: 310px;
  width: 680px;
  max-width: 680px;
}

.slide-fade-leave-active {
  transition: all .3s ease;
}

.slide-fade-leave-to {
  transform: translateY(-10px);
  opacity: 0;
}

.mainContent {
  display: flex;
  width: 100%;
  flex-wrap: wrap;
  align-items: flex-start;
  justify-content: center;
  align-content: flex-start;
  max-width: 1100px;
  margin: 0 auto;
}

.mainContent h1 {
  margin: 0;
  font-style: normal;
  font-weight: 500;
  font-size: 24px;
  float: left;
}

#courses div a {
  float: right;
  font-style: normal;
  font-weight: 400;
  font-size: 18px;
  color: var(--blue);
  text-decoration: none;
  text-align: right;
}

@media (max-width: 1430px) {
  .TwoFor {
    height: 310px;
    width: 310px;
  }
}
</style>
