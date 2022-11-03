# cgo_demo

## demo1
仅启用cgo

## demo2
在go代码中添加C代码函数继续调用

## demo3

在go代码中加载标准的c代码（.c,.h）

## demo4

`gcc -c -o number.o number.c & ar rcs libnumber.a number.o`

在go中调用C静态库文件

## demo5

`gcc number.c -fPIC -shared -o libnumber.so`

在go中调用C动态库文件（需要设置LD_LIBRARY_PATH参数）


