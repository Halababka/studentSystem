<template>
  <div id="courses">
    <div class="blockHeader">
      <h1>Ваши курсы</h1>
      <router-link to="/mycourses">Смотреть все</router-link>
    </div>
    <div class="block-courses">
      <div class="course-card">
        <img src="@/assets/data_archiving_and_compression.svg">
        <p>Архивация и сжатие данных</p>
      </div>
      <div class="course-card">
        <img src="@/assets/botany_and_plant_physiology.svg">
        <p>Ботаника и физиология растений</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "CoursesWidget"
}
</script>

<style scoped>

</style>