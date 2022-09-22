import store from './store/index.js?version=1'
import routes from './routes.js?version=1'

export const app = new Vue({
  el: '#app',
  router: new VueRouter({ routes }),
  vuetify: new Vuetify(),
  store,
  data() {
    return {
      Buildnr: "",
    }
  },
  computed: {
    ...Vuex.mapState({
    
    })
  },
  created() {
    this.Buildnr = window.myapp.buildnr
  },
  methods: {

  },
  template: `
  <v-app class="grey lighten-4">
    <v-content class="mx-4 mb-4">
      <router-view></router-view>
    </v-content>
    <v-footer absolute>
      <v-col class="text-center caption" cols="12">
        {{ new Date().getFullYear() }} —
        <span>Buildnr: {{Buildnr}}</span>
      </v-col>
    </v-footer>
  </v-app>
`
})

console.log('Main is here!')