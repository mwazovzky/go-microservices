<template>
  <div class="menu">
    <ProductIndex :products="products" />
  </div>
</template>

<script>
// @ is an alias to /src
import ProductIndex from "@/components/Product/Index.vue";

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
    this.$http.get().then(res => {
      this.products = res.data;
    });
  },
};
</script>
