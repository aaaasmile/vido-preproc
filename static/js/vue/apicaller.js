
export default {
  CallPostEditor(that, req) {
    return that.$http.post("PostEditor", JSON.stringify(req), { headers: { "content-type": "application/json" } })
  },
  ListPost(that, params) {
    let req = { method: 'ListPost', Params: params }
    this.CallPostEditor(that, req).then(result => {
      console.log('Call terminated ', result.data)
      that.$store.commit('msgText', `Status: ${result.data.Status}`)
    }, error => {
      console.error(error);
      that.$store.commit('errorText', error.bodyText)
    });
  },  
}