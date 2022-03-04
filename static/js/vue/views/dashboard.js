import API from '../apicaller.js'

export default {
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
    ...Vuex.mapState({
      LastMsgText: state => {
        return state.gen.lastMsgText
      },
      ContentPost: state => {

      },
      TitlePost: state => {

      }
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
  <div class="ui container">
    <div class="ui attached message">
      <div class="header">Editing Post</div>
    </div>
    <div class="ui form attached fluid segment">
      <div class="field">
        <label>File name</label>
        <input
          placeholder="TitlePost"
          disabled="true"
          name="titlepost"
          type="text"
          size="35"
          :value="TitlePost"
        />
      </div>
      <div class="two fields">
        <div class="field">
          <label>Post</label>
          <textarea
            id="post"
            rows="25"
            cols="110"
            name="contentpost"
            placeholder="ContentPost"
            :value="ContentPost"
          ></textarea>
        </div>
        <div class="field">
          <label>Preview Area</label>
          <div id="preview"></div>
        </div>
      </div>
      <div>
        <el-tooltip
          class="item"
          effect="dark"
          content="Open Folder"
          placement="top-start"
        >
          <el-button @click="openFile" icon="el-icon-folder-opened"></el-button>
        </el-tooltip>
        <el-tooltip
          class="item"
          effect="dark"
          content="Save file"
          placement="top-start"
        >
          <el-button
            @click="saveFile"
            icon="el-icon-folder-checked"
          ></el-button>
        </el-tooltip>
        <el-tooltip
          class="item"
          effect="dark"
          content="Preview"
          placement="top-start"
        >
          <el-button @click="previewFile" icon="el-icon-view"></el-button>
        </el-tooltip>
        <el-button @click="createIndexPages">Create Index Pages</el-button>
        <el-button @click="startWebGen">Start Webgen</el-button>
        <el-button @click="navigateToWebGenOut"
          >Navigate to webgen out</el-button
        >
      </div>
    </div>

    <div v-if="LastMsgText" name="preprocessor" class="ui message transition">
      <i id="preproc-close" class="close icon"></i>

      <div class="header">Preprocessor</div>
      <p>{{ LastMsgText }}</p>
    </div>

    <div
      name="feedback"
      id="feedback"
      class="ui message transition hidden"
    ></div>

    <div class="ui vertical segment"></div>
    <div class="ui vertical segment">
      <p>
        Il post si edita con il formato
        <a href="http://borgar.github.com/textile-js/" target="_blank"
          >textile</a
        >
        per il render (in webgen è RedCloth). I link sono del formato
        <code>"testo":http://url</code>. Nota che le apici sono
        <b>SOLO</b> attorno alla prima parola, quella prima dei due punti.
      </p>
    </div>
    <div class="ui vertical segment">
      <p>
        Le immagini messe in src/images si referenziano nel post usando il
        seguente comando:
      </p>
      <code> {{ CommandImage }} </code>
      <p>Oppure in Redcloth:</p>
      <p>
        <code
          >!https://github.com/aaaasmile/live-omxctrl/blob/master/doc/05-12-_2020_22-23-43.png?raw=true!:https://github.com/aaaasmile/live-omxctrl/blob/master/doc/05-12-_2020_22-23-43.png?raw=true</code
        >
      </p>
      <p>
        Un riferimento completo di RedCloth si trova su
        <a href="https://github.com/jgarber/redcloth"
          >https://github.com/jgarber/redcloth</a
        >
      </p>
    </div>
    <div class="ui vertical segment">
      <p>
        Una volta salvato il post, bisogna ricreare gli indici con il comando
        <i>Create Index Pages</i>
      </p>
      <p>
        Poi con Webgen si crea il sito completo che va poi sincronizzato con WLC
        (comando ./sync_site_invido.sh)
      </p>
    </div>
  </div>`
}