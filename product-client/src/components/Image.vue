<template>
  <img class="image" :src="src" :alt="alt" />
</template>

<script>
export default {
  name: "Image",
  props: {
    src: {
      type: String,
      required: true,
    },
    alt: {
      type: String,
      required: "Some Image",
    },
  },
};
</script>

<style scoped lang="scss">
$grey: rgb(222, 222, 222);

.image {
  height: 200px;
  border: 1px solid $grey;
  border-radius: 8px;
}
</style>
