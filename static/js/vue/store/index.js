import Generic from './generic-store.js?version=1'
import Post from './post-store.js?version=1'

export default new Vuex.Store({
  modules: {
    gen: Generic,
    post: Post,
  }
})