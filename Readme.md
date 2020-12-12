# vido-preproc
Lo scopo di questa piccola utilità è quello di avere un pre-processor,
per il sito dell'invido, che crei le pagine indexxxx.page da sottoporre poi al
processo di webgen.

Webgen usa i file indexxx.page ma non i file singoli .post. Il problema 
principale di webgen è che non fa il merge dei post singoli nei files
indexxxx.page usando la cronologia. In precedenza editavo direttamente i post
nel file index.page spostando, di tanto in tanto, i post più vecchi nelle 
altra pagine di indice.

Un'altra funzione di vido-preproc è quella di avere un editor con il correttore ortografico,
in questo caso il browser (Visual Studio Code e Notepad++ non l'hanno così buono).

vido-preproc ha una funzione preview senza dover lanciare webgen e ricaricare la pagina di output.

vido-preproc è in grado di lanciare webgen senza dovere aprire una powershell e trovare
la versione corretta di ruby e webgen. Il file config.toml contiene tutte le info.

Per aggiornare vido-preproc uso lo script update-vidopre.ps1.

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
visualizza l'editor usando http get. Il rendering usa i template in go per creare
la pagina html ma anche per creare l'ossatura dei file di .post. 
Data la natura molto semplice dell'editor, non c'è bisogno di un framework in js, tipo reactjs,
ma solo delle librerie semanticui per la grafica e jquery per i comadi http.
Il rendering della preview avviene con il file in javascript textile.js.

Il comando usato durante lo sviluppo è (applicazione già aperta nel browser, altrimenti non mettere il flag --nobrowser)
go run .\main.go --uicmd last --nobrowser

## Comandi
I comandi vengono lanciati all'applicazione via post. Jquery serve per lanciare il post.
Semanticui è per avere un po' di grafica nell'editor.
custom.js serve per creare i popup dei pulsanti e interpretare il click.
La struttura di un comando usa la url do? con il comando nel parametro della request.

## Aggiungere un nuovo post
Si lancia:
.\vido-preproc.exe --title "Bel titolo del nuovo post"
Questo lancia nel browser l'editor con preparato il nuovo post con data e titolo.
Una volta editato il post, lo si salva nel nuovo file di testo sotto la dir post-src
- Poi si ricrea le varie pagine di iondex.html usando dal browser "Create Index  Pages"
- Ora si lancia webgen con il comando nel browser "Start Webgen"

## Aggiungere la repository a Github
Quando si crea una nuova repository su github viene spiegato come si aggiunge questa 
già inizializzata repository a github. Comunque i comandi sono due:
git remote add origin https://github.com/aaaasmile/vido-preproc.git
git push -u origin master

## TODO 
- Prima di sottoporre il preview a textile, si potrebbe risolvere {relocatable} per avere le immagini
nella preview.
- Editare altri post usando il titolo del file invece di editare sempre l'ultimo
- RedCloth redering in browser live (vedi sito RedCloth) e non dopo webgen.
RedCloth usa una soluzione server in ruby che non è utilizzabile. In JS direttamente 
posso usare https://github.com/borgar/textile-js. [DONE]
- Nella uicmd va implementato il comando save. [DONE]

## Comandi
Per vedere un po' il programma in azione si può usare:
vido-preproc -uicmd last

- Crea un nuovo post e lo edita nel browser a http://localhost:4200
vido-preproc --uicmd new --title "Ripreso a programmare la Cuperativa"

- Crea i files index.page con menu di navigazione nella page:
vido-preproc --cmd  createindex

Lo step successivo è quello di lanciare webgen per creare html.

## tags
git tag -a v0.3.20181229-00 -m "0.3.20181229-00 realese"
git push --tags
