module Main exposing (..)

import Browser
import Html exposing (Html, div, text, header, h1, p, input, button)
import Html.Attributes exposing (style, placeholder, value, class)
import Html.Events exposing (onInput, onClick)
import Http
import Random
import Dict exposing (Dict)
import Json.Decode as Decode exposing (Decoder)

-- 1. Le Modèle
type alias Meaning =
    { partOfSpeech : String
    , definitions : List Definition
    }

type alias Definition = String

type Model
    = Chargement
    | Erreur String
    | Succes 
        { tousLesMots : List String
        , motChoisi : String
        , meanings : List Meaning
        , saisie : String 
        , montrerSolution : Bool 
        }

-- 2. L'initialisation
init : () -> ( Model, Cmd Msg )
init _ =
    ( Chargement
    , Http.get 
        { url = "mots.txt"
        , expect = Http.expectString ReçuLeFichier
        }
    )

-- 3. Les Messages
type Msg
    = ReçuLeFichier (Result Http.Error String)
    | IndexChoisi Int
    | DefinitionsRecues (Result Http.Error (List Meaning))
    | ChangementSaisie String
    | AfficherSolution
    | ProchainMot

-- Logique de traitement des données
partOfSpeechPriority : String -> Int
partOfSpeechPriority pos =
    case pos of
        "noun" -> 0
        "verb" -> 1
        "proper noun" -> 2
        _ -> 99

insertMeaning : Meaning -> Dict String Meaning -> Dict String Meaning
insertMeaning meaning dict =
    Dict.update meaning.partOfSpeech
        (\maybeExisting ->
            case maybeExisting of
                Nothing -> Just meaning
                Just existing -> Just { existing | definitions = existing.definitions ++ meaning.definitions }
        )
        dict

groupMeanings : List Meaning -> List Meaning
groupMeanings meanings =
    meanings
        |> List.foldl insertMeaning Dict.empty
        |> Dict.values
        |> List.sortBy (\m -> partOfSpeechPriority m.partOfSpeech)

-- Decoders
definitionDecoder : Decoder String
definitionDecoder = Decode.field "definition" Decode.string

definitionsDecoder : Decoder (List String)
definitionsDecoder = Decode.field "definitions" (Decode.list definitionDecoder)

meaningsDecoder : Decoder Meaning
meaningsDecoder = Decode.map2 Meaning (Decode.field "partOfSpeech" Decode.string) definitionsDecoder 

dictionaryDecoder : Decoder (List Meaning)
dictionaryDecoder = Decode.list (Decode.field "meanings" (Decode.list meaningsDecoder)) |> Decode.map List.concat

-- Styles
stylesAnimation : Html Msg
stylesAnimation =
    Html.node "style" []
        [ text """
            @keyframes blink-green {
                0%, 100% { background-color: #6aaa64; color: white; border-color: #6aaa64; }
                50% { background-color: white; color: #121213; border-color: #d3d6da; }
            }
            .wordle-success {
                animation: blink-green 0.6s ease-in-out infinite;
                background-color: #6aaa64 !important;
                color: white !important;
                border-color: #6aaa64 !important;
            }
        """ ]

-- 4. L'Update
update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        ReçuLeFichier resultat ->
            case resultat of
                Ok contenu ->
                    let 
                        listeMots = String.words contenu
                        taille = List.length listeMots - 1
                    in
                    ( Succes 
                        { tousLesMots = listeMots
                        , motChoisi = ""
                        , meanings = []
                        , saisie = ""
                        , montrerSolution = False 
                        }
                    , Random.generate IndexChoisi (Random.int 0 taille)
                    )
                Err _ ->
                    ( Erreur "Impossible de lire le fichier .txt", Cmd.none )

        IndexChoisi index ->
            case model of
                Succes contenu ->
                    let 
                        mot = List.drop index contenu.tousLesMots |> List.head |> Maybe.withDefault ""
                        url = "https://api.dictionaryapi.dev/api/v2/entries/en/" ++ mot
                    in 
                    ( Succes { contenu | motChoisi = mot, meanings = [], saisie = "", montrerSolution = False }
                    , Http.get { url = url, expect = Http.expectJson DefinitionsRecues dictionaryDecoder }
                    )
                _ -> ( model, Cmd.none )

        DefinitionsRecues resultat ->
            case (resultat, model) of
                ( Ok meanings, Succes data ) ->
                    ( Succes { data | meanings = meanings }, Cmd.none )
                ( Err err, _ ) ->
                    ( Erreur ("Erreur dictionnaire : " ++ Debug.toString err), Cmd.none )
                _ -> ( model, Cmd.none )

        ChangementSaisie nouvelleSaisie ->
            case model of
                Succes data -> ( Succes { data | saisie = nouvelleSaisie }, Cmd.none )
                _ -> ( model, Cmd.none )

        AfficherSolution ->
            case model of
                Succes data -> ( Succes { data | montrerSolution = True }, Cmd.none )
                _ -> ( model, Cmd.none )

        ProchainMot ->
            case model of
                Succes data ->
                    let
                        taille = List.length data.tousLesMots - 1
                    in
                    ( model, Random.generate IndexChoisi (Random.int 0 taille) )
                _ -> ( model, Cmd.none )

