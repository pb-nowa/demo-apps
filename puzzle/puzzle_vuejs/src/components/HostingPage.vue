<style scoped lang="less">
.host {
  width: 100%;
  height: 100%;
}
</style>

<template >
  <div class="host">
    <welcome-page @join="join" v-if="step === 0"></welcome-page>
    <email-page @submit="submitEmail" v-if="step === 1"></email-page>
    <key-page :userKey="userKey" v-if="step === 2" @start="startGame"></key-page>
    <stake-page @stake="stake" @seeTutorial="seeTutorial" v-if="step === 3"></stake-page>
    <tutorial-page @done="doneTutorial" v-if="step === 4"></tutorial-page>
    <puzzle-page
      @openModal="openModal"
      @restart="restartGame"
      v-if="step === 5"></puzzle-page>

    <new-home-page @join="join" v-if="step === 10"></new-home-page>
    <new-welcome-page @join="join" v-if="step === 11"></new-welcome-page>
    <enter-page @join="join" v-if="step === 12"></enter-page>

    <b-button v-b-modal.modal-harmony>Launch demo modal</b-button>

    <b-modal id="modal-harmony" title="BootstrapVue" hide-footer hide-header>
      <p class="my-4 content">{{ this.modalMessage }}</p>
      <div class="buttons">
        <button class="btn-harmony"
                @click="$bvModal.hide('modal-harmony')">
          {{ noMessage }}
        </button>
        <button class="btn-harmony btn-right btn-harmony-secondary"
                @click="modalClickFn">
          {{ yesMessage }}
        </button>
      </div>
    </b-modal>
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
      userKey: "Oxhsa89sd23jkl3450stypose00",
      modalMessage: 'Do you really want to quit?',
      noMessage: '',
      yesMessage: '',
      modalClickFn: function () {}
    };
  },
  mounted: function() {},
  methods: {
    join(index) {
      if (index) {
        this.step = index;
        return;
      }

      this.step++;
    },
    openModal(yesMessage, noMessage, modalMessage, clickFn) {
      console.log("eeeeeeee openModal")
      this.yesMessage = yesMessage;
      this.noMessage = noMessage;
      this.modalMessage = modalMessage;
      this.modalClickFn = clickFn;

      this.$bvModal.show('modal-harmony');
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
