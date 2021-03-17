import axios from "axios";

const http = axios.create({
  baseURL: process.env.VUE_APP_PRODUCT_API_URL,
});

export default {
  install: (app) => {
    app.config.globalProperties.$http = http;
  },
};
