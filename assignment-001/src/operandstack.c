/*operand stack to store the operands*/
#include<stdio.h>
#define MAXSIZE 100
typedef struct
	{
	float data[MAXSIZE];
	int top;
	}STACK1;
int isempty1(STACK1 *s)
	{
	if(s->top==-1)
		return 1;
	return 0;
	}

int isfull1(STACK1 *s)
	{
	if(s->top==MAXSIZE-1)
		return 1;
	return 0;
	}

void push1(STACK1 *s,float num)
	{
	s->top++;
	s->data[s->top]=num;
	}

float pop1(STACK1 *s)
	{
	float num;
	num=s->data[s->top];
	s->top--;
	return num;
	}

int peek1(STACK1 *s)
	{
	float num;
	num=s->data[s->top];
	return num;
	}

void ini1(STACK1 *s)
	{
	s->top=-1;
	}
