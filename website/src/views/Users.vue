<template>
  <div class="users">
    <table style="width:100%">
      <thead>
      <tr>
        <th>Name</th>
        <th>Surname</th>
        <th>Age</th>
        <th>Sex</th>
        <th>City</th>
        <th></th>
      </tr>
      </thead>
      <tbody v-for="user in users" :key="user.id">
      <tr>
        <td>{{user.name}}</td>
        <td>{{user.surname}}</td>
        <td>{{user.age > 0 ? user.age : '-'}}</td>
        <td>{{user.sex}}</td>
        <td>{{user.city}}</td>
        <td>
          <button v-if="!isUserFollowed(user.id) && !isCurrentUser(user.id)" @click="follow(user.id)">Follow</button>
          <span v-else-if="isCurrentUser(user.id)">You</span>
          <span v-else>Followed</span>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      users: [],
    }
  },
  methods: {
    loadUsers: function() {
      axios({ url: '/users', method: 'GET', mode: 'no-cors'})
          .then(resp => {
            this.users = resp.data
          })
          .catch(err => {
            console.log(err)
          })
    },
    isCurrentUser: function(id) {
      return this.profile.id === id
    },
    isUserFollowed: function(id) {
      if (this.profile.friends === null) {
        return false
      }

      return this.profile.friends.includes(id)
    },
    follow: function(id) {
      this.$store.dispatch("follow", id)
          .catch(err => console.log(err));
    },
  },
  computed: {
    profile() {
      return this.$store.getters.profile
    },
  },
  beforeMount(){
    this.loadUsers()
  },
};
</script>
