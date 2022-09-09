points = []
for _ in range(int(input())):
    points.append(list(map(int,input().split())))

count , dots = 0 , []

while len(points) > 0:
    i = min(points , key= lambda x : x[1])[1]
    z = points[:]
    for p in z :
        if p[0] <= i <= p[1]:
            points.remove(p)
    if len(z) != len(points):
        count += 1
        dots.append(i)
print(count)
print(*dots)