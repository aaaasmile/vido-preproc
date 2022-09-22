import API from '../apicaller.js?version=1'
import Toast from '../components/toast.js?version=1'

export default {
  components: { Toast },
  data() {
    return {
      CommandImage: "",

    }
  },
  created() {
    this.CommandImage = "<a href=\"{relocatable: /images/nuovaimg.PNG}\"><img width=\"300\" src=\"{relocatable: /images/nuovaimg.PNG}\"></a>"
  },
  mounted() {
  },
  computed: {
    TitlePost: {
      get() {
        return this.$store.state.post.title_post
      },
      set(newVal) {
        this.$store.commit('setPostTitle', newVal)
      }
    },
    ContentPost: {
      get() {
        return this.$store.state.post.content_post
      },
      set(newVal) {
        this.$store.commit('setPostContent', newVal)
      }

    },
    ...Vuex.mapState({
      LastMsgText: state => {
        return state.gen.lastMsgText
      },  
    })
  },
  methods: {
    openFile() {
      console.log('Call open file')
      API.ListPost(this, {})
    },
    saveFile(){
      console.log('Save file')
    },
    previewFile(){
      console.log('Preview file')
    },
    createIndexPages(){
      console.log('create index pages')
    },
    startWebGen(){
      console.log('start web gen')
    },
    navigateToWebGenOut(){
      console.log('navigate to web gen out')
    },
  },
  template: `
  <v-row justify="center">
    <Toast></Toast>
    <h2>Invido Preprocessor</h2>
    <v-col cols="12">
      <v-card>
        <v-card-title>Editing Post</v-card-title>
        <v-expansion-panels :flat="true">
          <v-expansion-panel>
            <v-expansion-panel-header></v-expansion-panel-header>
            <v-expansion-panel-content>
              <v-container>
                <v-row>
                  <v-col cols="12" md="4">
                    <v-text-field
                      v-model="TitlePost"
                      label="Title Post"
                    ></v-text-field>
                  </v-col>
                </v-row>
              </v-container>
              <v-card-actions>
                <v-btn v-on:click="openFile">Open Folder</v-btn>
              </v-card-actions>
            </v-expansion-panel-content>
          </v-expansion-panel>
        </v-expansion-panels>
      </v-card>
    </v-col>
  </v-row>`
}