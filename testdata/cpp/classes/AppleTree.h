#ifndef APPLETREE_H
#define APPLETREE_H

#include "Tree.h"

class Apple;

class AppleTree : public Tree {
protected:
    int nb_apples;
    Apple* apples;
public:
   void addApple(); 
};

#endif
