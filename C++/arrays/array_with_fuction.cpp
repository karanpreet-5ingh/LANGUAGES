#include<iostream>
using namespace std;
void print_array(int arr[], int size){
    cout<<"printig the array"<<endl;
    //printinhg the array 
    for (int i=0 ;i<size;i++)
    {
        cout<<arr[i]<<" ";


    }
}

int main(){
    int fifth[10]={0};
    int   n=10;
    int fifth_size  =sizeof(fifth)/ sizeof(int); //so   here maa lo ki jo array h uska total actual size  40 hai pr aako wo chahie 4 because 4 indexes h usme 
    print_array(fifth,fifth_size);
}
