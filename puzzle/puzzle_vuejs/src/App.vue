<style lang="less">
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}
html {
  height: 100%;
}
body,
#app {
  height: 100%;
}

#app {
  background: url(./assets/bg.png);
  background-repeat: no-repeat;
}

body,
html {
  font-family: "Fira Sans", "Source Sans Pro", Arial, sans-serif;
  color: #2c3e50;
  font-size: 18px;
  margin: 0;
}

.btn-harmony {
  font-family: inherit;
  cursor: pointer;
  color: white;
  padding: 0.5em 1.5em 0.6em 1.5em !important;
  background: #05AAFF !important;
  border: 2px solid #13156A !important;
  border-radius: 12px !important;
  box-shadow:
    inset 0px 0px 0px 0px transparent,
    inset 0px -7px 0px 0px #304FBD;
}

.btn-level, .btn-reset {
  font-size: 18px;
  display: flex;
  align-items: center;

  svg {
    margin-right: 10px;
  }
}

.tabs {

  .tab-headers {
    display: flex;
  }

  .tab-header {
    display: flex;
    flex-grow: 1;
    flex-basis: 0;
    align-items: center;
    padding: 1rem 1.5rem;
    margin-right: 5px;
    border-top-left-radius: 12px;
    border-top-right-radius: 12px;
    border: 5px solid #243572;
    border-bottom: none;
    color: #747FA8;
    background: #243572;
  }

  .tab-header.active {
    color: white;
    background: #01A0FF;
    font-weight: bold;
  }

  .tab-header:last-child {
    margin-right: 0;
  }

  .list {
    list-style: none;
  }

  .item {
    padding: 3px;
    background: #13173C;

    .inner {
      display: flex;
      align-items: center;
      color: white;
      background: #2F6EB2;
      border: 2px solid #13173C;
      border-radius: 8px;
      box-shadow:
        inset 0px 0px 0px 0px transparent,
        inset 0px -5px 0px 0px #244C83;
    }
  }

  .item:last-child {
    border-bottom-right-radius: 8px;
    border-bottom-left-radius: 8px;
  }

  .badge {
    position: relative;
    font-weight: bold;
    border-top-left-radius: 8px;
    border-bottom-left-radius: 8px;

    img {
      width: 40px;
      position: absolute;
      z-index: -1;
    }
  }

  .reward {
    display: flex;
    align-items: center;
    margin-left: auto;
    margin-right: 80px;

    svg {
      margin-right: 5px;
    }
  }

  .name {
    margin-left: 20px;
  }

  .badge-1 {
    font-weight: bold;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 50px;
    height: 50px;
  }
}

.white-color {
  color: white;
}



.harmony-input {
  font-family: inherit;
  padding: 0.5em 1.5em 0.6em 1.5em !important;
  font-size: 24px;
  text-align: center;
  font-weight: bold;
  background: white;
  border: 2px solid #13156A;
  border-radius: 12px !important;
  box-shadow:
    inset 0px 0px 0px 0px transparent,
    inset 0px -5px 0px 0px #7C90E6;
}

.main-container {
  display: flex;
  flex-direction: column;
  margin: 0 auto;
  padding: 1em;
  /* Hack to improve transition performance on mobile devices. It enables GPU rendering. */
  transform: translateZ(0);
}

.main-container .score-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
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

@-webkit-keyframes up-disappear {
  0% {
    opacity: 0.7;
  }

  100% {
    opacity: 0;
    -webkit-transform: translateY(-40px);
  }
}

@keyframes chip-value-changed {
  0% {
    transform: scale(1);
  }

  50% {
    transform: scale(1.2);
  }

  100% {
    transform: scale(1);
  }
}

.btn-primary,
.btn-mini {
  background-color: #482bff;
  border-radius: 0.3em;
  border: 0;
  color: #fff;
  outline: none;
  &:disabled {
    opacity: 0.5;
    background-color: #aaa;
  }
}

.btn-primary {
  padding: 0.7em 1.0em;
  font-size: 1.1em;
  margin-right: 0em;
  margin-left: 0em;
  margin-top: 1em;

}

.btn-mini {
  display: inline-block;
  font-size: 0.8em;
  width: 2em;
  height: 2em;
  padding: 0;
  outline: none;
}

.logo {
  margin: 1em auto;
  color: red;
  height: 2em;
  width: 2em;
  background-image: url(./assets/logo.svg);
  background-size: contain;
  background-repeat: no-repeat;
}

.flex-veritcal {
  display: flex;
  flex-direction: column;
}

.flex-horizontal {
  display: flex;
  flex-direction: row;
  align-items: center;
}
.flex-hv-center {
  display: flex;
  align-items: center;
  justify-content: center;
}

.flex-grow {
  flex: auto;
  min-width: 0;
  min-height: 0;
}
.text-right {
  text-align: right;
}

.info-item {
  .flex-horizontal;
  color: #fff;
  padding: 0.5em;
  width: 5em;
  text-align: center;
  .content {
    margin-left: 0.5em;
    font-weight: bold;
  }
}

body {
}

.new-logo {
  width: 282px;
  height: 198px;
  margin: 0 auto;
  margin-top: 45px;
  background-image: url(./assets/new/puzzle-logo.png);
  background-size: contain;
  background-repeat: no-repeat;
}

p {
  margin: 0;
}
p + p {
  margin-top: 1em;
}
</style>

<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: "App"
};
</script>

<style>
</style>
