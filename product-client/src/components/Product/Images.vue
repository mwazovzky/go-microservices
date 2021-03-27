<template>
  <div class="product-item">
    <div class="image-index">
      <Image class="m-1" v-for="image in images" :key="image" :src="image" />
    </div>
    <FileForm class="mt-4" @submit="upload" />
  </div>
</template>

<script>
import axios from "axios";
import Image from "@/components/Image";
import FileForm from "@/components/FileForm";

const BASE_URL = process.env.VUE_APP_IMAGES_API_URL;

export default {
  name: "ProductImages",
  components: {
    Image,
    FileForm,
  },
  props: {
    product: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      files: [],
    };
  },
  computed: {
    url() {
      return `${BASE_URL}/${this.product.id}`;
    },
    images() {
      return this.files.map(file => `${this.url}/${file}`);
    },
  },
  created() {
    this.load();
  },
  methods: {
    load() {
      axios.get(this.url).then(res => {
        this.files = res.data;
      });
    },
    upload(formData) {
      formData.append("id", this.product.id);
      axios.post(BASE_URL, formData).then(() => {
        this.load();
      });
    },
  },
};
</script>

<style scoped lang="scss">
.m-1 {
  margin: 4px;
}

.mt-4 {
  margin-top: 16px;
}

.image-index {
  display: flex;
  flex-wrap: wrap;
}
</style>
