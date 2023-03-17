build:
	go build -o Imagenes main.go
run:
	./Imagenes
clearR:
	rm -R ./imgs/rojo
	@echo "Directorio rojo eliminado"
clearG:
	rm -R ./imgs/verde
	@echo "Directorio verde eliminado"
clearB:
	rm -R ./imgs/azul
	@echo "Directorio azul eliminado"
clearAll:
	make clearR
	make clearG
	make clearB
	@echo "Directorios eliminados"
