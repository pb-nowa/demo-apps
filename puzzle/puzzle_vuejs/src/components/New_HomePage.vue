<style scoped lang="less">
.content {
  margin: 3em;
}

.btn-primary {
  display: block;
  margin: auto;
}

.new-home-page, .container, .row, .new-home-logo {
  height: 100%;
  overflow: hidden;
}

.new-home-logo {
  display: flex;
  width: 100%;
  justify-content: center;
  align-items: center;
  cursor: pointer;

  img {
    width: 100%;
  }
}

// Extra large devices (large desktops, 1200px and up)
@media (min-width: 1200px) {
  .new-home-logo {
    img {
      width: 100%;
    }
  }
}

// Large devices (desktops, 992px and up)
@media (min-width: 992px) {
  .new-home-logo {
    img {
      width: 747px;
    }
  }
}

// Medium devices (tablets, 768px and up)
@media (min-width: 768px) {
  .new-home-logo {
    img {
      width: 600px;
    }
  }
}

@media (max-width: 575.98px) {
  .new-home-page {
    height: 100%;
    display: flex;
    align-items: center;
  }
}

</style>

<template >
  <div class="new-home-page">
    <div class="container">
      <div class="row">
        <div v-on:click="toWelcome" class="new-home-logo">
          <img src="../assets/new/puzzle-logo.png" alt="">
        </div>
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
