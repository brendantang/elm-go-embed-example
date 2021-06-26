module Pages.Echo.Foo_ exposing (Model, Msg, page)

import Gen.Params.Echo.Foo_ exposing (Params)
import Html
import Http
import Page
import Request
import Shared
import View exposing (View)


page : Shared.Model -> Request.With Params -> Page.With Model Msg
page shared req =
    Page.element
        { init = init req.params
        , update = update
        , view = view
        , subscriptions = subscriptions
        }



-- INIT


type Model
    = Loading
    | Failure
    | Success String


init : Params -> ( Model, Cmd Msg )
init params =
    ( Loading
    , Http.get
        { url = "/api/echo/" ++ params.foo
        , expect = Http.expectString GotBackendResponse
        }
    )



-- UPDATE


type Msg
    = GotBackendResponse (Result Http.Error String)


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        GotBackendResponse response ->
            case response of
                Ok responseText ->
                    ( Success responseText, Cmd.none )

                Err _ ->
                    ( Failure, Cmd.none )



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none



-- VIEW


view : Model -> View Msg
view model =
    { title = "Homepage"
    , body =
        [ Html.text <|
            case model of
                Loading ->
                    "Loading response from backend..."

                Success responseText ->
                    "Got response from backend:" ++ responseText

                Failure ->
                    "Error getting response from backend"
        ]
    }
