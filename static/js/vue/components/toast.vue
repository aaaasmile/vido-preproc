<template>
  <v-snackbar v-model="snackbar" :color="colorsnack">
      {{ textMsg }}
      <v-btn dark text @click="closeToast">Close</v-btn>
    </v-snackbar>
</template>