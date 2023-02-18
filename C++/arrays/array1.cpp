//array ko countinues location chahie to store the index
// array me 1st location ka index zero hoga


#include <iostream>
using namespace std;
int main() {
    int number[15];
    cout<<"value at 14 index "<<  number[15]<<endl;
    cout<<endl<<"Everthing is Fine "<<endl <<endl;
     // intialising an array
    int  second[3] = {5,7,11};
    cout<<"value at 2 index "<<  second[2]<<endl;
// abb agar to aapne array ko initialize kiya hai tbb to app us us values ko dekh paoge pr jo value aapne intialize nhi ki wo value apki 0 ho jaegi while printing

    int third[15] = {2,7};
    int n=12 ;
    cout<<"printing the array"<<endl;
    // printing the array 
    for(int i = 0; i< n; i++){
        cout<< third[i]<<" ";
    }
// not possible with any other value other than 0
    int fourth[10] = {21};// this can intialize all the element of array with zero 
    n = 10 ;
    cout<<"printing the array/n/n"<<endl;
    for(int i = 0; i < n; i++){
        cout<< fourth[i]<<" ";


    }




return 0;
}
