class Solution(object):
    def removeElement(self, nums, val):
        """
        :type nums: List[int]
        :type val: int
        :rtype: int
        """
        count = 0

        for v in nums[:]:
            if v == val:
                nums.remove(v)
                continue
            count += 1
    
        return count

        
            
        