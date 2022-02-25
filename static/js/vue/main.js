import store from './store/index.js'
import routes from './routes.js'


export const app = new Vue({
  el: '#app',
  router: new VueRouter({ routes }),
  components: {  },
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
  <div>
    <router-view></router-view>
    <div class="ui vertical footer segment">
      <div class="ui container">
        <div class="ui stackable divided equal ten stackable grid">
          <div class="five wide column">
            <h4 class="ui header">Version</h4>
            <p>Software build {{ Buildnr }}</p>
          </div>
          <div class="seven wide column">
            <h4 class="ui header">Info</h4>
            <p><i class="copyright icon"></i> {{ new Date().getFullYear() }} by Invido.it</p>
          </div>
        </div>
      </div>
    </div>
  </div>
})

console.log('Main is here!')