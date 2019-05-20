<style scoped lang="less">
.content {
  margin: 3em;
}
.intro-wrapper {
  position: relative;

  .intro {
    position: relative;
    z-index: 10;
    border-radius: 0.5em;
    // border: 0.15em solid #979797;
    padding: 1em;
    background-color: #fff;
    height: 10em;
    overflow: auto;
  }
}

.btn-primary {
  display: block;
  margin: auto;
}

.new-home-logo {
  position: absolute;
  width: 747px;
  height: 524px;
  // left: 347px;
  top: 100px;
  background: url(../assets/new/puzzle-logo.png);
}

</style>

<template >
  <div class="new-home-page">
    <div class="content">
      <div class="new-home-logo"></div>
    </div>
    <!-- <button class="btn-primary" @click="$emit('join')">Start Game</button> -->
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
  methods: {}
};
</script>
