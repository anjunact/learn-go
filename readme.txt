go doc builtin

数组同样是值类型的:将一个数组赋 值给另一个数组，会复制所有的元素
a := [...]int{1, 2, 3}
slice 是一个指向 array 的指针
utf8.RuneCountInString() 字符串个数
recover 仅在延迟函数中有效
new(File) 和 &File{} 是等价的