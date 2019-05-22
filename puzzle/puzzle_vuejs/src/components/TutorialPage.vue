<style scoped lang="less">
  .content {
    margin: 3em;
  }

  .inner {
    width: 800px;
    margin: 0 auto;
  }

  .top {
    display: flex;
    justify-content: center;
  }

  img {
    width: 60px;
  }

  .tutorial-content {
    background: white;
    padding: 1.5rem 2rem;
    border-radius: 8px;
    border: 2px solid #13156A;
    box-shadow:
      inset 0px 0px 0px 0px transparent,
      inset 0px -10px 0px 0px #7C90E6;
    margin: 0 1.5rem;
    height: 300px;
    width: 500px;
    text-align: center;
  }

  .left-button, .right-button {
    display: flex;
    align-items: center;
  }

  .submit-button {
    text-align: center;
    margin-top: 2rem;
  }
</style>

<template>
  <div class="tutorial-page">
    <div class="new-logo"></div>
    <div class="content">
      <div class="inner">
        <div class="top">
          <div class="left-button">
            <img src="../assets/left_button.png" alt="">
          </div>
          <div class="tutorial-content">
            Tutorial content
          </div>
          <div class="right-button">
            <img src="../assets/right_button.png" alt="">
          </div>
        </div>

        <div class="submit-button">
          <button class="btn-primary btn-harmony join-now" @click="$emit('join', 5)">Got it</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "TutorialPage",

    props: {},
    data() {
      return {};
    },
    mounted: function() {},
    methods: {
      backToEnterPage() {
        console.log("eeeeeeee ")
        this.$emit('join', 12)
      }
    }
  };
</script>
