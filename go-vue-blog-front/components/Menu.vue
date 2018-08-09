<template>
  <nav class="menu">
    <nuxt-link to="/" class="no-underline">
      <svg class="icon" aria-hidden="true">
        <use xlink:href="#icon-home"></use>
      </svg>首页
    </nuxt-link>
    <nuxt-link to="/archive" class="no-underline">
      <svg class="icon" aria-hidden="true">
        <use xlink:href="#icon-archive"></use>
      </svg>归档
    </nuxt-link>
    <nuxt-link to="/about" class="no-underline">
      <svg class="icon" aria-hidden="true">
        <use xlink:href="#icon-about"></use>
      </svg>关于我
    </nuxt-link>
    <a href="https://github.com/louyan" target="_blank" class="no-underline" @click="goToOtherWebsite('GITHUB')">
      <svg class="icon" aria-hidden="true">
        <use xlink:href="#icon-github"></use>
      </svg>GITHUB
    </a>
  </nav>
</template>

<script>
    export default {

        components: {},
    }
</script>

<style lang="scss" scoped>
  @import "../assets/sass/app.scss";
  .menu {
    position: fixed;
    top: 5.2em;
  a {
    display: block;
    margin-bottom: 0.8em;
    padding-left: 1em;
    height: 3em;
    font-size: 1.2em;
    color: #666;
    line-height: 3em;

  &.nuxt-link-exact-active {
     color: $base-color;
     background-color: #fff;
   }
  &.juejin {
  .icon {
    vertical-align: -0.05em;
  }
  }
  }
  .icon {
    width: 1em;
    height: 1em;
    margin-right: 0.5em;
  }
  }
</style>



