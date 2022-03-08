import ElInput from './src/input.js';

/* istanbul ignore next */
ElInput.install = function(Vue) {
  Vue.component(ElInput.name, ElInput);
};

export default ElInput;
