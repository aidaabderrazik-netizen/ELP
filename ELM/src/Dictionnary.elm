module Dictionnary exposing (dictionaryDecoder, groupMeanings)

import Dict exposing (Dict)
import Model exposing (Meaning)
import Json.Decode as Decode exposing (Decoder)

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
dictionaryDecoder = Decode.list (Decode.field "meanings" (Decode.list meaningsDecoder)) |> Decode.map (List.concat >> groupMeanings)
