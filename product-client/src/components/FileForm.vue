<template>
  <form class="image-form" @submit.prevent="submit">
    <div class="image-form-title">
      {{ title }}
    </div>
    <div class="image-form-input mt-4">
      <input type="file" name="file" ref="file" />
      <input class="btn-green" type="submit" value="Upload File" />
    </div>
  </form>
</template>

<script>
export default {
  name: "FileForm",
  props: {
    title: {
      type: String,
      default: "Upload File",
    },
  },
  methods: {
    submit() {
      const formData = new FormData();
      const file = this.$refs["file"].files[0];
      formData.append("file", file);

      this.$emit("submit", formData);
    },
  },
};
</script>

<style scoped lang="scss">
$grey: rgb(222, 222, 222);
$green: #42b983;

.mt-4 {
  margin-top: 16px;
}

.image-form {
  width: 400px;
  margin: 0 auto;
  padding: 16px;
  border: 1px solid $grey;
  border-radius: 4px;

  display: flex;
  flex-direction: column;

  div.image-form-title {
    width: 100%;
  }

  div.image-form-input {
    display: flex;
    justify-content: space-between;
  }
}

.btn-green {
  background-color: $green;
  color: white;
  border: 1px solid $green;
  border-radius: 4px;
}
</style>
