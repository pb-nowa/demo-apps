<style scoped lang="less">
.redeem-help-link:before {
  content: "want binance referral id?";

  // width: 25px;
  // height: 25px;
  // line-height: 25px;
  // display: inline-block;
  // background: #828282;
  // border-radius: 50%;
  // color: #fff;
  // font-weight: normal;
}

.btn-primary {
  background: #1B295E;
}

.score-container {
  margin-right: auto;
  margin-left: auto;
  justify-content: space-between;
  display: flex;
  flex: 1;
  flex-direction: row;
}

footer {
  margin: 1em auto 0;
  .btn-primary {
    font-size: 1em;
  }
}

.board-wrapper {
  position: relative;
  margin: 0 auto;
  flex-grow: 0;
  flex-shrink: 0;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}

.game-over-message {
  font-weight: bold;
  text-align: center;
  background-color: rgba(255, 255, 255, 1);
  border-radius: 0.3em;
}

.content-tutorial {
  height: 100%;
  padding: 0 12px;
  background: rgba(255, 255, 255, 0.2);
  display: flex;
  align-items: center;
  justify-content: center;
}

.continue-playing-button {
  margin-top: 9.17892px;
  background: none;
  border: 1px solid #1B295E;
  border-radius: 6px;
  color: #1B295E;
}

.content-level10 {
  .buttons {
    margin: 20px 0;
    width: 210px;
    margin: 0 auto;
    .btn-primary {
      width: 100%;
    }
  }
}

.main-container {
  height: 100%;
  .game-container {
    position: relative;
    display: flex;
    flex-direction: column;
    height: 100%;
    .overlay {
      width: 100%;
      height: 100%;
      position: absolute;
      z-index: 2;
      display: flex;
      align-items: center;
      justify-content: center;
    }
  }
}

.blur-text {
  opacity: 0.8;
}

.appearing {
  animation: appearing 1.2s;
  -webkit-animation: appearing 1.2s;
}
input{
    text-align:center;
}
.is-redeemed {
  animation: disappearing 3s;
  -webkit-animation: disappearing 3s;
}

@keyframes disappearing {
  0% {
    opacity: 1;
  }

  25% {
    opacity: 1;
  }

  100% {
    opacity: 0;
  }
}

@keyframes appearing {
  0% {
    opacity: 0;
  }

  100% {
    opacity: 1;
  }
}

.action-row + .action-row {
  margin-top: 1em;
}
.info-item {
  align-self: flex-end;
  .content {
    position: relative;
  }
}
.count-down {
  .seconds-left {
    &.hurry-up {
      color: #f6371d;
    }
    &.hurry-up {
      animation-name: headShake;
      animation-duration: 1s;
      animation-timing-function: ease-in-out;
      animation-iteration-count: infinite;
    }
  }
}

@keyframes headShake {
  0% {
    transform: translateX(0);
  }

  6.5% {
    transform: translateX(-6px) rotateY(-9deg);
  }

  18.5% {
    transform: translateX(5px) rotateY(7deg);
  }

  31.5% {
    transform: translateX(-3px) rotateY(-5deg);
  }

  43.5% {
    transform: translateX(2px) rotateY(3deg);
  }

  50% {
    transform: translateX(0);
  }
}

.number-increase {
  display: block;
  position: absolute;
  left: 0;
  top: 0;
  color: #2c3e50;
  width: 100%;
  animation: up-disappear 1.5s;
}
@keyframes up-disappear {
  0% {
    opacity: 0.7;
  }

  100% {
    opacity: 0;
    transform: translateY(-40px);
  }
}
.link-footer {
  left: 0;
  width: 100%;
  padding: 1em;
  text-align: center;
  z-index: 1000;
  flex: 1;
  display: flex;
  align-items: flex-end;
  justify-content: center;
  min-height: 40px;
}
.fake-footer {
  flex: 1;
}
.icon-clock,
.icon-token {
  background-size: contain;
}

