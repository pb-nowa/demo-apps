<template>
  <div class="chip" :style="style">{{ value }}</div>
</template>

<script>
const colors = [
  {
    bg: "#D4D3D4",
    fg: "white",
    insetColor: "#6A7C8B"
  },
  {
    bg: "#3FE6FF",
    fg: "white",
    insetColor: "#207F8D"
  },
  {
    bg: "#02D69D",
    fg: "white",
    insetColor: "#27876E"
  },
  {
    bg: "#482AFF",
    fg: "white"
  },
  {
    bg: "#BDFF00",
    fg: "#1B295E"
  },
  {
    bg: "#758796",
    fg: "white"
  },
  {
    bg: "#C6F6EF",
    fg: "#1B295E"
  },
  {
    bg: "#19586D",
    fg: "white"
  },
  {
    bg: "#01DBA7",
    fg: "#1B295E"
  }
];

export default {
  name: "Chip",
  props: ["value","boardSizePx"],
  //TODO:nit: extract these two lines as getColorIndex.
  computed: {
    backColor: function() {
      let idx = this.value % colors.length;
      if (idx < 0) idx += colors.length;
      return colors[idx].bg;
    },
    color: function() {
      let idx = this.value % colors.length;
      if (idx < 0) idx += colors.length;
      return colors[idx].fg;
    },
    insetColor: function() {
      let idx = this.value % colors.length;
      if (idx < 0) idx += colors.length;
      return colors[idx].insetColor;
    },
    style() {
      let boxShadowStyle;
      if (this.insetColor) {
        boxShadowStyle = `inset 0px 0px 0px 0px transparent, inset 0px -9px 0px 0px ${this.insetColor}`;
      } else {
        boxShadowStyle = "none";
      }
      return {
        textStroke: "3px black",
        fontSize: this.boardSizePx/8 + "px",
        backgroundColor: this.backColor,
        color: this.color,
        boxShadow: boxShadowStyle
      };
    }
  }
};
</script>