-- 5. La Vue
viewMeaning : Meaning -> Html Msg
viewMeaning meaning =
    div [ style "margin-bottom" "24px" ]
        [ div 
            [ style "color" "#818384"
            , style "text-transform" "uppercase"
            , style "font-weight" "bold"
            , style "font-size" "14px"
            , style "margin-bottom" "8px"
            ] 
            [ text meaning.partOfSpeech ]
        , div [ style "color" "#121213", style "line-height" "1.5" ] 
            (List.map (\def -> p [ style "margin" "4px 0" ] [ text ("• " ++ def) ]) meaning.definitions)
        ]

view : Model -> Html Msg
view model =
    case model of
        Chargement ->
            div [ style "padding" "20px" ] [ text "Chargement du dictionnaire..." ]

        Erreur message ->
            div [ style "padding" "20px", style "color" "red" ] [ text message ]

        Succes data ->
            let
                estJuste = String.toLower (String.trim data.saisie) == String.toLower data.motChoisi
                
                titreTexte = 
                    if data.montrerSolution then 
                        "THE WORD WAS: " ++ String.toUpper data.motChoisi 
                    else 
                        "GUESS IT !"
            in
            div 
                [ style "background-color" "white", style "min-height" "100vh"
                , style "display" "flex", style "flex-direction" "column", style "align-items" "center"
                , style "font-family" "'Helvetica Neue', Arial, sans-serif"
                ]
                [ stylesAnimation
                , header 
                    [ style "width" "100%", style "background-color" "white", style "color" "#121213"
                    , style "border-bottom" "1px solid #d3d6da", style "text-align" "center", style "padding" "15px 0"
                    , style "margin-bottom" "30px"
                    ]
                    [ h1 [ style "margin" "0", style "font-size" "32px" ] [ text titreTexte ]
                    
                    , div [ style "margin-top" "10px" ] 
                        [ if not data.montrerSolution && not estJuste then
                            button 
                                [ onClick AfficherSolution
                                , style "background" "none", style "border" "none", style "color" "#818384"
                                , style "text-decoration" "underline", style "cursor" "pointer", style "margin-right" "15px"
                                ] 
                                [ text "Give up? Show answer" ]
                          else text ""
                        
                        , button 
                            [ onClick ProchainMot
                            , style "background-color" "#121213", style "color" "white", style "border" "none"
                            , style "padding" "8px 16px", style "border-radius" "4px", style "cursor" "pointer", style "font-weight" "bold"
                            ] 
                            [ text "NEW WORD" ]
                        ]
                    ]

                , div [ style "max-width" "800px", style "width" "90%" ]
                    [ div [ style "margin-bottom" "30px" ] 
                        (List.map viewMeaning data.meanings)
                    
                    , input 
                        [ placeholder "Type your guess..."
                        , value data.saisie
                        , onInput ChangementSaisie 
                        , class (if estJuste then "wordle-success" else "")
                        , style "background-color" "white", style "border" "2px solid #d3d6da"
                        , style "color" "#121213", style "padding" "15px", style "width" "100%"
                        , style "font-size" "1.2rem", style "text-align" "center", style "text-transform" "uppercase"
                        , style "box-sizing" "border-box", style "outline" "none"
                        ]
                        []
                    ]
                ]

main =
    Browser.element
        { init = init
        , update = update
        , subscriptions = \_ -> Sub.none
        , view = view
        }