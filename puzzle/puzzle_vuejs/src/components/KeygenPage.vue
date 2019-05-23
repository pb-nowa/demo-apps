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
      color: white;
      text-decoration: underline;
    }
  }
</style>

<template>
  <div class="keygen-page">
    <div class="container">
      <div class="row">

        <div class="new-logo"></div>
      </div>
      <div class="row">
        <div class="inner">
          <div class="input-div input-token">
            <label for="" class="white-color">Your key has generated</label>
            <br>
            <input class="harmony-input"
                   v-model="userKey"
                   type="text">
          </div>

          <div class="submit-section">
            <button class="btn-primary btn-harmony"
                    @click="$emit('join', 5)">Start game</button>
          </div>

        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "KeygenPage",

    props: ['userKey'],

    data() {
      return {};
    },
    mounted: function() {
      console.log("eeeeeeee userKey", this.userKey);
    },
    methods: {
      submit() {
        // submit to backend

        // and change page ( generate token )
        this.$emit('join', 13);
      }
    }
  };
</script>
