/*operator stack to store the all operators*/
#include<stdio.h>
#define MAXSIZE 100
typedef struct
	{
	char data[MAXSIZE];
	int top;
	}STACK2;
int isempty2(STACK2 *s)
	{
	if(s->top==-1)
		return 1;
	return 0;
	}

int isfull2(STACK2 *s)
	{
	if(s->top==MAXSIZE-1)
		return 1;
	return 0;
	}

void push2(STACK2 *s,char ch)
	{
	s->top++;
	s->data[s->top]=ch;
	}

int pop2(STACK2 *s)
	{
	char ch;
	ch=s->data[s->top];
	s->top--;
	return ch;
	}

int peek2(STACK2 *s)
	{
	char ch;
	ch=s->data[s->top];
	return ch;
	}

void ini2(STACK2 *s)
	{
	s->top=-1;
	}
