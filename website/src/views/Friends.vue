<template>
  <div class="friends">
    <table style="width:100%">
      <thead>
      <tr>
        <th>Name</th>
        <th>Surname</th>
        <th>Age</th>
        <th>Sex</th>
        <th>City</th>
      </tr>
      </thead>
      <tbody v-for="user in users" :key="user.id">
        <tr>
          <td>{{user.name}}</td>
          <td>{{user.surname}}</td>
          <td>{{user.age > 0 ? user.age : '-'}}</td>
          <td>{{user.sex}}</td>
          <td>{{user.city}}</td>
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
      loadFriends: function() {
        axios({ url: '/profile/friends', method: 'GET', mode: 'no-cors'})
            .then(resp => {
              this.users = resp.data
            })
            .catch(err => {
              console.log(err)
            })
      }
    },
    beforeMount(){
      this.loadFriends()
    },
  };
</script>
