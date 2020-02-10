#include<stdio.h>
#include<stdlib.h>
#include<string.h>
#include"../include/evaluate.h"
#include"../include/check_exp.h"
int main()
{
	int res;
	char exp[100];
	while(1)	
	{
	printf("\nenter expression or press \\q to quit\n");
	gets(exp);
	if(strcmp(exp,"\\q")==0)
	{
		exit(1);
	}
	/*check given expression valid or not */
	res=check_exp(exp);
	if(res==1)
	{
		/*if valid display result*/
		printf("==>Result is%f\n",evaluate(exp));
	}
	else
	{
		/* if invalid display error message*/
		printf("==>Invalid expression\n");
	}
	}
	
	
	
}
