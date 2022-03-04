import ElFormItem from '../form/src/form-item.js';

/* istanbul ignore next */
ElFormItem.install = function(Vue) {
  Vue.component(ElFormItem.name, ElFormItem);
};

export default ElFormItem;
