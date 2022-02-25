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
  <div
    class="
      p-4
      max-w-sm
      mx-auto
      bg-white
      rounded-xl
      shadow-lg
      flex
      items-center
      space-x-4
    "
  >
    <div class="container mx-auto">
      <h2>ChitChat</h2>
      <p class="text-slate-500">You have a new message!</p>
    </div>
  </div>`
})

console.log('Main is here!')