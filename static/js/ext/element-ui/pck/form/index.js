import ElForm from './src/form.js';

/* istanbul ignore next */
ElForm.install = function(Vue) {
  Vue.component(ElForm.name, ElForm);
};

export default ElForm;
