<template>
  <div class="homepage">
    <AsideMenu/>
    <div class="home">
      <div class="content">
        <header>
          <TopMenu></TopMenu>
        </header>
        <router-view/>
      </div>
    </div>
  </div>
</template>

<script setup>
import AsideMenu from "@/components/AsideMenu";
import TopMenu from "@/components/TopMenu";
import {onBeforeMount} from "vue";
import {getUserData} from "@/composables/getUserData";

onBeforeMount(() => {
  getUserData();
});

</script>

<style scoped>
.homepage {
  display: flex !important;
  max-width: 1920px;
  margin: 0 auto;
}

.home {
  flex-grow: 1;
  background: var(--dark-mate-mate);
}

.content header {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  justify-content: flex-start;
  margin: 12px 0;
}

.content header h1 {
  font-weight: 500;
  font-size: 36px;
  margin: 0;
}

.content {
  padding: 15px 24px 24px;
  height: 100vh;
  overflow: auto;
}

@media (max-width: 700px) {
  #menuVertical {
    display: none;
  }

  .content header {
    width: 347px;
    margin: 0 auto;
  }
}

</style>