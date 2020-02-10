/*validate expression entered by user check bracket balancing as well as alphabeticals*/
#include<stdio.h>
#include<string.h>
#include<stdlib.h>
int check_exp(char exp[])
{
	int i,n,temp;
	int flag=0,bracket=0;
	n=strlen(exp);
	i=0;
	while(i<n){
		temp=(int)(exp[i]);
		if((int)(exp[i])>=65 && (int)(exp[i])<=122)
			return 0;
		if(exp[i]==' ')
			i++;
		if(exp[i]=='(')
		{
			bracket=bracket+1;
			i++;
		}
		if(exp[i]==')')
		{
			bracket=bracket-1;
			i++;
		}
		if((int)(exp[i])>=48 && (int)(exp[i])<=57)
		{
			while((int)(exp[i])>=48 && (int)(exp[i])<=57)
			{
			i++;
			}
			if(flag==1)
			{
				flag=0;
			}
			
		}
		if(exp[i]=='+' || exp[i]=='-' || exp[i]=='*' || exp[i]=='/')
		{
			i++;
			if(flag==0)
			{
				flag=1;
			}
			else
			{
				return 0;
			}

		}
		

	}
	if(flag!=0 || bracket!=0)
	{
		return 0;
	}
	else{
		return 1;
	}

}

