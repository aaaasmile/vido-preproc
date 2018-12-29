console.log('Hi, custom is loaded.')
$('.message .close')
  .on('click', function () {
    $(this)
      .closest('.message')
      //.fadeToggle() // Nota questo non è quello che si vede in semnati ui (.transition("fade")). Ma è la funzione presa da jquery senza usare semantic-2.4-2.min.js
      .transition('fade') // per usare questo fade, bisogna includere anche semantic-2.4-2.min.js (270k di roba, solo per questa funzione!!)
      ;
  })
  ;

$('#bttextile')
  .on('click', () => {
    console.log("Preview clicked!")
    //console.log( textile( "I am using __textile__." ) );
    //let htmlprev = textile( "I am using __textile__." )
    let rawtext = $('#post').val()
    let htmlprev = textile(rawtext)
    $('#preview').empty().append(htmlprev)
  });

$('#bttextile')
  .popup({
    inline: true
  })
  ;