module Update exposing (Msg(..), init, update, subscriptions)

import Http
import Random
import Time
import Model exposing (Model(..), GameMode(..))
import Dictionnary exposing (dictionaryDecoder)

type Msg
    = ReçuLeFichier (Result Http.Error String)
    | IndexChoisi Int
    | DefinitionsRecues (Result Http.Error (List Model.Meaning))
    | ChangementSaisie String
    | AfficherSolution
    | ProchainMot
    | StartTimerGame
    | Tick Time.Posix


-- 2. L'initialisation
init : () -> ( Model, Cmd Msg )
init _ =
    ( Chargement
    , Http.get 
        { url = "mots.txt"
        , expect = Http.expectString ReçuLeFichier
        }
    )

update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        ReçuLeFichier resultat ->
            case resultat of
                Ok contenu ->
                    let 
                        listeMots = String.words contenu
                    in
                    ( Succes 
                        { tousLesMots = listeMots
                        , motChoisi = ""
                        , meanings = []
                        , saisie = ""
                        , montrerSolution = False 
                        , score = 0 
                        , tempsRestant = 120
                        , mode = Idle
                        }
                    , Random.generate IndexChoisi (Random.int 0 (List.length listeMots - 1))
                    )
                Err _ -> ( Erreur "Impossible de lire le fichier .txt", Cmd.none )

        IndexChoisi index ->
            case model of
                Succes data ->
                    let 
                        mot = List.drop index data.tousLesMots |> List.head |> Maybe.withDefault ""
                    in 
                    ( Succes { data | motChoisi = mot, meanings = [], saisie = "", montrerSolution = False }
                    , Http.get { url = "https://api.dictionaryapi.dev/api/v2/entries/en/" ++ mot, expect = Http.expectJson DefinitionsRecues dictionaryDecoder }
                    )
                _ -> ( model, Cmd.none )

        DefinitionsRecues resultat ->
            case (resultat, model) of
                ( Ok meanings, Succes data ) -> ( Succes { data | meanings = meanings }, Cmd.none )
                _ -> ( model, Cmd.none )

        StartTimerGame ->
            case model of
                Succes data -> 
                    ( Succes { data | mode = Playing, score = 0, tempsRestant = 120, montrerSolution = False }
                    -- On demande un nouvel index au hasard dès qu'on appuie sur Start
                    , Random.generate IndexChoisi (Random.int 0 (List.length data.tousLesMots - 1))
                    )
                _ -> ( model, Cmd.none )
        Tick _ ->
            case model of
                Succes data ->
                    if data.mode == Playing then
                        if data.tempsRestant <= 0 then
                            ( Succes { data | mode = Idle, montrerSolution = True }, Cmd.none )
                        else
                            ( Succes { data | tempsRestant = data.tempsRestant - 1 }, Cmd.none )
                    else ( model, Cmd.none )
                _ -> ( model, Cmd.none )

        ChangementSaisie nouvelleSaisie ->
            case model of
                Succes data -> 
                    let 
                        estJuste = String.toLower (String.trim nouvelleSaisie) == String.toLower data.motChoisi
                    in
                    if estJuste && data.mode == Playing then
                        ( Succes { data | score = data.score + 1, saisie = "" }
                        , Random.generate IndexChoisi (Random.int 0 (List.length data.tousLesMots - 1))
                        )
                    else
                        ( Succes { data | saisie = nouvelleSaisie }, Cmd.none )
                _ -> ( model, Cmd.none )

        AfficherSolution ->
            case model of
                Succes data -> ( Succes { data | montrerSolution = True }, Cmd.none )
                _ -> ( model, Cmd.none )

        ProchainMot ->
            case model of
                Succes data -> ( model, Random.generate IndexChoisi (Random.int 0 (List.length data.tousLesMots - 1)) )
                _ -> ( model, Cmd.none )

subscriptions : Model -> Sub Msg
subscriptions model =
    case model of
        Succes data ->
            if data.mode == Playing then
                Time.every 1000 Tick
            else
                Sub.none

        _ ->
            Sub.none
