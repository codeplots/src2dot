# src2dot
[![Build Status](https://travis-ci.org/codeplots/src2dot.svg?branch=master)](https://travis-ci.org/codeplots/src2dot) [![codecov](https://codecov.io/gh/codeplots/src2dot/branch/master/graph/badge.svg)](https://codecov.io/gh/codeplots/src2dot) 
![dependency graph of src2dot](/docs/examples/src2dot_dependencies.png)

```
cd your-repo

# run universal-ctags for all supported languages
ctags --all-kinds=* --fields=* -R --extras=+frF .

# run cscope for C/C++/Java 
cscope -kbc

# run starscope for Go
starscope -e cscope

# run pycscope for Python
pycscope -R

# finally run src2dot
# use grep to filter out unwanted nodes and pipe output into dot
./src2dot | grep -v "test" | dot -Tpng -o dependency_graph.png
```
