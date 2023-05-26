import {createApp} from "vue";
import App from "./App.vue";
import router from "./router";

const cors = require("cors");
createApp(App).use(router,cors()).mount("#app");

