# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def preorderTraversal(self, root):
        nodes=[]
        self._preorder(root,nodes)
        return nodes
    def _preorder(self,p,nodes):
        if p is None:
            return
        nodes.append(p.val)
        self._preorder(p.left,nodes)
        self._preorder(p.right,nodes)



root = TreeNode(1)
# root.left = TreeNode(None)
root.right = TreeNode(2)
root.right.left = TreeNode(3)

# root = TreeNode()

s = Solution()
print(s.preorderTraversal(root))

# printPreorder(root)