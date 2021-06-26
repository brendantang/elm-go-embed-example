module Pages.Home_ exposing (view)

import Gen.Route
import Html
import Html.Attributes
import View exposing (View)


view : View msg
view =
    { title = "Homepage"
    , body =
        [ Html.text "Hello, world from Elm Spa!"
        , Html.a
            [ Html.Attributes.href (Gen.Route.toHref Gen.Route.Other)
            ]
            [ Html.text "go to other page" ]
        ]
    }
