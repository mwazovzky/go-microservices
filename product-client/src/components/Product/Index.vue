<template>
  <div class="product-index">
    <h1>Menu</h1>
    <ProductCard class="mt-1" v-for="product in products" :key="product.id" :product="product" />
  </div>
</template>

<script>
import ProductCard from "./Card";

export default {
  name: "ProductIndex",
  components: {
    ProductCard,
  },
  props: {
    products: {
      type: Array,
      default: () => [],
    },
  },
};
</script>

<style scoped lang="scss">
.mt-1 {
  margin-top: 4px;
}
</style>
