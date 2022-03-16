import Container from './src/main.js';

Container.install = function(Vue) {
  Vue.component(Container.name, Container);
};

export default Container;
