<template>
  <div>
    <h4>Sign Up</h4>
    <form @submit.prevent="register">

      <label for="name">Name</label>
      <div>
        <input id="name" type="text" v-model="name" required autofocus>
      </div>

      <label for="surname">Surname</label>
      <div>
        <input id="surname" type="text" v-model="surname" required autofocus>
      </div>

      <label for="email">E-Mail Address</label>
      <div>
        <input id="email" type="email" v-model="email" required>
      </div>

      <label for="password">Password</label>
      <div>
        <input id="password" type="password" v-model="password" required>
      </div>
<!--      <label for="password-confirm">Confirm Password</label>-->
<!--      <div>-->
<!--        <input id="password-confirm" type="password" v-model="password_confirmation" required>-->
<!--      </div>-->

      <label for="city">City</label>
      <div>
        <input id="city" type="text" v-model="city" required autofocus>
      </div>

      <label for="age">Age</label>
      <div>
        <input id="age" type="number" min="1" max="150" v-model="age" required autofocus>
      </div>

      <label for="age">Sex</label>
      <div>
        <select v-model="sex" required autofocus>
          <option value="male">Male</option>
          <option value="female">Female</option>
        </select>
      </div>

      <label for="interests">Interests</label>
      <div>
        <input id="interests" type="text" v-model="interests" autofocus>
      </div>

      <br>

      <div>
        <button type="submit">Sign Up</button>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      name: "",
      surname: "",
      email: "",
      password: "",
      password_confirmation: "",
      age: "",
      city: "",
      sex: "",
      interests: "",
    };
  },
  methods: {
    register: function() {
      let data = {
        name: this.name,
        surname: this.surname,
        email: this.email,
        password: this.password,
        age: this.age,
        city: this.city,
        sex: this.sex,
        //interests: this.interests,
      };
      this.$store
          .dispatch("register", data)
          .then(() => this.$router.push("/"))
          .catch(err => console.log(err));
    }
  }
};
</script>
