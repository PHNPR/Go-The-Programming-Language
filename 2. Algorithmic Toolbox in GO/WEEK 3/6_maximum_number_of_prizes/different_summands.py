def optimal_summands(n):
    for i in range(1,n+1):
        if i * (i + 1) <= 2 * n < (i + 1)*(i + 2):
            a = [j for j in range(1,i)] 
            a.append(n - sum(a))
            return a


if __name__ == '__main__':
    n = int(input())
    summands = optimal_summands(n)
    print(len(summands))
    print(*summands)