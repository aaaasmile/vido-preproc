import Tooltip from './src/main.js';

/* istanbul ignore next */
Tooltip.install = function(Vue) {
  Vue.component(Tooltip.name, Tooltip);
};

export default Tooltip;
