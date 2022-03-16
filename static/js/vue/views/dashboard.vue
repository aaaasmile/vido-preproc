<template>
  <el-container>
    <el-col :span="24">
      <el-row>
        <el-col :span="12">
          <el-row>Editing Post</el-row>
          <el-form ref="form" label-width="80px">
            <el-form-item label="Title Post">
              <el-input v-model="TitlePost"></el-input>
            </el-form-item>
            <el-form-item label="Post">
              <el-input type="textarea" v-model="ContentPost"></el-input>
            </el-form-item>
          </el-form>
        </el-col>
        <el-col :span="12">
          <el-header>Preview Area</el-header>
          <div id="preview"></div>
        </el-col>
      </el-row>
      <el-row>
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
      </el-row>
      <el-row>
        <div
          v-if="LastMsgText"
          name="preprocessor"
          class="ui message transition"
        >
          <i id="preproc-close" class="close icon"></i>

          <div class="header">Preprocessor</div>
          <p>{{ LastMsgText }}</p>
        </div>
      </el-row>
      <el-row>
        Il post si edita con il formato
        <a href="http://borgar.github.com/textile-js/" target="_blank"
          >textile</a
        >
        per il render (in webgen è RedCloth). I link sono del formato
        <code>"testo":http://url</code>. Nota che le apici sono
        <b>SOLO</b> attorno alla prima parola, quella prima dei due punti.
      </el-row>
      <el-row>
        Le immagini messe in src/images si referenziano nel post usando il
        seguente comando:
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
      </el-row>
      <el-row>
        Una volta salvato il post, bisogna ricreare gli indici con il comando
        <i>Create Index Pages</i>
      </el-row>
      <el-row>
        Poi con Webgen si crea il sito completo che va poi sincronizzato con WLC
        (comando ./sync_site_invido.sh)
      </el-row>
    </el-col>
  </el-container>
</template>