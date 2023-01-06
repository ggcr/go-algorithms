import unittest
from code import *

class Tests(unittest.TestCase):
    def test_initial(self):
        self.assertEqual(maxIceCream([1,3,2,4,1], 7), 4)

if __name__ == '__main__':
    unittest.main()
