<template>
  <div class="menu">
    <ProductItem v-if="product" :product="product" />
  </div>
</template>

<script>
import ProductItem from "@/components/Product/Item.vue";

export default {
  name: "Product",
  components: {
    ProductItem,
  },
  data() {
    return {
      product: null,
    };
  },
  created() {
    const id = this.$route.params.id;
    this.$http.get(id).then(res => {
      this.product = res.data;
    });
  },
};
</script>
