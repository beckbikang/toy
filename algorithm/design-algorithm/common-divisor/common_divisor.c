#include<stdio.h>


int gcd(int m, int n){

    int r;

    while(n >0){
        r = m %n;
        m = n;
        n = r;
    }
    return m;
}

int max(int x, int y ){
    if(x > y) {
        return  x;
    }
    return y;
}

float getPi(long n){

    long i = 1;
    long cent = 0;
    float ret = 0;

    while(i < n){
        cent = 2*i;
        ret += 1.0/(cent-1) - 1.0/(cent+1);
        i = i+2;
    }

    return  ret * 4;
}


int main(int argc, char const *argv[])
{
    printf("%d,%d => %d\n",60,24, 12);
    printf("%2.10f", getPi(1000000000));
    return 0;
}
