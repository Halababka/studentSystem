<template>
  <header>
    <ul>
      <li>
        <h2 class="logo">Stusy</h2>
      </li>
      <li class="header-buttons">
        <router-link to="/auth">Вход</router-link>
        <button>Демо</button>
      </li>
    </ul>
  </header>
  <main>
    <h1>Улучшайте учебный процесс</h1>
    <h2>Вместе с Stust</h2>
    <button>Попробовать демо</button>
    <div class="cards">
      <div class="card">
        <h3>Для студентов</h3>
        <ul>
          <li>Помощь с организацией</li>
          <li>Оценка знаний</li>
          <li>Персонализированные материалы</li>
        </ul>
      </div>
      <div class="card">
        <h3>Для абитуриентов</h3>
        <ul>
          <li>Оценка знаний</li>
          <li>Персонализированные материалы</li>
        </ul>
      </div>
      <div class="card">
        <h3>Для преподавателей</h3>
        <ul>
          <li>Оценка качества обучения</li>
        </ul>
      </div>
      <div class="card">
        <h3>Для преподавателей</h3>
        <ul>
          <li>Оценка и мониторинг качества знаний</li>
        </ul>
      </div>

    </div>
  </main>

  <!--  <a href="/auth">Вход</a>-->
</template>

<script>
export default {
  name: "MainView"
}
</script>

<style scoped>

button {
  color: #ffffff;
  background: var(--blue);
  border: none;
  border-radius: 15px;
  font-size: 18px;
  box-sizing: border-box;
  width: 100%;
  padding: 15px 25px;
  margin: 10px;
  transition: 0.3s;
  cursor: pointer;

}

button:hover {
  background: var(--blue-mate);
}

header {
  max-width: 1920px;
  background: var(--light);
  border-bottom-left-radius: 15px;
  border-bottom-right-radius: 15px;
}

header ul {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin: 0 auto;
  list-style: none;
}

.logo {
  color: var(--blue);
  user-select: none;
  cursor: pointer;
}

.header-buttons {
  display: flex;
  align-items: center;
}

.header-buttons a {
  color: var(--blue);
  margin-right: 15px;
  font-size: 18px;
  text-decoration: var(--blue);
}

ul {
  max-width: 1042px;
}

main {
  display: flex;
  flex-direction: column;
  gap: 35px;
  align-items: center;
}

main h1 {
  font-size: 46pt;
  margin-top: 35px;
}

main h2 {
  font-size: 36pt;
  color: var(--blue);
}

main button {
  width: auto;
}

.cards {
  display: flex;
  flex-wrap: wrap;
  max-width: 1000px;
  margin: 0 auto;
  gap: 25px;
  justify-content: center;
}

.card {
  background: var(--light);
  padding: 15px;
  border-radius: 15px;
  height: 106px;
  width: 300px;
}

.card ul {
  list-style: none;
}

.card li {
  margin-top: 10px;
}
</style>