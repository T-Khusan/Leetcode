class Solution:
    def restoreString(self, s, indices):
        result = list(range(len(indices)))
        index = 0
        for item in indices:
            result[item] = s[index]
            index += 1
        return "".join(result)
a = Solution()
print(a.restoreString("codeleet",[4,5,6,7,0,2,1,3]))