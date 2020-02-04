#include<stdio.h>
#include "../include/main.h"
#include"../include/operations.h"
int main()
{
	float op1,op2;
	int ch;
	do{
	printf("Enter 1 No:\n");
	scanf("%f",&op1);
	printf("Enter No 2:\n");
	scanf("%f",&op2);
	printf("1.Addition\n2.Subtracton\n3.multiplication\n4.Division\n5.Quit\n");
	printf("Enter one of The Operation");
	scanf("%d",&ch);
	printf("%d",ch);
	switch(ch)
	{
		case 1:
				printf("Addition is: %f\n",addition(op1,op2));
				break;
		case 2:

				printf("Subtraction is: %f\n",subtraction(op1,op2));
				break;
		case 3:

				printf("Multiplication is: %f\n",multiplication(op1,op2));
				break;
		case 4:

				printf("Division is: %f\n",division(op1,op2));
				break;
		case 5:
				exit(0);

			

	}
	}while(ch!=5);

}
