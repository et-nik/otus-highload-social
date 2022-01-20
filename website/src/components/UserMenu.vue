<template>
  <div class="nav" v-if="isLoggedIn">
    <a href="/profile">Profile</a>&nbsp;|&nbsp;
    <a href="/friends">Friends</a>&nbsp;|&nbsp;
    <a href="/users">Users</a>&nbsp;|&nbsp;
    <a href="/sign-out">Sign Out</a>
  </div>
  <div class="nav" v-else>
    <a href="/sign-in">Sign In</a>&nbsp;|&nbsp;
    <a href="/sign-up">Sign Up</a>
  </div>

</template>

<script>
export default {
  name: 'UserMenu',
  computed: {
    isLoggedIn() {
      return this.$store.getters.isLoggedIn
    }
  },
}
</script>
