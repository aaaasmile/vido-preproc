import Generic from './generic-store.js'
import Post from './post-store.js'

export default new Vuex.Store({
  modules: {
    gen: Generic,
    post: Post,
  }
})