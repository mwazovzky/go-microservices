<template>
  <div class="product-item">
    <h1>
      {{ product.name }}
    </h1>
    <ProductImages :product="product" />
  </div>
</template>

<script>
import ProductImages from "./Images";

export default {
  name: "ProductItem",
  components: {
    ProductImages,
  },
  props: {
    product: {
      type: Object,
      required: true,
    },
  },
};
</script>

<style scoped lang="scss">
//
</style>
