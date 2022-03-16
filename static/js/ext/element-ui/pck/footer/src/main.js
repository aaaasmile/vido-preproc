export default {
  name: 'ElFooter',

  componentName: 'ElFooter',

  props: {
    height: {
      type: String,
      default: '60px'
    }
  },
  template: `
  <footer class="el-footer" :style="{ height }">
    <slot></slot>
  </footer>`
};
