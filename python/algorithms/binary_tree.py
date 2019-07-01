class Node:
    def __init__(self, value=None, child=[]):
        self.value = value
        self.child = child

    def add_child(self, node):
        self.child.append(node)

class BinaryTree:
    def __init__(self):
        self.root = Node(child=[None,None])

    def add(self, value: int) -> int:
        target = self.root
        while True:
            if target.value is None:
                target.value = value
                return target.value
            elif value < target.value:
                if target.child[0] is None:
                    target.child[0] = Node(child=[None,None])
                target = target.child[0]
            else:
                if target.child[1] is None:
                    target.child[1] = Node(child=[None,None])
                target = target.child[1]

    def search(self, value: int):
        return self.__search(self.root, value)

    def __search(self, node, value) -> bool:
        if node.value == value:
            return True
        if value < node.value and not node.child[0] is None:
            return self.__search(node.child[0], value)
        elif not node.child[1] is None:
            return self.__search(node.child[1], value)
        return False

if __name__ == '__main__':
   b = BinaryTree() 
   assert b.add(1) == 1
   assert b.add(2) == 2
   assert b.search(1) == True
   assert b.search(2) == True
   assert b.search(3) == False
