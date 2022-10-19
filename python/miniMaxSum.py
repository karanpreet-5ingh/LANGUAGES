arr = [2,1,3,4]
n = len(arr)
sum = 0
arr.sort()
print(arr)
for i in range(n-2):

    sum = sum + arr[i]
print(sum)