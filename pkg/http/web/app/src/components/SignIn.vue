<template>
  <div>
    <form class="login" @submit.prevent="login">
      <h1>Sign in</h1>

      <label>Email</label>
      <div>
        <input required v-model="email" type="email" placeholder="Name"/>
      </div>

      <label>Password</label>
      <div>
        <input required v-model="password" type="password" placeholder="Password"/>
      </div>

      <hr/>
      <button type="submit">Sign In</button>
    </form>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        email: "",
        password: ""
      };
    },
    methods: {
      login: function() {
        let email = this.email;
        let password = this.password;
        this.$store
            .dispatch("login", { email, password })
            .then(() => this.$router.push("/profile"))
            .catch(err => console.log(err));
      }
    }
  };
</script>
