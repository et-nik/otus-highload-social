<template>
  <div class="profile">
    <table>
      <tbody>

      <tr>
        <td><strong>Email</strong></td>
        <td>{{profile.email}}</td>
      </tr>

      <tr>
        <td><strong>Name</strong></td>
        <td>{{profile.name}}</td>
      </tr>

      <tr>
        <td><strong>Surname</strong></td>
        <td>{{profile.surname}}</td>
      </tr>

      <tr>
        <td><strong>Age</strong></td>
        <td>{{age}}</td>
      </tr>

      <tr>
        <td><strong>Sex</strong></td>
        <td>{{profile.sex}}</td>
      </tr>

      <tr>
        <td><strong>City</strong></td>
        <td>{{profile.city}}</td>
      </tr>

      <tr>
        <td><strong>Interests</strong></td>
        <td>{{interests}}</td>
      </tr>

      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  methods: {
    loadProfile: function() {
      this.$store
          .dispatch("profile")
          .then(() => this.$router.push("/profile"))
          .catch(err => console.log(err));
    }
  },
  computed: {
    profile() {
      return this.$store.getters.profile
    },
    age() {
      return this.profile.age > 0
          ? this.profile.age
          : '-'
    },
    interests() {
      if (this.profile.interests) {
        return this.profile.interests.join(', ')
      } else {
        return ''
      }
    }
  },
  // beforeMount(){
  //   this.loadProfile()
  // },
};
</script>
