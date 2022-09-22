
// Vuex.mapState in computed (received events)
// Vuex.mapMutations in methods, (emit events)

export default {
  data() {
    return {
    }
  },
  computed: {
    snackbar: {
      get() {
        return (this.$store.state.gen.errorText !== '') || (this.$store.state.gen.msgText !== '')
      },
      set(newVal) {
        if (!newVal) {
          this.$store.commit('clearErrorText')
          this.$store.commit('clearMsgText')
        }
      }
    },
    ...Vuex.mapState({
      textMsg: state => {
        if (state.gen.errorText !== '') {
          return state.gen.errorText
        }
        return state.gen.msgText
      },
      colorsnack: state => {
        if (state.gen.errorText !== "") {
          return "red darken-4"
        }
        return ''
      }
    })
  },
  methods: {
    closeToast() {
      console.log('close toast is called')
      this.$store.commit('clearErrorText')
      this.$store.commit('clearMsgText')
    },
  },
  template: `
  <v-snackbar v-model="snackbar" :color="colorsnack">
      {{ textMsg }}
      <v-btn dark text @click="closeToast">Close</v-btn>
    </v-snackbar>
`
}
