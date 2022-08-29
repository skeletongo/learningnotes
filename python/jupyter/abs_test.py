import unittest

from abs import abs as myabs

class TestAbs(unittest.TestCase):

    # setUp,tearDown 每调用一个测试方法的前后分别被执行

    def setUp(self):
        print('setUp...')

    def tearDown(self):
        print('tearDown...')

    def test_n(self):
        print("test_n...")
        self.assertTrue(myabs(0)==0)
        self.assertEqual(myabs(1),1)
        self.assertEqual(myabs(-1), 1)
        self.assertEqual(myabs(1.2),1.2)
        self.assertEqual(myabs(-1.2),1.2)

    def test_typerror(self):
        print("test_typerror...")
        with self.assertRaises(TypeError):
            myabs("a")

if __name__ == '__main__':
    unittest.main()

# python3  -m unittest .\abs_test.py