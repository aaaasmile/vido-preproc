export default {
    state: {
        title_post : '',
        content_post: '',
    },
    mutations: {
        setPostTitle(state, newval){
            state.title_post = newval
        },
        setPostContent(state, newval){
            state.content_post = newval
        }
    }        
}