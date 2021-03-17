<template>
  <div class="product-card">
    <div class="fb-20">{{ product.name }}</div>
    <div class="fb-20">{{ product.sku }}</div>
    <div class="flex-grow">{{ product.description }}</div>
    <div class="w-50">{{ product.price }}</div>
  </div>
</template>

<script>
export default {
  name: "ProductCard",
  props: {
    product: {
      type: Object,
      required: true,
    },
  },
};
</script>

<style scoped lang="scss">
$grey: rgb(222, 222, 222);

.product-card {
  padding: 8px;
  border: 1px solid $grey;
  border-radius: 4px;
  display: flex;
  justify-content: space-between;
  div {
    text-align: left;
  }
}
.w-50 {
  width: 50px;
}
.fb-20 {
  flex-basis: 20%;
}
.flex-grow {
  flex-grow: 1;
}
</style>
