# vido-preproc
Lo scopo di questa piccola utilità è quello di avere un pre-processor,
per il sito dell'invido, che crei le pagine indexxxx.page da sottoporre poi al
processo di [Webgen](https://github.com/gettalong/webgen). Il sito [invido.it](https://invido.it) viene infatti
generato con generatore statico di HTML.  

![alt text](https://github.com/aaaasmile/vido-preproc/blob/master/doc/12-12-_2020_18-51-18.png?raw=true)

Questo programma serve per aggiornare i sorgenti del sito [invido su github](https://github.com/aaaasmile/InvidositeHtmlgit).

Webgen usa per il suo processing i files con un suffisso .page.
Nel sito ho messo diversi post raggruppati in pagine tipo indexxx.page. Un problema affiora quando
i post cominciano a crescere, ma le pagine indexx.page non cambiano.
Se usassi per ogni singolo post una page dedicata non avrei il raggruppamento di più post su una 
pagina singola.  
Magari è possibile risolvere il problema creando un'estension per Webgen.
Ho fatto però prima a scrivere un pre-processor, che mi risolve anche il problema del
correttore ortografico.
 
In precedenza editavo direttamente i post
nel file index.page spostando, di tanto in tanto, i post più vecchi nelle 
altra pagine di indice.

Un'altra funzione di vido-preproc è quella di avere un editor con il correttore ortografico,
in questo caso il browser (Visual Studio Code e Notepad++ non l'hanno così buono).

vido-preproc ha una funzione preview senza dover lanciare webgen e ricaricare la pagina di output.

vido-preproc è in grado di lanciare webgen senza dovere aprire una powershell e trovare
la versione corretta di ruby e webgen. Il file config.toml contiene tutte le info.

Per aggiornare vido-preproc uso lo script .\start_publish.ps1.

## Come funziona

I singoli post vengono elaborati e suddivisi in diversi .post files.
Ogni file .post contiene un singolo articolo.

Un file .post ha un pattern ben definito con titolo e data
e rispecchia la struttura attuale. Il nome del file contiene la data rovesciata 
per il sort.

Una volta che i post sono suddivisi in file singoli, vido-preproc effettua
un merge che genera il file indexxxx.page finale. Anche il menu di navigazione viene aggiustato.

Alla fine si deve lanciare webgen che crea il sito. L'aggiornamento in remoto avviene con rsynch.

## Sviluppo
vido-preproc è scritto in go ed è una web app che gestisce comandi lanciati via http post e 
visualizza l'editor usando http get.  
Il rendering usa i template in go per creare
la pagina html ma anche per creare l'ossatura dei file di .post. 
Data la natura molto semplice dell'editor, non c'è bisogno di un framework in js, tipo reactjs,
ma solo delle librerie semanticui per la grafica e jquery per i comadi http.  
Il rendering della preview avviene con il file in javascript textile.js.

Il comando usato durante lo sviluppo è (applicazione già aperta nel browser, altrimenti non mettere il flag --nobrowser)  
```go run .\main.go --uicmd last --nobrowser ```

## Comandi
I comandi vengono lanciati all'applicazione via post. Jquery serve per lanciare il post.
Semanticui è per avere un po' di grafica nell'editor.
custom.js serve per creare i popup dei pulsanti e interpretare il click.
La struttura di un comando usa la url do? con il comando nel parametro della request.

## Aggiungere un nuovo post
Si lancia:  
```.\vido-preproc.exe --title "Bel titolo del nuovo post" ```

Questo lancia nel browser l'editor con preparato il nuovo post con data e titolo.
Una volta editato il post, lo si salva nel nuovo file di testo sotto la dir post-src
- Poi si ricrea le varie pagine di index.html usando dal browser "Create Index  Pages"
- Ora si lancia webgen con il comando nel browser "Start Webgen"


### TODO 
- Prima di sottoporre il preview a textile, si potrebbe risolvere {relocatable} per avere le immagini
nella preview.
- Editare altri post e pagine usando aprendo direttamente il file .post o .page
- RedCloth redering in browser live (vedi sito RedCloth) e non dopo webgen.
RedCloth usa una soluzione server in ruby che non è utilizzabile. In JS direttamente 
posso usare https://github.com/borgar/textile-js. [DONE]
- Nella uicmd va implementato il comando save. [DONE]

## Comandi
Per vedere un po' il programma in azione si può usare:  
```vido-preproc -uicmd last```

- Crea un nuovo post e lo edita nel browser a http://localhost:4200   
```vido-preproc --uicmd new --title "Ripreso a programmare la Cuperativa"```

- Crea i files index.page con menu di navigazione nella page:
```vido-preproc --cmd  createindex```

Lo step successivo è quello di lanciare webgen per creare html.


## Element-UI
Per aggiungere ElementUI ho usato il video qui: https://morioh.com/p/83da739c9c76
Si tratta di creare un hello world con vue cli 3.0 e poi creare la distribuzione

    npm install @vue/cli@3.0
    node_modules/.bin/vue create hello-world
    cd hello-world/
    ../node_modules/.bin/vue add element
    npm run serve
    npm run build
Ora nella dir dist si trovano i css e js dei plugin vendor, nel mio caso solo element.
Questo sopra per usare la mia vue 2.0

### Integrazione nella App
Non ho trovato di meglio che prendere i sorgenti dei packages dal node_module e metterli nella sottodirectory element-ui
Sorgenti che vanno splittati in vue e js.
I singoli componenti li ho poi messi in main.js. Molto lavoro manuale, ma almeno c'è il vantaggio che metto solo i componenti che 
mi servono.
Per le immagini delle icone vedi: https://element.eleme.io/#/en-US/component/icon#icon
Per esempio voglio il controllo el-input.
1) Vado nella dir reference\hello-world\node_modules\element-ui\packages e prendo la dir input
e la copio nella dir vido-preproc\vido-preproc\static\js\ext\element-ui\pck
2) cambio il file index.js in element-ui\pck\input
3) dal file input.vue creo manualmente nella stessa directory il file input.js per "Vue Template Js"
In input.js trasporto la sezione Script e creo la sezione template con la string vuota. Poi in input.vue applico "Vue Template Js"
4) I vari import vanno adattati per il browser (path e .js alla fine)
5) Input va messo nella lista dei controlli in main.js
Nota che la variabile process.env.NODE_ENV diventa window.env.NODE_ENV 

