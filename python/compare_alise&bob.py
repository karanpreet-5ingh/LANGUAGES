def compareTriplets(a, b):
    # Write your code here
    a_n = len(a)
    ar = [0,0]
    for i in range(a_n):
        if(a[i] > b[i]):
            ar[0] += 1
        elif(a[i] < b[i]):
            ar[1] += 1
            
    return ar 

a = [1, 2, 3]
b = [3, 2, 1]    
ar = compareTriplets(a,b)
print(ar)