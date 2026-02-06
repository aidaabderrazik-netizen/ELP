module View exposing (view)

import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Model exposing (Model(..), Meaning, GameMode(..))
import Update exposing(Msg(..))

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


-- 5. La Vue
viewMeaning : Meaning -> Html Msg
viewMeaning meaning =
    div [ style "margin-bottom" "24px" ]
        [ div [ style "color" "#818384", style "text-transform" "uppercase", style "font-weight" "bold", style "font-size" "14px", style "margin-bottom" "8px" ] [ text meaning.partOfSpeech ]
        , div [ style "color" "#121213", style "line-height" "1.5" ] (List.map (\def -> p [ style "margin" "4px 0" ] [ text ("• " ++ def) ]) meaning.definitions)
        ]

view : Model -> Html Msg
view model =
    case model of
        Chargement -> div [ style "padding" "20px" ] [ text "Loading dictionary..." ]
        Erreur message -> div [ style "padding" "20px", style "color" "red" ] [ text message ]
        Succes data ->
            let
                estJuste = String.toLower (String.trim data.saisie) == String.toLower data.motChoisi
                formatTime s = (String.fromInt (s // 60)) ++ ":" ++ (String.padLeft 2 '0' (String.fromInt (remainderBy 60 s)))
                titreTexte = if data.montrerSolution then "THE WORD WAS: " ++ String.toUpper data.motChoisi else "GUESS IT !"
            in
            div [ style "background-color" "white", style "min-height" "100vh", style "display" "flex", style "flex-direction" "column", style "align-items" "center", style "font-family" "sans-serif" ]
                [ stylesAnimation
                , header [ style "width" "100%", style "border-bottom" "1px solid #d3d6da", style "text-align" "center", style "padding" "15px 0", style "margin-bottom" "30px" ]
                    [ h1 [ style "margin" "0", style "font-size" "32px" ] [ text titreTexte ]
                    , div [ style "font-size" "20px", style "margin" "10px", style "font-weight" "bold" ] 
                        [ text ("TIME: " ++ formatTime data.tempsRestant)
                        , span [ style "margin-left" "20px", style "color" "#6aaa64" ] [ text ("SCORE: " ++ String.fromInt data.score) ]
                        ]
                    , if data.mode == Idle then
                        button [ onClick StartTimerGame, style "background-color" "#6aaa64", style "color" "white", style "border" "none", style "padding" "10px 20px", style "border-radius" "4px", style "cursor" "pointer", style "font-weight" "bold" ] [ text "START CHALLENGE" ]
                      else text ""
                    ]
                , div [ style "max-width" "800px", style "width" "90%" ]
                    [ div [ style "margin-bottom" "30px" ] (List.map viewMeaning data.meanings)
                    , input 
                        [ placeholder "Type your guess..."
                        , value data.saisie
                        , onInput ChangementSaisie 
                        , class (if estJuste then "wordle-success" else "")
                        , disabled (data.mode == Idle && data.tempsRestant <= 0)
                        , style "width" "100%", style "padding" "15px", style "font-size" "1.2rem", style "text-align" "center", style "text-transform" "uppercase", style "border" "2px solid #d3d6da"
                        ] []
                    ]
                ]
