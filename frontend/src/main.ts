import { createApp } from "vue";
import { createPinia } from "pinia";
import 'mdui/mdui.css';
import 'mdui';
import App from "./App.vue";
import router from "./router";
import i18n from "./i18n";

import "./style.scss";
const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(i18n);

app.mount("#app");
