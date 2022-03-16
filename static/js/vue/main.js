import store from './store/index.js'
import routes from './routes.js'
import Toast from './components/toast.js'
//Element ui components
import Button from '../ext/element-ui/pck/button/index.js'
import Tooltip from '../ext/element-ui/pck/tooltip/index.js'
import Form from '../ext/element-ui/pck/form/index.js'
import FormItem from '../ext/element-ui/pck/form-item/index.js';
import Input from '../ext/element-ui/pck/input/index.js';
import Container from '../ext/element-ui/pck/container/index.js';
import Row from '../ext/element-ui/pck/row/index.js';
import Col from '../ext/element-ui/pck/col/index.js';

const components = [
  Container,
  Row,
  Col,
  Button,
  Form,
  FormItem,
  Input,
  Tooltip
]

components.forEach(component => {
  Vue.component(component.name, component);
});

export const app = new Vue({
  el: '#app',
  router: new VueRouter({ routes }),
  components: { Toast },
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
    <Toast></Toast>
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
            <p>
              <i class="copyright icon"></i> {{ new Date().getFullYear() }} by
              Invido.it
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
`
})

console.log('Main is here!')