module Main exposing (main)

import Browser
import Update exposing (init, update, subscriptions)
import View exposing (view)

main =
    Browser.element
        { init = init
        , update = update
        , view = view
        , subscriptions = subscriptions
        }
