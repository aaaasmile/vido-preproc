$('.message .close')
  .on('click', function() {
    $(this)
      .closest('.message')
      //.fadeToggle() // Nota questo non è quello che si vede in semnati ui (.transition("fade")). Ma è la funzione presa da jquery senza usare semantic-2.4-2.min.js
      .transition('fade') // per usare questo fade, bisogna includere anche semantic-2.4-2.min.js (270k di roba, solo per questa funzione!!)
    ;
  })
;