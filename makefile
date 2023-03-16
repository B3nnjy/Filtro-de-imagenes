build:
	go build -o Imagenes main.go
run:
	./Imagenes
clearR:
	rm -R ./imgs/rojo
clearG:
	rm -R ./imgs/verde
clearB:
	rm -R ./imgs/azul
clearAll:
	make clearR
	make clearG
	make clearB
	clear
	echo "Directorios eliminados"