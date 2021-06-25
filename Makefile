all: frontend server

frontend: 
	cd elm-frontend && elm-spa build 

server:
	go build -o dist
