<template>
  <div
    class="ui message"
    :hidden="snackbar"
  >
    <i class="close icon" @click="closeToast"></i>
    <div class="header">Message</div>
    <p>
      {{ textMsg }}
    </p>
  </div>
</template>