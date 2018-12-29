applyCloseMsg = (sel) => {
  $(sel)
    .on('click', function () {
      $(this)
        .closest('.message')
        //.fadeToggle() // Nota questo non è quello che si vede in semnati ui (.transition("fade")). Ma è la funzione presa da jquery senza usare semantic-2.4-2.min.js
        .transition('fade') // per usare questo fade, bisogna includere anche semantic-2.4-2.min.js (270k di roba, solo per questa funzione!!)
        ;
      // send also a message to the server to clear the message
      let name = $(this).closest('.message').attr("name");
      if (name === "preprocessor" || name === "feedback") {
        console.log("Clear this: ", name)
        let url = "do?"
        url = url + $.param({ "clear": "preprocessor" })
        $.post(url, res => {
          console.log(res)
        });
      }
    })
}

$('#bttextile')
  .on('click', () => {
    console.log("Preview clicked")
    //console.log( textile( "I am using __textile__." ) );
    //let htmlprev = textile( "I am using __textile__." )
    let rawtext = $('#post').val()
    let htmlprev = textile(rawtext)
    $('#preview').empty().append(htmlprev);
  });

// popup in semnatic ui are not active by default
// activate it only in bttextile. The data-content is expected.
$('#bttextile')
  .popup({
    inline: true
  })
  ;

$('#btgotowebgenout')
  .popup({
    inline: true
  })
  .on('click', () => {
    console.log("Go to webgen output clicked")
    let url = "do?"
    url = url + $.param({ "openwebgenout": '' })
    $.post(url, res => {
      console.log(res)
    })
  })
  ;
$('#btsave')
  .popup({
    inline: true
  })
  .on('click', () => {
    console.log("Save the current post")
    let url = "do?"
    let rawtext = $('#post').val()
    url = url + $.param({ "save": '' })
    $.post(url, rawtext, res => { // rawtext is a data, not a param in url
      let cont = JSON.parse(res) // res is always in Json, also if I send a simply string
      console.log(cont)
      writeFeedback(cont)
    })
  })
  ;

writeFeedback = (x) => {
  let html = `<i id="feedback-close" class="close icon"></i><div class="header">Result</div><p>${x}</p>`
  //console.log(html)
  $('#feedback')
    .empty()
    .removeClass("hidden")
    .append(html)
    ;
  applyCloseMsg('#feedback-close') // reapply because the close icon handler
}


applyCloseMsg('#preproc-close')
console.log('Custom script is ready.')