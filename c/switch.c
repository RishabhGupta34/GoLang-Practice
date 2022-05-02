#include<stdio.h>

int main(){
    int n;
    printf("Enter an integer: ");
    scanf("%d",&n);
    switch(n){
        case 1:
            printf("Color is Red %d",n);
            break;
        case 2:
            printf("Color is Green %d",n);
            break;
        case 3:
            printf("Color is Blue %d",n);
            break;
        default:
            printf("Color is Invalid %d",n);
    }
    printf("\n");
}