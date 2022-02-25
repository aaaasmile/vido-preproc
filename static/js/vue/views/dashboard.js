export default {
    data() {
        return {
            CommandImage: "",

        }
    },
    created() {
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
            TitlePost:state => {

            }
        })
    },
    methods: {
        echo() {
            console.log('Call echo')
        },
    },
    template: `
  <div class="ui container">
    <div class="ui attached message">
      <div class="header">Editing Post</div>
    </div>
    <form
      class="ui form attached fluid segment"
      action="/save-post/"
      method="POST"
    >
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
            >{{ ContentPost }}</textarea
          >
        </div>
        <div class="field">
          <label>Preview</label>
          <div id="preview"></div>
        </div>
      </div>
      <div>
        <!-- <input class="ui blue submit button" type="submit" value="Save" > -->
        <!-- Save with page reload -->
        <a class="ui button" id="btsave" data-content="Save the current post"
          ><i class="save outline icon"></i
        ></a>
        <!-- Save without page reload -->
        <a class="ui button" id="bttextile" data-content="Preview"
          ><i class="code icon"></i
        ></a>
        <a
          class="ui button"
          id="btbuildindex"
          data-content="Create all index-00-99.page files"
          >Create Index Pages</a
        >
        <a
          class="ui button"
          id="btrunwebgen"
          data-content="Run webgen to update the full site"
          >Start Webgen</a
        >
        <a
          class="ui button"
          id="btgotowebgenout"
          data-content="Open webgen output in a new browser window"
          >Navigate to webgen out</a
        >
      </div>
    </form>

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
  </div>
`
}