.icon-clock {
  background-image: url(../assets/clock.svg);
}
.icon-token {
  background-image: url(../assets/token.svg);
}
.level-text {
  font-weight: bold;
}
.logo {
  align-self: flex-start;
}
.link {
  font-size: 0.8em;
  text-align: center;
  text-decoration: none;
}
.input {
  border-radius: 0.3em;
  border: 1px solid black;
  background-color: #f5f5f5;
  display: block;
  width: 100%;
  overflow: auto;
  -webkit-appearance: none;
  font-size: initial;
  outline: none;
}

.inputs {
  display: flex;
  margin: 0 10px;
}

  .is-level10 {
    .inputs {
      background-color: "red" !important;
    }

    .input-error {
      font-size: 14px;
      color: red;
    }

    .texts {
      margin-top: 10px;
    }
  }

  img.loading {
    width: 50px;
  }

  .redeemed-section {
    font-size: 50px;
    color: darkgreen;
  }

  .lose-hope-image {
    img {
      width: 90px;
    }

    margin-bottom: 10px;
  }

.congrats-trophy {
  margin-top: -5px;
}

// Large devices (desktops, less than 1200px)
@media (max-width: 1199.98px) {
  .congrats-trophy img {
    width: 150px;
  }
}

// Medium devices (tablets, less than 992px)
@media (max-width: 991.98px) {
  .congrats-trophy img {
    width: 100px;
  }
}

// Small devices (landscape phones, less than 768px)
@media (max-width: 767.98px) {
  .congrats-trophy img {
    width: 100px;
  }
}

// Extra small devices (portrait phones, less than 576px)
@media (max-width: 575.98px) {
  .congrats-trophy img {
    width: 100px;
  }
}

</style>

<template>
  <div id="app">
    <div class="main-container appearing">
      <div class="game-container" ref="gameContainer">
        <a
          :href="'https://explorer2.harmony.one/#/address/' + globalData.address"
          class="logo"
          target="_blank"
        ></a>
        <div class="score-container" :style="{ width: boardSizePx + 'px' }">
          <div class="balance info-item" :style="infoItemStyle">
            <div class="label">
              <div class="icon-token" :style="iconTokenStyle"></div>
            </div>
            <div class="content">
              {{ globalData.balance }}
              <transition>
                <span v-if="balanceIncrease!=''" class="number-increase">{{ balanceIncrease }}</span>
              </transition>
            </div>
          </div>
          <div class="count-down info-item" :style="infoItemStyle">
            <div class="label">
              <div class="icon-clock" :style="iconClockStyle"></div>
            </div>
            <div class="content">
              <div
                class="seconds-left"
                :class="{ 'hurry-up': secondsLeft && secondsLeft <= 12, 'game-over': !secondsLeft }"
              >{{ secondsLeft | time }}</div>
              <transition>
                <span v-if="timeIncrease!=''" class="number-increase">
                  {{
                  timeIncrease
                  }}
                </span>
              </transition>
            </div>
          </div>
        </div>

        <div class="board-wrapper" :style="boardWrapperStyle">
          <div v-if="gameEnded">
