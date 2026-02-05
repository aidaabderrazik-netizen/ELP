
module Model exposing (Model(..), Meaning, Definition, GameMode(..))

-- 1. Le Modèle
type alias Meaning =
    { partOfSpeech : String
    , definitions : List Definition
    }

type GameMode = Idle | Playing

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
        , score : Int
        , tempsRestant : Int
        , mode : GameMode
        }

