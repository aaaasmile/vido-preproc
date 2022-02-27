<template>
  <div 
    class="ui message transition" 
    :class="{ hidden: isHidden }"
  >
    <i class="close icon" @click="closeToast"></i>
    <div class="header">Message</div>
    <p>
      {{ textMsg }}
    </p>
  </div>
</template>