build:
	gcc  src/main.o -Lsrc/ -lcal -o main
run:
	./main
clean:
	rm main	
