<template>
  <div id="login">
    <div class="left-content"></div>
    <div class="right-content">
      <transition name="fade">
        <ModalWindow v-if="loading">
          <ScreenLoader/>
        </ModalWindow>
      </transition>
      <h1 class="title">StudentSystem</h1>
      <form class="form" v-on:submit.prevent="auth">
        <h1>Вход</h1>
        <div
            class="error-block"
            v-bind:style="[error ? { opacity: 1 } : { opacity: 0 }]"
        >
          <p>{{ errorText }}</p>
        </div>
        <div class="input-form">
          <h2>Электронная почта</h2>
          <input
              type="text"
              placeholder="name@gmail.com"
              v-model="email"
              required
          />
        </div>
        <div class="input-form">
          <h2 style="float: left">Пароль</h2>
          <a href="#" class="forget" style="float: right">Забыли пароль?</a>
          <input
              type="password"
              placeholder="Введите пароль"
              v-model="password"
              required
          />
        </div>
        <div class="button-login">
          <input type="submit" value="Войти"/>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import {getCookie, url} from "@/global";
import ScreenLoader from "@/components/ScreenLoader";
import ModalWindow from "@/components/ModalWindow";

export default {
  name: "AuthPage",
  components: {ScreenLoader, ModalWindow},

  data() {
    return {
      email: "",
      password: "",
      error: false,
      loading: false,
      errorText: "Сервер не отвечает. Попробуйте повторить запрос позже."
    };
  },
  methods: {
    auth() {
      this.loading = true;
      fetch(`${url}/user/auth`, {
        method: "POST",
        body: JSON.stringify({
          email: this.email,
          password: this.password
        })
      }).then((response) => {
        return response.json();
      })
          .then((data) => {
            if (data["error_message"]) {
              this.error = true;
              this.errorText = data["error_message"];
            }
            this.loading = false;
            if (data["access_token"]) {
              this.error = false;
              document.cookie = "TOKEN=" + data["access_token"];
              document.cookie = "ID=" + data["uid"];
              this.$router.push("/home");
            }
          })
          .catch((err) => {
            console.error("Невозможно отправить запрос", err);
            this.loading = false;
            this.error = true;
            this.errorText =
                "Сервер не отвечает. Попробуйте повторить запрос позже.";
          });
    }
  },
  mounted() {
    if (getCookie("TOKEN") && getCookie("ID") !== undefined) {
      this.$router.push("/home");
    }
  }
};
</script>

<style scoped>
#login {
  background: var(--dark-mate);
  display: flex;
}

.left-content {
  width: 60%;
}

.right-content {
  position: relative;
  margin-left: auto;
  padding: 10px;
  width: 40%;
  background: var(--light);
  display: flex;
  align-items: center;
  /*overflow: hidden;*/
  flex-direction: column;
  justify-content: space-around;
}

h1 {
  text-align: left;
  margin: 39px 57px;
  font-size: 36px;
  color: var(--blue);
}

.form {
  max-width: 460px;
  display: flex;
  flex-direction: column;
}

.form h1 {
  margin: 0;
  padding: 0;
  text-align: center;
  color: #000000;
  font-size: 36px;
}

.form h2 {
  margin: 0;
  padding: 0;
  text-align: left;
  color: #000000;
  font-size: 18px;
}

.input-form a {
  display: block;
  text-align: right;
  color: var(--blue);
  font-size: 18px;
  text-decoration: none;
  transition: 0.3s;
}

.input-form a:hover {
  color: var(--blue-mate);
}

h1.title {
  top: 10px;
}

.error-block {
  border-left: 3px solid red;
  box-shadow: 1px 1px 14px 1px #0000003b;
  border-top-right-radius: 5px;
  border-bottom-right-radius: 5px;
  margin: 20px 0;
  transition: 0.3s;
  opacity: 0;
}

.error-block p {
  color: red;
  padding: 16px;
}

@media (max-width: 1350px) {
  .left-content {
    display: none;
  }

  .right-content {
    margin-left: 0;
    width: 100%;
    padding: 0;
  }
}

@media (max-width: 480px) {
  .form {
    padding: 15px;
  }

  h1 {
    font-size: 11vw;
  }

  .form h1 {
    font-size: 9vw;
  }
}

@media (min-height: 620px) {
  .right-content {
    height: 100vh;
    justify-content: flex-start
  }
}
</style>
