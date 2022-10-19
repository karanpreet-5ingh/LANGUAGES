
arr = [-4, 3, -9, 0, 4, 1]
n = len(arr)

a = [0,0,0]

for i in arr:
    if i > 0:
        a[0] = a[0] +1
    
    elif i < 0 :
        a[1] = a[1]  +1
    
    else :
        a[2] = a[2] +1
        
print('{0:.6f}'.format((a[0]/n)))
print('{0:.6f}'.format((a[1]/n)))
print('{0:.6f}'.format((a[2]/n)))