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
import Footer from '../ext/element-ui/pck/footer/index.js';
import Header from '../ext/element-ui/pck/header/index.js';

const components = [
  Container,
  Row,
  Col,
  Footer,
  Header,
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
  <el-container>
    <Toast></Toast>
    <router-view></router-view>
    <el-footer>
      <el-row>
        <el-col>
          <el-header>Version</el-header>
          <el-row>Software build {{ Buildnr }}</el-row>
        </el-col>
        <el-col>
          <el-header>Info</el-header>
          <el-row>
            <i class="copyright icon"></i> {{ new Date().getFullYear() }} by
            Invido.it
          </el-row>
        </el-col>
      </el-row>
    </el-footer>
  </el-container>
`
})

console.log('Main is here!')