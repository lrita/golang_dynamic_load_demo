### Golang热插件demo
* 将需要插件化的代码编译为一个动态库，在package的`init()`函数中启动逻辑运行
* 采用`dlopen`来加载插件的动态库，加载后动态库的`init()`函数被执行，启动插件逻辑


### 遇到的问题
* 在Darwin平台上，golang的runtime在`dlopen`后初始化时，有bug[#11100](https://github.com/golang/go/issues/11100)、[##](https://groups.google.com/forum/#!topic/Golang-nuts/Vy8r05reLyw)，因此此方案并不是全平台通用方案。
* 在`dlopen`加载插件后，插件引用的第三方package变量与主进程引用的同名package不是同一内存空间，插件引用的package中相关变量会重新初始化一份，使得插件与主进程不能交互同一变量，使得插件机制的交互性很差（这个问题没有找到其他人相关反馈的印证，是我自己实现的问题还是golang本身机制问题就不得而知了）。
* 在1.6的cgo机制下，golang的export函数不能return go的内存空间、go的专属变量(GoInterface、GoPointer)，插件给主进程传递变量变得困难。（或许这个问题可以通过go调用cgo -> cgo调用c -> c调用cgo -> cgo 调用go的长链条实现?待实验）
