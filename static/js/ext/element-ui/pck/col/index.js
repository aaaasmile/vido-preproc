import ElCol from './src/col.js';

/* istanbul ignore next */
ElCol.install = function(Vue) {
  Vue.component(ElCol.name, ElCol);
};

export default ElCol;

