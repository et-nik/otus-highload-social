<template>
  <div id="app">
    <user-menu></user-menu>
    <router-view></router-view>
  </div>
</template>

<script>
  import UserMenu from "@/components/UserMenu";

  export default {
    name: 'App',
    components: {
      'user-menu': UserMenu,
    }
  }
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

.nav {
  text-align: center;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}

table {
  width: 100%;
  margin-bottom: 20px;
  border: 1px solid #dddddd;
  border-collapse: collapse;
  margin-top: 20px;
}
table th {
  font-weight: bold;
  padding: 5px;
  background: #efefef;
  border: 1px solid #dddddd;
}
table td {
  border: 1px solid #dddddd;
  padding: 5px;
}
</style>
