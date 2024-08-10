# link to question: https://leetcode.com/problems/regions-cut-by-slashes/
from typing import List

class UnionFind:
    def __init__(self, size):
        self.parent = list(range(size))
        self.rank = [1] * size

    def find(self, x):
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x, y):
        rootX = self.find(x)
        rootY = self.find(y)
        if rootX != rootY:
            if self.rank[rootX] > self.rank[rootY]:
                self.parent[rootY] = rootX
            elif self.rank[rootX] < self.rank[rootY]:
                self.parent[rootX] = rootY
            else:
                self.parent[rootY] = rootX
                self.rank[rootX] += 1

class Solution:
    def regionsBySlashes(self, grid: List[str]) -> int:
        n = len(grid)
        size = 4 * n * n  # Setiap sel dibagi menjadi 4 wilayah
        uf = UnionFind(size)

        for i in range(n):
            for j in range(n):
                # Indeks untuk wilayah segitiga kiri atas dalam sel saat ini
                index = 4 * (i * n + j)
                if grid[i][j] == '/':
                    uf.union(index, index + 3)  # kiri atas dan kanan bawah
                    uf.union(index + 1, index + 2)  # kanan atas dan kiri bawah
                elif grid[i][j] == '\\':
                    uf.union(index, index + 1)  # kiri atas dan kanan atas
                    uf.union(index + 2, index + 3)  # kiri bawah dan kanan bawah
                else:
                    # satukan semua 4 wilayah bersama-sama
                    uf.union(index, index + 1)
                    uf.union(index + 1, index + 2)
                    uf.union(index + 2, index + 3)

                # Union dengan sel di sebelah kanan
                if j + 1 < n:
                    uf.union(index + 1, 4 * (i * n + j + 1) + 3)

                # Union dengan sel di bawah
                if i + 1 < n:
                    uf.union(index + 2, 4 * ((i + 1) * n + j))

        # Hitung jumlah wilayah yang berbeda dengan menghitung jumlah akar unik
        return sum(uf.find(x) == x for x in range(size))


# Test case
sol = Solution()
print(sol.regionsBySlashes([" /", "/ "]))  # 2
print(sol.regionsBySlashes([" /", "  "]))  # 1
print(sol.regionsBySlashes(["\\/", "/\\"]))  # 4