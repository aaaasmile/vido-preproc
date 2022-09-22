<template>
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
  </v-row>
</template>