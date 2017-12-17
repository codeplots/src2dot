import trees

class Apple(object):
    colour = "red"
    def __init__(self):
        print("apple.init")
        self.size = 100
        self.tree = trees.AppleTree()
    def ripe(self):
        print("apple.ripe")
