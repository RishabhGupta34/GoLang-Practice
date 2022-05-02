#include<stdio.h>
#include<string.h>

int main(){
    char first[50],second[50];
    printf("First string: ");
    fgets(first,50,stdin);
    printf("Second string: ");
    fgets(second,50,stdin);
    char third[100];
    first[strlen(first)-1]='\0';
    second[strlen(second)-1]='\0';
    strcpy(third,first);
    strcat(third,second);
    printf("Final string: %s\n",third);
    return(0);
}