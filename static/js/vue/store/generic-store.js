export default {
  state: {
    errorText: '',
    msgText: '',
    lastmsgText: '',
    datalog: [],
  },
  mutations: {    
    errorText(state, msg) {
      state.errorText = msg
      state.lastmsgText = msg
    },
    msgText(state, msg) {
      state.msgText = msg
      state.lastmsgText = msg
    },
    msgTextStatus(state, msg) {
      state.msgText = ''
      state.lastmsgText = msg
    },
    lastMsgText(state, msg) {
      state.lastmsgText = msg
    },
    msginfolog(state, datalog) {
      state.datalog = []
      for (let ix = 0; ix < datalog.length; ix++) {
        state.datalog.push({ key: ix, text: datalog[ix] })
      }
    },
    clearErrorText(state) {
      if (state.errorText !== '') {
        state.errorText = ''
      }
    },
    clearMsgText(state) {
      if (state.msgText !== '') {
        state.msgText = ''
      }
    }
  }
}