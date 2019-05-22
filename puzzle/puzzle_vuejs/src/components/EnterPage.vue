<style scoped lang="less">
  .content {
    margin: 3em;
  }

  .inner {
    width: 500px;
    margin: 0 auto;
  }

  input {
    width: 100%;
  }

  .small small {
    float: right;
  }

  .input-token {
    padding-bottom: 30px;
  }

  .input-email {
    margin-bottom: 40px;
  }

  .submit-section {
    text-align: center;
    margin-bottom: 20px;
  }

  .view-tutorial {
    text-align: center;
    a {
      color: white
    }
  }
</style>

<template>
  <div class="enter-page">
    <div class="new-logo"></div>
    <div class="content">
      <div class="inner">
        <div class="input-div input-token">
          <label for="" class="white-color">You have 100 Harmony tokens. How much you want to stake?</label>
          <br>
          <input class="harmony-input" placeholder="100" type="number">
          <br>
          <div class="small">
            <small class="white-color">(Stake 2x, 3x increases reward)</small>
          </div>
        </div>

        <div class="input-div input-email">
          <input class="harmony-input" placeholder="user@email.com" type="email">
        </div>

        <div class="submit-section">
          <button class="btn-primary btn-harmony" @click="$emit('join')">Submit</button>
        </div>

        <div class="view-tutorial">
          <a v-on:click="$emit('join', 4)" href="#">View Tutorial</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "EnterPage",

    props: {},
    data() {
      return {};
    },
    mounted: function() {},
    methods: {}
  };
</script>
