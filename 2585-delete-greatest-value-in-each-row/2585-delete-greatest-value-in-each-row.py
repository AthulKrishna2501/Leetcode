class Solution:
    def deleteGreatestValue(self, grid: List[List[int]]) -> int:
        ans = 0
        for i in grid:
            i.sort()



        for i in range(len(grid[0])):
            temp=0
            for j in grid:
                temp=max(j[i],temp)

            ans+=temp           


        return ans    