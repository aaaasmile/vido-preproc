  export default {
    name: 'ElHeader',

    componentName: 'ElHeader',

    props: {
      height: {
        type: String,
        default: '60px'
      }
    },
    template: `
  <header class="el-header" :style="{ height }">
    <slot></slot>
  </header>`
  };