<!--            <div v-if="true">-->
            <div class="overlay game-over-message appearing">
              <div class="content content-level10">
                <div>
                  <div class="blur-text" :style="gameTutorialStyle" v-if="this.levelIndex > showCouponLevel">
                    <div class="congrats-trophy">
                      <img src="../assets/congrats_trophy.svg" alt="">
                    </div>
                    <span :style="gameTutorialSmallStyle">Congrats!</span>
                    <br>
                    <span :style="gameTutorialSmallStyle">You finished level {{ this.levelIndex }}</span>
                    <br>
                    <div
                      v-if="gameEnded"
                      style="
                        font-size: 17px;
                        margin: 2em 0 1.5em;
                      "
                    >Tweet your success!</div>
                  </div>
                </div>

                <div>
                  <div class="blur-text" :style="gameTutorialStyle" v-if="showNoLoseLevel">
                    <div class="lose-hope-image">
                      <img src="../assets/lose_hope.svg" alt="">
                    </div>
                    <span :style="gameTutorialSmallStyle">Don't lose hope!</span>
                    <br>
                    <span :style="gameTutorialSmallStyle">Try Again!</span>
                    <br>
                    <br>
                    <br>
                  </div>
                </div>

                <div v-if="this.levelIndex === fireworkLevel">
                  <Fireworks/>
                </div>

                <div class="buttons" v-if="showTwitterLevel">
                  <div>
                    <social-sharing :title="twitterTitle"
                                    url=""
                                    inline-template>
                      <network network="twitter">
                        <a class="btn-primary btn-twitter">
                          <i class="fab fa-twitter"></i> Tweet
                        </a>
                      </network>
                    </social-sharing>
                  </div>
                </div>
                <div>

                </div>
              </div>
            </div>
          </div>

          <div class="is-level10"
               v-if="isLevel10 && !gameEnded">
            <div
              v-bind:class="[isRedeemed ? 'is-redeemed' : 'appearing']"
              class="overlay game-over-message">
              <div class="content content-level10">
                <div>
                  <div
                    style="margin-bottom: 25px;"
                    v-if="!isRedeemed && !isRedeeming"
                  >
                    <img
                      style="width: 24%;"
                      src="../assets/win-trophy.svg" alt="Win trophy" class="win-trophy" />

                    <p style = "margin-bottom: 12px;"
                      class="blur-text" :style="gameTutorialStyle">

                      <span :style="gameTutorialSmallGoodStyle">You have a chance to win</span>
                      <br>

                      <span :style="gameTutorialGoodStyle">$20 of ONE tokens</span>
                      <br>

                    </p>
                  </div>

                  <div v-if="isRedeeming" class="loading-section">
                    <img class="loading" src="../assets/loading.svg" alt="">
                  </div>

                  <div v-if="isRedeemed" class="redeemed-section">
                    <i class="fas fa-check-circle"></i>
                  </div>
                 <!-- <span :style="gameTutorialMediumStyle">Enter Binance Coupon Code:</span> -->

                  <div v-if="!isRedeeming" class="inputs">
                    <input
                      :style= "bijanInputStyle" v-if="!isRedeemed" class="input" v-model="couponCode"
                      @input="onCouponChange"
                      placeholder="Enter Binance referral ID" />
                    <!-- <a
                      href="http://harmony.one"
                      class="redeem-help-link"
                    ></a> -->
                    <div class="buttons" style="margin-left: 5px">
                      <button :style="bijanStyle" v-if="!isRedeemed"
                        class="btn-primary" @click="enterCouponCode">Redeem
                      </button>
                    </div>
                  </div>

                  <div v-if="!isRedeeming">
                    <p
                      style="max-width: 20em;"
                      v-bind:class="{'input-error': !isRedeemed, 'input-success': isRedeemed}"
                    >
                      {{this.redeemMessage}}
                    </p>
                  </div>

                  <div class="buttons">
                  <div >
                    <button
                      :style="bijanLessStyle"
                      v-if="!gameEnded && !isRedeemed && !isRedeeming"
                      class="btn-primary continue-playing-button" @click="keepPlaying">
                    Continue Playing
                    </button>
                  </div>
                </div>
                </div>

                <div>
                </div>
              </div>
            </div>
          </div>

          <transition name="fade" v-for="(level, i) in levels" :key="i">
            <Game
              :ref="'game' + i"
              :listen-own-key-events-only="false"
              :tab-index="1"
              :board-size-px="boardSizePx"
              :game="level"
              :gameLevel="levelIndex+1"
              :isLevel10="isLevel10"
              :gameStarted="gameStarted"
              :gameEnded="gameEnded"
              @completeLevel="onLevelComplete"
              v-if="i === levelIndex"
            ></Game>
          </transition>
        </div>

        <stake-row
          v-if="!gameStarted && !gameEnded"
          @stake="startGame"
          :style="stakeRowStyle"
          @stakeToken="resetLevel"
          :isLevel10="isLevel10"
          :gameEnded="gameEnded"
          :playSound="playSound"
          :mute="mute"
        ></stake-row>

        <footer class="flex-vertical" :style="{ width: boardSizePx + 'px' }" v-if="gameStarted">
          <div class="flex-horizontal action-row">
            <span
              class="flex-grow level-text"
              :style="levelTextStyle"
            >Level: {{ levelIndex + 1 }} / {{ levels.length }}</span>
            <button
              v-if="showResetButton"
              class="btn-primary"
              @click="resetLevel"
              :style="{
                visibility: gameEnded ? 'hidden':'visible',
                fontSize: boardSizePx / 20 + 'px'
                }"
            >
              <font-awesome-icon icon="sync"></font-awesome-icon>
            </button>
          </div>
        </footer>
        <footer class="flex-vertical" :style="{ width: boardSizePx + 'px' }" v-if="gameEnded">
          <div class="flex-horizontal action-row">
          <span
              class="flex-grow level-text"
              :style="levelTextStyle"
            >Level: {{ levelIndex === levels.length ? levelIndex : levelIndex + 1 }} / {{ levels.length }}</span>

            <button v-if="gameEnded"  class="btn-primary" @click="reloadGame">
              Play again!
            </button>
          </div>
        </footer>
        <div class="link-footer"></div>
      </div>
    </div>
  </div>
