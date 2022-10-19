
from ast import type_ignore
from ctypes import sizeof
from ctypes.wintypes import SIZE


arr = [[1,1,1,0,0,0],
       [0,1,0,0,0,0],
       [1,1,1,0,0,0],
       [0,0,2,4,4,0],
       [0,0,0,2,0,0],
       [0,0,1,2,4,0]]


print(type(arr))
n = len(arr)
max = -1
for r in range(n-2):
    for c in range(n-2):
        print(arr[r][c] , arr[r][c+1] , arr[r][c+2])
        print(" ",arr[r+1][c+1]," "              )   
        print(arr[r+2][c] ,arr[r+2][c+1] , arr[r+2][c+2],"\n")


        sum = arr[r][c] + arr[r][c+1] + arr[r][c+2] + arr[r+1][c+1] + arr[r+2][c] +arr[r+2][c+1] + arr[r+2][c+2]
        if sum > max:
            max = sum


print(max)