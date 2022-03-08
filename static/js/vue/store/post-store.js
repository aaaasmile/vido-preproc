export default {
    state: {
        title_post : ''
    },
    mutations: {
        setPostTitle(state, newval){
            state.title_post = newval
        }
    }        
}