build:
	gcc src/main.c src/operations/addition.c src/operations/subtraction.c src/operations/multiplication.c src/operations/division.c src/operandstack.c src/operatorstack.c src/check_exp.c src/evaluate.c -o main
build-static:
	gcc  src/main.o -Lsrc/ -lcal -o main
build-shared:
	gcc src/main.c src/libdynamiccal.so -o main
run:
	./main
clean:
	rm main
