export default {
  state: {
    errorText: '',
    msgText: '',
    navbackend: '',
    navbackend_system: '',
    lastmsgText: '',
    infodata: [],
    reslog: [],
    user_name: '',
    company: '',
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
    msginfolog(state, infodata) {
      state.infodata = []
      for (let ix = 0; ix < infodata.length; ix++) {
        state.infodata.push({ key: ix, text: infodata[ix] })
      }
    },
    clearErrorText(state) {
      if (state.errorText !== '') {
        state.errorText = ''
      }
    },
    clearMsgText(state) {
      state.msgText = ''
    },
    clearAll(state) {
      state.msgText = ''
      state.errorText = ''
      state.infodata = [],
      state.reslog = [],
      state.lastmsgText = ''
    },
    scriptResDatalog(state, datalog) {
      state.reslog = []
      for (let ix = 0; ix < datalog.length; ix++) {
        state.reslog.push({ key: ix, text: datalog[ix] })
      }
    }
  }
}