<script>
  export default {
    methods: {
      logout: function() {
        this.$store
            .dispatch("logout")
            .then(() => this.$router.push("/"))
            .catch(err => console.log(err));
      }
    },
    beforeMount(){
      this.logout()
    },
  };
</script>
