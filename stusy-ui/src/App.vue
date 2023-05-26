<template>
  <component :is="layout">
    <router-view/>
  </component>

</template>

<script>
import HomeLayout from "@/layouts/HomeLayout";
import EmptyLayout from "@/layouts/EmptyLayout";

export default {
  computed: {
    layout() {
      return this.$route.meta.layout
    }
  },
  components: {
    HomeLayout, EmptyLayout
  }
}
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Ubuntu:ital,wght@0,300;0,400;0,500;0,700;1,300;1,400;1,500;1,700&display=swap");

* {
  margin: 0;
  padding: 0;
}

body {
  display: block;
  font-family: "Ubuntu", sans-serif;
  background: var(--dark-mate-mate);

  --dark: #121212;
  --dark-mate: #292a2d;
  --dark-mate-mate: #dbdee8;

  --light: #fff;
  --light-mate: #dcfffd;

  --blue: #2255f4;
  --blue-mate: #1e42bb;

  overflow: hidden;

  margin: 0 auto;
}

.input-form {
  margin: 10px 0;
  text-align: left;
}

input {
  font-style: normal;
  font-weight: 500;
  font-size: 18px;
  line-height: 21px;
}

textarea {
  resize: none;
  height: 10rem;
  overflow: auto;
  max-width: 950px;
  padding: 20px;
  font-weight: 300;
  font-size: 24px;
  border: 1px solid var(--blue);
  border-radius: 10px;
}

textarea:focus {
  outline: none !important;
  border: 1px solid var(--blue-mate);
}

.input-form input {
  margin: 10px 0;
  padding: 25px 25px;
  width: 100%;
  background: rgba(69, 70, 79, 0.1);
  font-size: 18px;
  border: none;
  border-radius: 15px;
  box-sizing: border-box;
  outline: none;
}

.main {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  justify-content: flex-start;
  gap: 29px;
  width: 100%;
  padding: 15px 24px 24px;
  height: 100vh;
  overflow: auto;
}

.main header {
  display: flex;
  align-items: center;
  justify-content: flex-end;
}

.mainContent {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  gap: 16px;
  margin-bottom: 30px;
}

.courses {
  display: flex !important;
  background: var(--dark-mate-mate);
  min-height: 100vh;
}

.coursesMainLeft {
  display: flex;
  flex: 3;
  flex-direction: column;
  justify-content: flex-start;
  gap: 21px;
  padding: 21px 26px;
  max-width: 70%;
  min-height: 80%;
  background: var(--light);
  border-radius: 25px;
}

.coursesMainAbout {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 14px;
  max-width: 22%;
  min-height: 70%;
  padding: 21px 20px;
  background: var(--light);
  border-radius: 25px;
}

.coursesMainAbout h1 {
  font-style: normal;
  font-weight: 500;
  font-size: 24px;
  line-height: 28px;
}

.button-login input[type='submit'] {
  color: #ffffff;
  background: var(--blue);
  border: none;
  border-radius: 15px;
  font-size: 18px;
  box-sizing: border-box;
  width: 100%;
  padding: 25px;
  transition: 0.3s;
  cursor: pointer;
}

.button-login input[type='submit']:hover {
  background: var(--blue-mate);
}

.button-save {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  align-content: center;
  padding: 17px 15px;
  gap: 13px;
  background: var(--blue);
  border-radius: 15px;
  cursor: pointer;
}

.button-save input {
  cursor: pointer;
  color: #ffffff;
  border: none;
  background: none;
}

.button-save:hover, .sub-menu li:hover {
  background: var(--blue-mate);
}

.redactorContentCreating-topRow {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.redactorContentCreating-headerOfSection {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: flex-start;
  gap: 26px;
}

.groupForm-row {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 21px;
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
}

.groupForm-row .input-form-allWidth {
  flex-grow: 1;
}


.accordionMenu-list {
  display: flex;
  flex-direction: column;
  gap: 30px;
  list-style-type: none;
}

.radio-button {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: flex-start;
  gap: 20px;
}

.modal-window {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: stretch;
  width: 30%;
  gap: 17px;
  background-color: var(--light);
  border-radius: 15px;
  padding: 30px 20px;
}

.modal-window h1 {
  margin-bottom: 10px;
  font-style: normal;
  font-weight: 400;
  font-size: 36px;
  line-height: 41px;
}

.modal-window label {
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
}

.close-button {
  position: absolute;
  width: 20px;
  height: 20px;
  top: 10px;
  right: 10px;
  cursor: pointer;
}

.block-courses {
  display: flex;
  width: 100%;
  justify-content: space-evenly;
  gap: 5px;
}

.course-card {
  border: 1px solid var(--blue);
  border-radius: 25px;
  padding: 21px 10px;
  display: flex;
  flex-direction: column;
  text-align: center;
  height: 160px;
  width: 195px;
  user-select: none;
  cursor: pointer;
}

.course-card img {
  display: block;
  margin: auto;
  width: 110px;
}

@media (min-width: 1920px) {
  #menuVertical {
    display: block;
    text-align: center;
    padding: 39px 24px;
    margin: 27px 0 0 27px;
    width: 300px;
    height: fit-content;
    box-sizing: border-box;
    border-radius: 25px !important;
    background: var(--light);
  }
}

.skeleton {
  position: relative;
  cursor: default;
}

.skeleton::after {
  background-image: linear-gradient(
      90deg,
      #EEEEEE 0%, #EEEEEE 40%,
      #DDDDDD 50%, #DDDDDD 55%,
      #EEEEEE 65%, #EEEEEE 100%
  );
  background-size: 400%;
  animation: skeleton 1500ms infinite;
  content: '';
  z-index: 2;
  width: 100%;
  height: 100%;
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 5px;
}

@keyframes skeleton {
  from {
    background-position: 100% 100%;
  }
  to {
    background-position: 0 0;
  }
}


</style>
