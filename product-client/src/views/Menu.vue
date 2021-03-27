<template>
  <div class="menu">
    <ProductIndex :products="products" />
  </div>
</template>

<script>
// @ is an alias to /src
import axios from "axios";
import ProductIndex from "@/components/Product/Index.vue";

const BASE_URL = process.env.VUE_APP_PRODUCT_API_URL;

export default {
  name: "Menu",
  components: {
    ProductIndex,
  },
  data() {
    return {
      products: [],
    };
  },
  created() {
    axios.get(BASE_URL).then(res => {
      this.products = res.data;
    });
  },
};
</script>
