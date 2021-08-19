package chapters

import (
	"fmt"
)

// go通过匿名组合的方式实现继承

type Rect struct {
	x, y float64
	width, height float64
}

func (rect *Rect) Close() error {
	return  nil
}
func (rect Rect) PrintCircle() {
	fmt.Printf("width:%d * height:%d = %d", rect.width, rect.height, rect.width * rect.height)
}
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

type Base struct {
	Name string
}

func (base *Base) Foo() {
	fmt.Println("base foo called")
}
func (base *Base) Bar() {
	fmt.Println("base bar called")
}

type Foo struct {
	Base  // 也可以是 *Base指针类型，但实例化Foo对象时需要外部传递Base类型的指针
	rect *Rect
	Name string
}
func (foo Foo) Bar() {
	foo.Base.Bar()
	foo.Foo()  // 匿名方式可直接调用Foo
	foo.rect.PrintCircle() // 无法直接调用  foo.PrintCircle()
	fmt.Println("Name:", foo.Name)  // 只能访问最外层Name
	fmt.Println("foo bar called")
}

func Oop_test()  {
	base := Base{"XXXXXXXXXXX"}
	rect := NewRect(1, 2, 3, 4)
	foo := Foo{base, rect, "YYYYYYYYYYYy"}
	foo.Bar()
}

//////////////////////////////////////////

// Go实现接口是 非侵入性接口

type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}
type IReader interface {
	Read(buf []byte) (n int, err error)
}
type IWriter interface {
	Write(buf []byte) (n int, err error)
}
type ICloser interface {
	Close() error
}

type File struct {
	// ...
}
func (f *File) Read(buf []byte) (n int, err error) {
	return 0, nil
}
func (f *File) Write(buf []byte) (n int, err error) {
	return 0, nil
}
func (f *File) Seek(off int64, whence int) (pos int64, err error) {
	return 0, nil
}
func (f *File) Close() error {
	return  nil
}

func Interface_test()  {
	var file1 IFile = new(File)
	var file2 IReader = new(File)
	var file3 IWriter = new(File)
	var file4 ICloser = new(File)
	var file5 IReader = file1   // IReader接口是IFile子集，所以IFile可以赋值给IReader,反过来不行
	if file6, ok := file2.(IFile); ok { // 检查file2接口指向的对象实例是否实现了IFile，这里返回true， file为IFile接口类型
		fmt.Println(file6, ok)
	}
	if file7, ok := file1.(IReader); ok {// 这里返回true， file7为IReader接口类型
		fmt.Println(file7, ok)
	}
	if file8, ok := file1.(*File); ok {// 查询file1接口指向的类对象类型是否为File，这里返回true，file8位File对象
		fmt.Println(file8, ok)
	}
	file9, ok := file4.(*Rect) // 判断file4接口指向的对象类型是否为Rect，这里返回false， file9位nil
	if ok {
		fmt.Println(file9, ok)
	}
	fmt.Println(&file1, &file2, &file3, &file4, &file5)

	MyPrintln(1, "2222", 1.2)
}

/////////////////////////////////

func MyPrintln(args ...interface{}) {
	for _, v := range  args {

		switch type1 := v.(type){
		case int:
			fmt.Println(type1, v)
		case string:
			fmt.Println(type1, v)
		default:
			fmt.Println(type1, v)
		}
	}
}



