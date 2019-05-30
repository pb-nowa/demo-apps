<style scoped lang="less">
.page-content {
  margin: 3em;
}

.logo {
  font-size: 2em;
}

@media (min-width: 800px) {
  .logo {
    font-size: 2.5em;
  }
}
.msg {
  padding: 0.5em;
  color: #59504d;
  font-family: Fira Sans, sans-serif;
  font-size: 0.8em;
}
.value {
  padding: 1em;
  border-radius: 0.5em;
  color: #59504d;
  font-family: Fira Sans, sans-serif;
  // border: 0.15em solid #979797;
  overflow: auto;
  background-color: #fff;
  margin: 0 auto;
  text-align: center;
}
footer {
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  padding: 2em 1em;
  .link {
    font-size: 0.8em;
    text-align: center;
  }
}
.host {
  max-width: 600px;
  margin: 0 auto;
}
.info-item > .content {
  font-size: 1.4em;
}
.btn-mini.start-btn {
  width: auto;
  padding: 0 1em;
}
.stake-amount {
  margin: 0 0.5em;
  border-radius: 0.3em;
  border: 0;
  color: #1b2a5d;
  height: 2em;
  width: 4em;
  font-weight: bold;
}
.stake-row {
  margin: 2em auto 0;
  padding: 0 1rem;
}
.icon-dark-token {
  background-size: contain;
  height: 1.4em;
  width: 1.2em;
  background-image: url(../assets/dark-token.svg);
  background-repeat: no-repeat;
  margin-right: 0.5em;
}
.stake-buttons {
  background-color: #fff;
  border-radius: 0.3em;
}
.btn-mini {
  font-size: 1em;
  background-color: transparent;
  border: 0;
  color: #482aff;
  outline: none;
  &:disabled {
    opacity: 0.5;
    color: #ddd;
  }
}

.btn-primary {
  font-size: 1em;
  background-color: #482bff;
}

.plus-minus-section {
  display: flex;
  position: relative;
  width: 100%;

  .plus-button, .minus-button {
    width: 74px;
    height: 74px;
    padding: 0.5rem !important;
    position: absolute;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .plus-button {
    right: 0px;
    top: -6px;
  }

  .minus-button {
    left: 0;
    top: -6px;
  }

  .value {
    width: 100%;
    font-weight: bold;
    font-size: 28px !important;
    color: white;
    text-align: center;
    background: #35439B;
    border: 2px solid #13156A;
    -webkit-border-radius: 12px;
    -moz-border-radius: 12px;
    border-radius: 12px;
    height: 60px;
    display: flex;
    justify-content: center;
    flex-direction: column;
    overflow: hidden;
  }
}

@media (max-width: 575.98px) {
  .stake-row {
    flex-direction: column;
    margin-top: 0.5rem;
    .plus-minus-section {
      margin: 10px;
      width: 200px;
      .value {
        height: 30px;
        font-size: 16px !important;
      }
    }
    .minus-button, .plus-button {
      width: 50px;
      height: 50px;
    }
    svg {
      width: 50px;
      height: auto;
    }
    .btn-startgame {
      margin: 10px;
    }
  }
}

// Small devices (landscape phones, 576px and up)
@media (min-width: 576px) {
  .stake-row {
    .btn-startgame {
      margin-left: auto;
      font-size: 26px;
    }
  }
}

// Large devices (desktops, 992px and up)
@media (min-width: 768px) {
  .plus-minus-section {
    width: 280px;
  }
}
</style>

<template >
  <div class="flex-horizontal stake-row">
    <div class="plus-minus-section">
      <button class="btn-harmony minus-button"
              @click="minus"
              :disabled="globalData.stake <= 20">
        <svg width="52" height="12" viewBox="0 0 52 12" fill="none" xmlns="http://www.w3.org/2000/svg">
          <rect width="52" height="12" rx="6" fill="white"/>
        </svg>
      </button>

      <div class="value">
        {{ globalData.stake }}
      </div>

      <button class="btn-harmony plus-button"
              @click="plus"
              :disabled="globalData.stake + 20 > globalData.balance">
        <svg width="52" height="52" viewBox="0 0 52 52" fill="none" xmlns="http://www.w3.org/2000/svg">
          <rect y="20" width="52" height="12" rx="6" fill="white"/>
          <rect x="32" width="52" height="12" rx="6" transform="rotate(90 32 0)" fill="white"/>
        </svg>
      </button>
    </div>

    <button
      class="btn-harmony btn-startgame"
      @click="stakeToken"
      :disabled="globalData.balance < 20"
    >Start Game</button>
  </div>
</template>

<script>
import service from "../service";
import store from "../store";
export default {
  name: "StakeRow",
  data() {
    return {
      globalData: store.data
    };
  },
  methods: {
    minus() {
      if (this.globalData.stake <= 20) return;
      this.globalData.stake -= 20;
    },
    plus() {
      if (this.globalData.stake + 20 > this.globalData.balance) return;
      this.globalData.stake += 20;
    },
    stakeToken() {
      playBackgroundMusic();
      service
        .stakeToken(this.globalData.privkey, this.globalData.stake)
        .then(() => {
          this.$emit("stake", this.globalData.stake);
        });
    }
  }
};
</script>
