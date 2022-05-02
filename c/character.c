#include<stdio.h>
#include<math.h>
int main(){
	// printf("Hello\n");
	// int c;
	// c=getchar();
	// putchar(c);
	// printf("\n%c\n",c);
	// printf("\n%d\n",c);
	// int d=2;
	// printf("%d\n\n",d);
	// char name[10],name2[10];
	// scanf("%s",name);
	// scanf("%s",name2);
	// printf("%s\n",name);
	// printf("%s\n",name2);
	char name3[10],name4[10];
	fgets(name3,10,stdin);
	while ((getchar()) != '\n'); 
	fgets(name4,10,stdin);
	printf("%s\n",name3);    
	printf("%s\n",name4);
	printf("%f",pow(2.0,4.0));
	return 0;
}