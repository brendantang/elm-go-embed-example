all: frontend server

frontend: 
	cd elm-frontend && elm make src/Main.elm --output build/index.html

server:
	go build -o dist
