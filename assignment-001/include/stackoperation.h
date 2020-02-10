#define MAXSIZE 100
typedef struct
{
        char data[MAXSIZE];
        int top;
}STACK2;
typedef struct
{
        float data[MAXSIZE];
        int top;
}STACK1;

int isempty1(STACK1 *s);
int isfull1(STACK1 *s);
void push1(STACK1 *s,float num);
float pop1(STACK1 *s);
int peek1(STACK1 *s);
void ini1(STACK1 *s);

int isempty2(STACK2 *s);
int isfull2(STACK2 *s);
void push2(STACK2 *s,char ch);
int pop2(STACK2 *s);
int peek2(STACK2 *s);
void ini2(STACK2 *s);
