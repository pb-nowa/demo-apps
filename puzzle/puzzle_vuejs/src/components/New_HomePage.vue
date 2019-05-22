<style scoped lang="less">
.content {
  margin: 3em;
}

.btn-primary {
  display: block;
  margin: auto;
}

.new-home-logo {
  width: 747px;
  height: auto;
  margin: 0 auto;
  margin-top: 200px;
  cursor: pointer;
}
</style>

<template >
  <div class="new-home-page">
    <div class="content">
      <div v-on:click="toWelcome" class="new-home-logo">
        <img src="../assets/new/puzzle-logo.png" alt="">
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "NewHomePage",

  props: {},
  data() {
    return {};
  },
  mounted: function() {},
  methods: {
    toWelcome: function(event) {
      this.$emit('join', 11);
    }
  }
};
</script>
