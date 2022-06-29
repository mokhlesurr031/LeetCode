class Solution:
    def sumOfSquare(self, n):
        sq = 0
        while n > 0:
            rem = n % 10
            sq = sq + (rem * rem)
            n = n // 10
        return sq

    def isHappy(self, n: int) -> bool:
        visited = set()

        while n not in visited:
            visited.add(n)
            sum_of_sq = self.sumOfSquare(n)
            if sum_of_sq==1:
                return True
            n = sum_of_sq
        return False
