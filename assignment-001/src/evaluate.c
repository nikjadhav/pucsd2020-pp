#include"../include/stackoperation.h"
#include"../include/operations.h"
#include<string.h>
#include<stdio.h>
#include<stdlib.h>
float check_operation(char ch,float opnd1,float opnd2)
{
	switch(ch)
	{
		case '+':
			return addition(opnd1,opnd2);
		case '-':
			return subtraction(opnd1,opnd2);
		case '*':
			return multiplication(opnd1,opnd2);
		case '/':
			return division(opnd1,opnd2);
	}
}

int precedence(char ch)
{
	switch(ch)
	{
		case '@':
			return 0;

		case '+':

		case '-':
			return 1;
		case '*':

		case '/':
			return 2;
		case '^':
			return 3;
		case '(':
			return 4;
	}
}
float evaluate(char infix[])
{
char optr,ch;
char temp[100];
float opnd1,opnd2,value;
STACK1 operand;
STACK2 operator;
int i=0,n;
n=strlen(infix);
ini1(&operand);
ini2(&operator);
push2(&operator,'@');
while(i<n)
{
	if(infix[i]>='0' && infix[i]<='9')
	{
		while(infix[i]>='0' && infix[i]<='9')
		{
			ch=infix[i];
			strncat(temp,&ch,1);
			i=i+1;
		}
		i=i-1;
		push1(&operand,(float)atoi(temp));
		for(int j=0;j<100;j++)
		{
			temp[j]='\0';
		}
	}
	else
	if(infix[i]=='+' || infix[i]=='*' || infix[i]=='-' || infix[i]=='/')
	{
		optr=pop2(&operator);
		if(precedence(infix[i])>precedence(optr))
		{
			push2(&operator,optr);
			push2(&operator,infix[i]);
		}
		else
		{
			while(precedence(infix[i])<=precedence(optr))
			{
				opnd2=pop1(&operand);
				opnd1=pop1(&operand);
				value=check_operation(optr,opnd1,opnd2);
				push1(&operand,value);
				optr=pop2(&operator);
			}
			push2(&operator,optr);
			push2(&operator,infix[i]);

		}
	}
	i=i+1;
}
	while((ch=pop2(&operator))!='@')
        {
                opnd2=pop1(&operand);
                opnd1=pop1(&operand);
                value=check_operation(ch,opnd1,opnd2);
                push1(&operand,value);
        }
        return (pop1(&operand));
}
