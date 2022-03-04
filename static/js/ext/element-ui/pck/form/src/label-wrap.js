
export default {
    props: {
        isAutoWidth: Boolean,
        updateAll: Boolean
    },

    inject: ['elForm', 'elFormItem'],

    render() {
        var h = arguments[0];

        var slots = this.$slots.default;
        if (!slots) return null;
        if (this.isAutoWidth) {
            var autoLabelWidth = this.elForm.autoLabelWidth;
            var style = {};
            if (autoLabelWidth && autoLabelWidth !== 'auto') {
                var marginLeft = parseInt(autoLabelWidth, 10) - this.computedWidth;
                if (marginLeft) {
                    style.marginLeft = marginLeft + 'px';
                }
            }
            return h(
                'div',
                { 'class': 'el-form-item__label-wrap', style: style },
                [slots]
            );
        } else {
            return slots[0];
        }
    },

    methods: {
        getLabelWidth() {
            if (this.$el && this.$el.firstElementChild) {
                const computedWidth = window.getComputedStyle(this.$el.firstElementChild).width;
                return Math.ceil(parseFloat(computedWidth));
            } else {
                return 0;
            }
        },
        updateLabelWidth(action = 'update') {
            if (this.$slots.default && this.isAutoWidth && this.$el.firstElementChild) {
                if (action === 'update') {
                    this.computedWidth = this.getLabelWidth();
                } else if (action === 'remove') {
                    this.elForm.deregisterLabelWidth(this.computedWidth);
                }
            }
        }
    },

    watch: {
        computedWidth(val, oldVal) {
            if (this.updateAll) {
                this.elForm.registerLabelWidth(val, oldVal);
                this.elFormItem.updateComputedLabelWidth(val);
            }
        }
    },

    data() {
        return {
            computedWidth: 0
        };
    },

    mounted() {
        this.updateLabelWidth('update');
    },

    updated() {
        this.updateLabelWidth('update');
    },

    beforeDestroy() {
        this.updateLabelWidth('remove');
    }
};