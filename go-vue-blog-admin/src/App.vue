<template>
  <div id="app">
    <!--<img src="./assets/logo.png">-->
    <back-menu></back-menu>
    <router-view></router-view>
  </div>
</template>

<script>
    import Menu from './components/Menu';
    export default {
        name: 'app',
        components: {
            'back-menu': Menu
        }
    };
</script>

<style lang="scss">
  @import './assets/sass/app.scss';
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    height: 100%;
    /* text-align: center;
    color: #2c3e50;
    margin-top: 60px; */
  }
</style>