</template>

<script>
import Game from "./Game";
import Chip from "./Chip";
import StakeRow from "./StakeRow";
import TxHistoryLink from "./TxHistoryLink";
import RedeemPanel from "./RedeemPanel";
import { TweenLite } from "gsap/TweenMax";
import Vue from "vue";
import service from "../service";
import store from "../store";
import { levels } from "../level-generator";
import { setInterval, clearInterval } from "timers";
import Fireworks from "./Fireworks";
import { VALIDATE } from "../common/validate";
import utils from '../utils/utils';

const InitialSeconds = 15;
function guid() {
  return "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, function(c) {
    var r = (Math.random() * 16) | 0,
      v = c == "x" ? r : (r & 0x3) | 0x8;
    return v.toString(16);
  });
}

function getParameterByName(name) {
  var undefined;
  name = name.replace(/[\[]/, "\\[").replace(/[\]]/, "\\]");
  var regex = new RegExp("[\\?&]" + name + "=([^&#]*)"),
    results = regex.exec(location.search);
  return results === null
    ? undefined
    : decodeURIComponent(results[1].replace(/\+/g, " "));
}

export default {
  name: "PuzzlePage",
  components: {
    Game,
    Chip,
    StakeRow,
    TxHistoryLink,
    RedeemPanel,
    Fireworks
  },
  data() {
    return {
      // constants
      fireworkLevel: 99,
      showCouponLevel: 1,

      // variables
      globalData: store.data,
      levelIndex: 0,
      // levels: [],
      levels: levels(), // start with level 0 instead of demo
      boardSizePx: 0,
      size: 3,
      gameStarted: false,
      gameEnded: false,
      isLevel10: false,
      secondsLeft: InitialSeconds,
      timer: null,
      timePlayedTimer: null,
      timePlayed: 0,
      timeIncrease: "",
      balanceIncrease: "",
      isMobile: utils.mobilecheck(),
      reward: 0,
      cancelEmail: false,
      couponCode: "",

      // redeem code component
      redeemMessage: "",
      isRedeemed: false,
      isRedeeming: false,

      // sound
      mute: utils.mobilecheck()
    };
  },
  mounted() {
    let _vm = this;
    let id = getParameterByName("cos");
    this.levels = levels();
    //Set board size follow height of screen when change screen size
    window.addEventListener(
      "resize",
      () => {
        this.boardSizePx = Math.min(
          this.$refs.gameContainer.clientWidth,
          window.innerHeight / 1.7
        );
        this.$forceUpdate;
      },
      false
    );
    // Set board size follow height of screen
    this.boardSizePx = Math.min(
      this.$refs.gameContainer.clientWidth,
      window.innerHeight / 1.7
    );
    service.register(id);

    this.$root.$on('social_shares_open', function(network, url) {
      _vm.trackTweet()
    })
  },
  computed: {
    gameOverStyle() {
      return { fontSize: this.boardSizePx / 6 + "px" };
    },
    //TODO
    //can we find a better way to update the styles?
    //I don't want to have so many style related code in the view model.
    // Ideally it should only contain work-flow related logic.
    // Possible solution: setting the font-size of container and use em to control those sizes using css?
    ///Then we only need to use JS to change one thing -- font-size of container.
    gameTutorialStyle() {
      return { fontSize: this.boardSizePx / 14 + "px" };
    },
    /*MOVE CURSOR TO ADJACENT CELLS...*/
    gameTutorialSmallStyle() {
      return { fontSize: this.boardSizePx / 16 + "px" };
    },
    gameTutorialSmallGoodStyle() {
      return { fontSize: this.boardSizePx / 16 + "px",
      color: '#1B295E'};
    },
    gameTutorialTinyStyle() {
      return { fontSize: this.boardSizePx / 32 + "px" };
    },
    gameTutorialMediumStyle() {
      return { fontSize: this.boardSizePx / 24 + "px" };
    },
    bijanStyle() {
      return {
        marginTop: 0,
        marginLeft: 0,
      };
    },
    bijanLessStyle() {
      return {
        "margin-top": 30 + "px",
        "margin-bottom": 10 + "px",
      };
    },
    bijanInputStyle(){
      return {
        // height: this.boardSizePx / 10 + "px"
      };
    },
    gameTutorialLargeStyle() {
      return { fontSize: this.boardSizePx / 8 + "px" };
    },
    gameTutorialGoodStyle() {
      return { fontSize: this.boardSizePx / 14 + "px",
              color: '#1B295E'};
    },
    infoItemStyle() {
      return { fontSize: this.boardSizePx / 18 + "px" };
    },
    levelTextStyle() {
      return { fontSize: this.boardSizePx / 18 + "px" };
    },
    /*tokens and start button on bottom*/
    stakeRowStyle() {
      return {
        width: this.boardSizePx + "px",
        fontSize: this.boardSizePx / 20 + "px"
      };
    },
    titleStyle() {
      return {
        fontSize: this.boardSizePx / 20 + "px"
      };
    },
    iconTokenStyle() {
      return {
        width: this.boardSizePx / 7.6 + "px",
        height: this.boardSizePx / 12 + "px"
      };
    },
    iconClockStyle() {
      return {
        width: this.boardSizePx / 12 + "px",
        height: this.boardSizePx / 12 + "px"
      };
    },
    boardWrapperStyle() {
      return {
        width: this.boardSizePx + "px",
        height: this.boardSizePx + "px"
      };
    },
    level() {
      return this.levels[this.levelIndex];
    },
    /**
     * Twitter title
     * @return {string}
     */
    twitterTitle() {
      return `I finished level ${this.levelIndex} of #harmonypuzzle! See my winning moves on @harmonyprotocol #blockchain https://explorer2.harmony.one/#/address/${this.globalData.address} Play it at https://puzzle.harmony.one`
    },

    showTwitterLevel() {
      return this.levelIndex > this.showCouponLevel;
    },

    showNoLoseLevel() {
      return this.levelIndex <= this.showCouponLevel;
    },

    /**
     * Show reset button
     * @return {boolean}
     */
    showResetButton() {
      return !(this.isLevel10 && !this.gameEnded)
    }
  },
  destroyed() {
    // Remove event change screen
    window.removeEventListener("resize", this.handleResize);
  },
  methods: {
    startGame() {
      this.gameStarted = true;
      this.gameEnded = false;
      this.cancelEmail = false;
      this.levelIndex = 0;
      this.reward = 0;
      this.levels = levels();
      this.secondsLeft = InitialSeconds;
      this.startTimePlayed();
      this.timer = setInterval(() => {
        this.secondsLeft--;
        if (this.secondsLeft <= 0) {
          this.endGame();
        }
      }, 1000);
    },
    resetLevel() {
      this.$refs[`game${this.levelIndex}`][0].reset();
    },

    userGameLevel() {
      return `user-game-level-${this.levelIndex + 1}`
    },

    /***
     * Track analytics current level
     * @param level
     */
    gaTrack(level) {
      const userLevel = this.userGameLevel()
      this.$ga.event('puzzle-game', 'game-level', userLevel, 1)
    },
    /***
     * Track analytics current level
     * @param level
     */
    trackTweet(level) {
      const userLevel = this.userGameLevel()
      this.$ga.event('puzzle-game', 'tweet', userLevel)
    },
    onLevelComplete(moves) {
      this.gaTrack(this.levelIndex);
      this.trackTimePlayed();

      if (this.levelIndex === this.showCouponLevel) {
        this.endLevel10()
        return;
      }

      if (this.levelIndex === this.levels.length - 1) {
        this.endLastGame();
        return;
      }
      service
        .completeLevel(this.globalData.privkey, this.levelIndex + 1, moves)
        .then(rewards => {
          this.levelIndex++;
          let timeChange = 15;
          this.secondsLeft += timeChange;
          this.timeIncrease = `+${timeChange}`;
          this.balanceIncrease = `+${rewards}`;
          this.reward += rewards;
          Vue.nextTick(() => {
            this.timeIncrease = "";
            this.balanceIncrease = "";
          });
        });
    },
    endLevel10() {
      stopBackgroundMusic()
      this.isLevel10 = true;
      clearInterval(this.timer);
      clearInterval(this.timePlayedTimer)
    },
    endGame() {
      stopBackgroundMusic();
      this.isLevel10 = true;
      this.gameEnded = true;
      this.gameStarted = false;
      store.data.stake = 20;
      clearInterval(this.timer);
      clearInterval(this.timePlayedTimer)
      this.timer = null;
    },
    endLastGame() {
      this.endGame();
      this.levelIndex = this.levels.length;
    },
    restart() {
      this.gameEnded = false;
    },
    closeEmailPopup() {
      this.cancelEmail = true;
    },
    keepPlaying() {
      this.playSound();
      this.isLevel10 = false;
      this.levelIndex++;
      this.startTimer();
    },
    startTimer() {
      this.timer = setInterval(() => {
        this.secondsLeft--;
        if (this.secondsLeft <= 0) {
          this.endGame();
        }
      }, 1000);
    },

    startTimePlayed() {
      this.timePlayedTimer = setInterval(() => {
        this.timePlayed++;
      }, 1000);
    },

    resetTimePlayed() {
      this.timePlayed = 0;
    },

    trackTimePlayed() {
      const currentLevel = this.levelIndex + 1;
      this.$ga.event('time-played', `game-level-${currentLevel}`, this.timePlayed.toString())
      this.resetTimePlayed();
    },

    /**
     * Restart the game but still keeps the current point
     */
    restartGame() {
      this.playSound();
      this.levelIndex = 0;
      this.isLevel10 = false;
      this.gameStarted = true;
      this.gameEnded = false;
      this.secondsLeft = InitialSeconds;
      this.startTimer();

      service
        .stakeToken(this.globalData.privkey, this.globalData.stake)
        .then(() => {
          this.$emit("stake", this.globalData.stake);
        });
    },

    /**
     * Reload the game entirely
     */
    reloadGame() {
      this.$ga.event('puzzle-game', 'reload-game', this.userGameLevel())
      window.location.reload();
    },

    validCouponCode(s) {
      return VALIDATE.validateNumber(s);
    },

    /**
     * Enter coupon code
     */
    enterCouponCode() {
      const isValidCouponCode = this.validCouponCode(this.couponCode);
      if (!isValidCouponCode) {
        this.redeemMessage = 'ID is not valid.';
        return;
      }

      this.isRedeeming = true;
      service.submitCoupon(this.couponCode).then((res) => {
        this.redeemMessage = 'You have successfully entered the Binance referral ID';
        this.isRedeemed = true;
        setTimeout(() => {
          this.keepPlaying();
        }, 2000)
      }).catch((err) => {
        console.error(`There was an error while submitting the referral ID ${this.couponCode}: ${err}`)
        this.redeemMessage = 'Server error, please try again later. Sorry for the inconvenience.'
        this.isRedeemed = false;
      }).finally(() => {
        this.isRedeeming = false;
      })
    },

    onCouponChange(e) {
      if (this.couponCode === '') {
        this.redeemMessage = '';
        return;
      }
      const isValidCouponCode = this.validCouponCode(this.couponCode);
      if (!isValidCouponCode) {
        this.redeemMessage = 'Referral ID is not valid.';
        return;
      }

      this.redeemMessage = '';
    },

    playSound() {
      if (this.mute) {
        stopBackgroundMusic()
        return;
      }

      playBackgroundMusic();
    }
  }
};
</script>
