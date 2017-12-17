#include "Tree.h"
#include "AppleTree.h"
#include <iostream>

using namespace std;

int main() {
    Tree some_tree;
    AppleTree* apple_tree = new AppleTree();

    cout << "Hello from main" << endl;
    delete apple_tree;
}

