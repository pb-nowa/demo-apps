<style scoped lang="less">
.host {
  max-width: 1600px;
  margin: 0 auto;
  height: 100%;
}
</style>

<template >
  <div>
<!--    <div class="host">-->
    <welcome-page @join="join" v-if="step === 0"></welcome-page>
    <email-page @submit="submitEmail" v-if="step === 1"></email-page>
    <key-page :userKey="userKey" v-if="step === 2" @start="startGame"></key-page>
    <stake-page @stake="stake" @seeTutorial="seeTutorial" v-if="step === 3"></stake-page>
    <tutorial-page @done="doneTutorial" v-if="step === 4"></tutorial-page>
    <puzzle-page @restart="restartGame" v-if="step === 5"></puzzle-page>

    <new-home-page @join="join" v-if="step === 10"></new-home-page>
    <new-welcome-page @join="join" v-if="step === 11"></new-welcome-page>
    <enter-page @join="join" v-if="step === 12"></enter-page>
  </div>
</template>

<script>
import WelcomePage from "./WelcomePage";
import PuzzlePage from "./PuzzlePage";
import EmailPage from "./EmailPage";
import KeyPage from "./KeyPage";
import StakePage from "./StakePage";
import service from "../service";

import NewHomePage from "./New_HomePage";
import NewWelcomePage from "./New_WelcomePage";
import EnterPage from "./EnterPage";
import TutorialPage from "./TutorialPage";

const StakePageIndex = 3;
const TutorialPageIndex = 4;
const PuzzlePageIndex = 5;

export default {
  name: "HostingPage",
  components: {
    WelcomePage,
    EmailPage,
    KeyPage,
    TutorialPage,
    StakePage,
    PuzzlePage,
    NewHomePage,
    NewWelcomePage,
    EnterPage
  },
  data() {
    return {
      step: 5,
      userKey: "Oxhsa89sd23jkl3450stypose00"
    };
  },
  mounted: function() {},
  methods: {
    join(index) {
      console.log("eeeeeeee index", index)
      if (index) {
        this.step = index;
        return;
      }

      this.step++;
    },
    submitEmail(email) {
      service.register(email).then(() => {
        this.step++;
      });
    },
    stake(key, value) {
      service.stakeToken(key, value).then(() => {
        this.step = PuzzlePageIndex;
      });
    },
    startNewWelcome() {
      this.step = NewWelcomePage;
    },
    startGame() {
      this.step++;
    },
    seeTutorial() {
      this.step = TutorialPageIndex;
    },
    doneTutorial() {
      this.step = StakePageIndex;
    },
    restartGame() {
      this.step = PuzzlePageIndex;
    }
  }
};
</script>
