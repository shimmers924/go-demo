package main

import (
	"demo/test"
	"fmt"
)

//定义 ：结构体是用户定义的类型，表示若干个字段的集合。
//命名的结构体
type Employee struct {
	name, position string
	age, sex       int
}

type address struct {
	city  string
	state string
}

type person struct {
	firstName string
	lastName  string
	address
}

type user struct {
	string
	int
}

type image struct {
	data map[int]int
}

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() { //每一个go程序中都需要有一个main函数作为包的入口文件
	// fmt.Println("hello")

	// // 声明语句 var,const,type,func
	// var num = rand.Intn(10) + 1 //变量  只可在函数内访问
	// fmt.Println(num)            //打印可换行

	// var num2 string = "wuhongjie"
	// fmt.Printf(num2)

	// // 声明一组变量
	// var arr1, arr2 = 10, 20
	// fmt.Println(arr1, arr2)

	// var (
	// 	arr3 = "22"
	// 	arr4 = 200
	// 	arr5 = 10.2
	// )
	// fmt.Print(arr3, arr4, arr5)
	// fmt.Printf("arr3 is %v,arr4 is %v", arr3, arr4)

	// const STATUS = 10 //常量 可以在函数外访问
	// fmt.Println(STATUS)

	// // 简短变量的声明方式
	// path := 1
	// fmt.Println(path)

	// c := rand.Float64() * 10
	// d := rand.Float32() * 10
	// hh := rand.Intn(200)
	// fmt.Println(c, d, hh)

	// //声明一组变量
	// e, f := "hahah", 20
	// fmt.Println(e, f)

	// //常量
	// const hello = "hello words"
	// fmt.Print(hello)
	// fmt.Printf(hello)
	// fmt.Println(hello)
	// //创建一个带类型的常量
	// const a string = "hdjdhfs"
	// const b int = 34284924
	// fmt.Println(a)
	// fmt.Println(b)

	// // ---------------------------------数组 start⬇️ -----------------------------------
	// var array [3]int //[n]t n表示元素数量 t表示元素类型。元素的数量 n 也是该类型的一部分
	// array[0] = 1
	// array[1] = 2
	// array[2] = 3
	// //数组简略声明
	// first := 1
	// second := 2
	// array2 := [5]int{first, second, 7, 8, 9}
	// array3 := [5]int{first, second}
	// array4 := [...]string{"fggsf", "武红杰"}
	// //传入数组格式的参数并返回多个值
	// data, lenght := a(array2)
	// fmt.Println(array, array2, array3, array4, data, lenght)
	// //使用range遍历数组
	// sum := 0
	// for i, v := range array2 {
	// 	fmt.Printf("%d the element of a is %.2f\n", i, v)
	// 	sum += v
	// }
	// fmt.Println("\nsum of all elements of a", sum)

	// //多维数组 array[x][y]t  x表示行 y表示列
	// array5 := [3][2]string{
	// 	{"name", "武红杰"},
	// 	{"age", "20"},
	// 	{"height", "165"},
	// }
	// selectArray(array5) //遍历数组
	// var array6 [3][2]string
	// array6[0][0] = "apple"
	// array6[0][1] = "iphone"
	// array6[1][0] = "iphone"
	// array6[1][1] = "mac"
	// selectArray(array6)

	// // ---------------------------- 切片 start ⬇️ ------------------------------------
	// //使用语法a[start:end] 创建一个从 a 数组索引 start 开始到 end - 1 结束的切片 切片是由数组建立的一种方便、灵活且功能强大的包装（Wrapper）。切片本身不拥有任何数据。它们只是对现有数组的引用。
	// var slice1 []int = array2[1:5] // creates a slice from a[1] to a[3]
	// // 修改切片的值 ，所做更改将反应再底层数组中
	// for i := range slice1 {
	// 	slice1[i]++
	// }
	// // 当多个切片共用相同的底层数组时，每个切片所做的更改将反映在数组中。
	// slice2 := array2[1:3]
	// for i := range slice2 {
	// 	slice2[i]--
	// }

	// slice3 := array2[3:4]
	// // slice3[0] = 100
	// num := len(slice2) //切片的长度是切片中的元素数
	// // cap := cap(slice2)            // 切片的容量是从创建切片索引开始的底层数组中元素数
	// slice3 = slice3[:cap(slice3)] // 从 创建切片索引开始重置容量 re-slicing furitslice till its capacity
	// fmt.Print(slice1, slice2, num, cap(slice3))

	// i := make([]int, 3, 3) //使用make 创建一个切片
	// fmt.Print(i)
	// cars := []string{"php", "go", "python", "java"}
	// fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
	// cars = append(cars, "css")                                                         //向切片中追加一个元素
	// cars = append(cars, "javascript", "vue", "c++", "c#")                              //向切片中追加多个元素
	// fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) // capacity of cars is doubled to 6

	// //多维切片
	// slice4 := [][]string{
	// 	{"c", "c++"},
	// 	{"python"},
	// 	{"Go", "Rust"},
	// }
	// fmt.Print(slice4)
	// for _, v1 := range slice4 {
	// 	for _, v2 := range v1 {
	// 		fmt.Printf("%s ", v2)
	// 	}
	// 	fmt.Printf("\n")
	// }

	// //切片引起的内存优化问题   切片持有对底层数组的引用。只要切片在内存中，数组就不能被垃圾回收一种解决方法是使用 copy 函数 func copy(dst，src[]T)int 来生成一个切片的副本。这样我们可以使用新的切片，原始数组可以被垃圾回收。
	// countriesNeeded := countries()
	// fmt.Println(countriesNeeded)

	// //--------------------------------- 函数 start ⬇️ ----------------------------
	// // var back = calculateBill("PH729753958358357", "849248") //单返回值
	// // h, d := manyCallback(10.5, 8.9)                         //多返回值
	// // k, j := namespace(10, 8)                                //命名回值
	// // m, _ := whiteFull(10, 8)                                //空白符
	// // fmt.Println(back)
	// // fmt.Println(h, d)
	// // fmt.Println(k, j)
	// // fmt.Println(m)

	// //--------------------------------- 包 start ⬇️ ----------------------------
	// // process.Process()
	// // fmt.Print("武红杰11")

	// //--------------------------------- 可变函数 start ⬇️ ----------------------------
	// //可变函数 传入int数值，编译器转换为一个int类型切片被传入find函数
	// find(3, 43, 348, 249024, 242, 242)
	// //给可变函数传入切片
	// slice5 := []int{1, 3, 2}
	// // fmt.Print(slice5)
	// find(2, slice5...)

	// //--------------------------------- map start⬇️ ----------------------------
	// //定义：map 是在 Go 中将值（value）与键（key）关联的内置类型。通过相应的键可以获取到值。
	// //创建语法：make(map[type of key]type of value)
	// //键不一定只能是 string 类型。所有可比较的类型，如 boolean，interger，float，complex，string 等，都可以作为键
	// var map1 map[string]int
	// if map1 == nil {
	// 	fmt.Println("map is nil")
	// 	map1 = make(map[string]int)
	// }
	// map2 := make(map[string]int)
	// map2["stave"] = 12000
	// map2["stavew"] = 1000
	// map2["stavef"] = 12000
	// map2["staved"] = 12000
	// fmt.Println(map2)

	// //声明时初始化map
	// map3 := map[string]string{
	// 	"name": "武红杰",
	// 	"age":  "18",
	// }
	// map3["sex"] = "18"

	// //获取map元素
	// mapValue := map3["name"]   //获取一个存在的值
	// mapValue2 := map3["wight"] //获取一个不存在的值 会根据类型返回相应的空值
	// fmt.Println(map3, mapValue, mapValue2)

	// //检测某个key是否存在
	// key := "age"
	// value, ok := map3[key]
	// if ok == true {
	// 	fmt.Print("info of", value)
	// } else {
	// 	fmt.Print("not found this key", key)
	// }
	// //遍历map
	// for i, value := range map3 {
	// 	fmt.Printf("map3[%s] = %d\n", i, value)
	// }
	// //删除map中元素
	// delete(map3, "age")
	// fmt.Println(map3)

	// //获取map长度
	// count := len(map3)
	// fmt.Println(count)

	// //map 赋值为引用类型。给一个新的变量赋值，他们之乡同一个数据结构，因此改变一个变量，就会影响到另一个变量
	// //map 相等性 map之间不能使用== 操作符判断，== 只能用来检查是否为nil。判断两个 map 是否相等的方法是遍历比较两个 map 中的每个元素

	// //--------------------------------- string start⬇️ ----------------------------
	// // string1 := "这是一个字符串"
	// string1 := "Hello World"
	// //单独获取字符串的每一个字节
	// for i := 0; i < len(string1); i++ { //len 返回字符串中字节的数量
	// 	// fmt.Printf("%x ", string1[i]) //%x 格式限定符用于指定16进制编码
	// 	fmt.Printf("%c ", string1[i]) //%c  格式限定符用于打印字符串的字符
	// }

	// // rune是go内建类型 也是int32的别称 在go语言中，rune表示一个代码点，代码点无论占用多少个字节。都可以用一个rune来表示
	// string2 := []rune("这是一个字符串")
	// for i := 0; i < len(string2); i++ { //len 返回字符串中字节的数量
	// 	// fmt.Printf("%x ", string1[i]) //%x 格式限定符用于指定16进制编码
	// 	fmt.Printf("%c ", string2[i]) //%c  格式限定符用于打印字符串的字符
	// }

	// //循环字符串
	// string3 := "循环字符串"
	// for index, value := range string3 {
	// 	fmt.Printf("%c start at byte %d\n", value, index)
	// }

	// //用字节切片构建字符串
	// byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9} //16进制
	// byteSlice2 := []byte{67, 97, 102, 195, 169}       //10 进制
	// string4 := string(byteSlice)
	// string5 := string(byteSlice2)
	// fmt.Println(string4, string5)

	// //用rune 构建字符串
	// runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072} //6 进制的 Unicode 代码点
	// string6 := string(runeSlice)
	// fmt.Println(string6)

	// //字符串长度
	// string7 := "获取字符串长度"
	// fmt.Println(utf8.RuneCountInString(string7))

	// //字符串不可变  go的字符串一旦被创建，便无法修改 如果想要修改字符串需要使用rune转为切片 然后再转为一个字符串
	// string8 := "字符串无法改变如果想要修改字符串需要使用rune转为切片然后再转为一个字符串"
	// coptyString8 := []rune(string8)
	// coptyString8[0] = '哈'
	// string9 := string(coptyString8)
	// fmt.Println(string9)

	//--------------------------------- 指针 start⬇️ ----------------------------
	//指针是一种存储变量内存地址（memory address）的变量 变量 b 的值为 156，而 b 的内存地址为 0x1040a124。变量 a 存储了 b 的地址。我们就称 a 指向了 b。
	//声明指针  指针变量的类型为 *T，该指针指向一个 T 类型的变量。
	pointer1 := 2
	var pointer2 *int = &pointer1
	pointer1 = 3
	fmt.Printf("Type of pointer1 is %T\n", pointer2) // Type of pointer1 is *int
	fmt.Println("address of b is", pointer2)         //address of b is 0xc00001a1a0
	fmt.Println(pointer1, pointer2)                  //值为3 0xc00001a1a0

	// 指针的零值
	a := 8
	var pointer3 *int
	if pointer3 == nil {
		fmt.Println(a, pointer3)
		pointer3 = &a
		fmt.Println("pointer3 after initialization is", pointer3)
	}

	//指针的解引用 语法为：*a 指针的解引用可以获取指针所指向变量的值
	c := 10
	var pointer4 = &c
	fmt.Println(pointer4)
	*pointer4++
	fmt.Println(c) //输出11
	//向函数传递指针参数
	e := 10
	var d *int = &e
	pointerParam(d)
	fmt.Println(e)

	//不要向函数传递数组的指针，而应该使用切片 这种方式向函数传递一个数组指针参数，并在函数内修改数组。尽管它是有效的，但却不是 Go 语言惯用的实现方式。我们最好使用切片来处理。
	f := [3]string{"php", "mysql", "nginx"}
	slice6 := f[:]
	slice6[0] = "go"
	fmt.Println(slice6)

	//go 不支持指针运算

	//--------------------------------- 结构体 start⬇️ ----------------------------
	// 命名结构体赋值
	emp1 := Employee{
		name:     "这是一个命名结构体",
		position: "职位",
		age:      18,
		sex:      2,
	}
	em2 := Employee{"臭沙雕", "总经理", 18, 3} //省略字段名时，需要保证顺序与定义结构体顺序一致
	fmt.Println(emp1, em2)

	//创建匿名结构体并赋值
	em3 := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
	}
	fmt.Println(em3)

	//结构体零值
	var em4 Employee
	fmt.Println(em4)

	//访问结构体的字段  (点号操作符 . 用于访问结构体的字段)
	emp6 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("Name:", emp6.name)
	fmt.Println("Position:", emp6.position)
	fmt.Println("Age:", emp6.age)
	fmt.Println("Sex:", emp6.sex)

	// 创建0值的struct 以后再给各个字段赋值。
	var emp7 Employee
	emp7.name = "Jack"
	emp7.position = "Adams"
	fmt.Println("Employee 7:", emp7)

	//创建指向结构体的指针
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", (*emp8).name)
	fmt.Println("First Name:", emp8.name) // 可以使用 emp8.firstName 来代替显式的解引用 (*emp8).firstName
	fmt.Println("Age:", (*emp8).age)

	//匿名字段 当我们创建结构体时，字段可以只有类型，而没有字段名。这样的字段称为匿名字段（Anonymous Field）。
	u := user{"Naveen", 50}
	fmt.Println(u)

	//嵌套结构体 结构体的字段有可能也是一个结构体。这样的结构体称为嵌套结构体。
	var pp person
	pp.firstName = "Naveen"
	pp.lastName = "H"
	pp.address = address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", pp.firstName)
	fmt.Println("Age:", pp.lastName)
	fmt.Println("City:", pp.address.city)   //city is promoted field 提升字段
	fmt.Println("State:", pp.address.state) //state is promoted field	提升字段

	//导出结构体和字段 如果结构体名称以大写字母开头，则它是其他包可以访问的导出类型（Exported Type）。同样，如果结构体里的字段首字母大写，它也能被其他包访问到。
	var spec test.Object
	spec.Maker = "apple"
	spec.Price = 10000
	fmt.Println("spec:", spec)

	//结构体相等性 结构体是值类型。如果它的每一个字段都是可比较的，则该结构体也是可比较的。如果两个结构体变量的对应字段相等，则这两个变量也是相等的。
	name1 := address{"Steve", "hahah2"}
	name2 := address{"Steve", "hahah"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := address{city: "Steve", state: "Jobs"}
	name4 := address{}
	name4.city = "Steve"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}

	// 如果结构体包含不可比较的字段 则结构体变量也不可比较
	// image1 := image{data: map[int]int{
	// 	0: 100,
	// }}

	// image2 := image{data: map[int]int{
	// 	1: 200,
	// }}

	// if image1 == image2 {
	// 	fmt.Println("image1 and image2 are equal")
	// } else {
	// 	fmt.Println("image1 and image2 are not equal")
	// }

	//--------------------------------- 方法 start⬇️ ----------------------------
	//方法其实就是一个函数，在 func 这个关键字和方法名中间加入了一个特殊的接收器类型。接收器可以是结构体类型或者是非结构体类型。接收器是可以在方法的内部访问的。
	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "Los Angeles",
			state: "California",
		},
	}
	p.fullAddress()       //调用 person 类型的 fullAddress() 方法
	fullAddressfun(name1) //调用 fullAddressfun() 函数

	//为什么我们已经有函数了还需要方法呢？
	//	1.Go 不是纯粹的面向对象编程语言，而且Go不支持类。因此，基于类型的方法是一种实现和类相似行为的途径。
	//  2.相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的。

	//指针接收器与值接收器
	// 一般来说，指针接收器可以使用在：对方法内部的接收器所做的改变应该对调用者可见时。
	// 指针接收器也可以被使用在如下场景：当拷贝一个结构体的代价过于昂贵时。考虑下一个结构体有很多的字段。在方法内使用这个结构体做为值接收器需要拷贝整个结构体，这是很昂贵的。在这种情况下使用指针接收器，结构体不会被拷贝，只会传递一个指针到方法内部使用。
	// 在其他的所有情况，值接收器都可以被使用.
	em5 := Employee{"臭沙雕哈哈哈", "总经理", 18, 3}
	// 调用值接收器方法
	em5.changeName("为什么还不发工资！！！！")
	//调用指针接收器方法
	(&em5).changeAge(20)

	//匿名字段的方法
	// p := person{
	// 	firstName: "Elon",
	// 	lastName:  "Musk",
	// 	address: address{
	// 		city:  "Los Angeles",
	// 		state: "California",
	// 	},
	// }
	// p.fullAddress() //访问 address 结构体的 fullAddress 方法

	//在方法中使用值接收器 与 在函数中使用值参数
	// 	1.当一个函数有一个值参数，它只能接受一个值参数。
	// 	2.当一个方法有一个值接收器，它可以接受值接收器和指针接收器。
	//在方法中使用指针接收器 与 在函数中使用指针参数
	// 	1.函数使用指针参数只接受指针，而使用指针接收器的方法可以使用值接收器和指针接收器

	//在非结构体上定义方法
	// 为了在一个类型上定义一个方法，方法的接收器类型定义和方法的定义应该在同一个包中。到目前为止，我们定义的所有结构体和结构体上的方法都是在同一个 main 包中，因此它们是可以运行的
	// int1 := myInt(5)
	// int10 := myInt(10)
	// sum := int1.add(int10)
	// fmt.Println("Sum is", sum)

	// --------------------------------- 接口 start⬇️  ---------------------------------
	name := test.MyString("Sam Anderson")
	var v test.VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels())

	pemp1 := test.Permanent{1, 5000, 20}
	pemp2 := test.Permanent{2, 6000, 30}
	cemp1 := test.Contract{3, 3000}
	employees := []test.SalaryCalculator{pemp1, pemp2, cemp1}
	test.TotalExpense(employees)

	fmt.Println()

	//空接口
	s := "hello word"
	describe(s)
	m := 55
	describe(m)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(strt)

	//类型断言 (类型断言用于提取接口的底层值（Underlying Value))
	var k interface{} = 56
	q := k.(int) //在语法 i.(T) 中，接口 i 的具体类型是 T，该语法用于获得接口的底层值。
	fmt.Println(q)

	var r interface{} = "这是一个字符串"
	t, ok := r.(int) //如果 r 的具体类型是 T，那么 t 赋值为 r 的底层值，而 ok 赋值为 true。如果 r 的具体类型不是 T，那么 ok 赋值为 false，t 赋值为 r 类型的零值
	fmt.Println(t, ok)

	//类型选择 语法为i.(type)
	var z interface{} = 56
	switch z.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", z.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", z.(int))
	default:
		fmt.Printf("Unknown type\n")
	}

	//实现接口 使用指针接收
	// var d1 test.Add
	// p1 := test.Address{}
	// d1 = p1
	// d1.Add()
	// p2 := test.Address{"ahjdahjd", "hirwiri"}
	// d1 = &p2
	// d1.Add()

	var d2 test.Add
	address1 := test.Address{"ahjdahjd", "hirwiri"}
	// d2 = address1 //对于使用指针接受者的方法，用一个指针或者一个可取得地址的值来调用都是合法的。但接口中存储的具体值（Concrete Value）并不能取到地址，对于编译器无法自动获取 a 的地址，于是程序报错。
	d2 = &address1
	d2.Add()

	// 实现多个接口
	cc := test.Contract{1, 100000}
	var ss test.SalaryCalculator2 = cc
	fmt.Println("\nSalaryCalculator  =", ss.CalculateSalary2())
	var l test.SalaryCalculator = cc
	fmt.Println("\nSalaryCalculator  =", l.CalculateSalary())

	// 接口嵌套 (尽管 Go 语言没有提供继承机制，但可以通过嵌套其他的接口，创建一个新接口。)

	// 接口的零值 （接口的零值是 nil。对于值为 nil 的接口，其底层值（Underlying Value）和具体类型（Concrete Type）都为 nil。）

	// //--------------------------------- 并发 start⬇️ ----------------------------
	// 并发 Go 是并发式语言，而不是并行式语言。在讨论 Go 如何处理并发之前，我们必须理解何为并发，以及并发与并行的区别。
	// 并发是指立即处理多个任务的能力 我们可以想象一个人正在跑步。假如在他晨跑时，鞋带突然松了。于是他停下来，系一下鞋带，接下来继续跑。这个例子就是典型的并发。这个人能够一下搞定跑步和系鞋带两件事，即立即处理多个任务。
	// 并行是指同时处理多个任务 我们同样用这个跑步的例子来帮助理解。假如这个人在慢跑时，还在用他的 iPod 听着音乐。在这里，他是在跑步的同时听音乐，也就是同时处理多个任务。这称之为并行。
	// 并行不一定会加快运行速度，因为并行运行的组件之间可能需要相互通信。在我们浏览器的例子里，当文件下载完成后，应当对用户进行提醒，比如弹出一个窗口。于是，在负责下载的组件和负责渲染用户界面的组件之间，就产生了通信。在并发系统上，这种通信开销很小。但在多核的并行系统上，组件间的通信开销就很高了。所以，并行不一定会加快运行速度！
	// Go 编程语言原生支持并发。Go 使用 Go 协程（Goroutine） 和信道（Channel）来处理并发。

	// --------------------------------- 协程 start⬇️ ----------------------------
	// Go 协程是与其他函数或方法一起并发运行的函数或方法

}

// //--------------------------------- Defer start⬇️ ----------------------------
// //含有 defer 语句的函数，会在该函数将要返回之前，调用另一个函数
// nums := []int{78, 109, 2, 563, 300}
// largest(nums)
// //defer 不仅限于函数的调用，调用方法也是合法的。

// 单返回值函数
func calculateBill(orderSn string, userId string) string {
	var name string = orderSn + "\n" + userId
	return name
}

// 多返回值函数
func manyCallback(width, height float32) (float32, float32) {
	var area = width * height
	var perimeter = (width + height) * 2
	return area, perimeter
}

// 命名返回值
func namespace(width, height int) (area, perimeter int) {
	area = width * height
	perimeter = (width + height) * 2
	return //不需要明确指定返回值，默认返回area,perimeter
}

// _空白符
func whiteFull(width, height int) (area, perimeter int) {
	area = width * height
	perimeter = (width + height) * 2
	return //不需要明确指定返回值，默认返回area,perimeter
}

// 参数为数组类型的函数  todo需要再好好研究一下原始值为什么会改变
func a(datas [5]int) (b [5]int, lenght int) {
	datas[0] = 100
	b = datas
	lenght = len(datas)
	return b, lenght
}

// 使用range遍历二维数组
func selectArray(array [3][2]string) {
	for _, v1 := range array {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Print("\n")
	}
}

// 使用copy函数复制一个切片
func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}

// 可变函数   语法：如果函数最后一个参数被记作 ...T ，这时函数可以接受任意个 T 类型参数作为最后一个参数。
// 可变参数函数的工作原理是把可变参数转换为一个新的切片。以上面程序中的第 22 行为例，find 函数中的可变参数是 89，90，95 。 find 函数接受一个 int 类型的可变参数。因此这三个参数被编译器转换为一个 int 类型切片 int []int{89, 90, 95} 然后被传入 find函数。
func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for _, v := range nums {
		if v == num {
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

//参数为指针参数
func pointerParam(val *int) {
	*val = 55
}

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

// 创意一个值接收器的方法 ：语法为 func(t 结构体名称) 方法名(参数) (返回值){}
func (ee Employee) changeName(newName string) {
	fmt.Println(ee.name, ee.position, ee.age, ee.sex)
	ee.name = newName
	fmt.Println(ee.name, ee.position, ee.age, ee.sex)
}

//  创意一个指针接收器的方法 ：语法为 func(t *结构体名称) 方法名(参数) (返回值){}
// 值接收器和指针接收器之间的区别在于，在指针接收器的方法内部的改变对于调用者是可见的
func (ee *Employee) changeAge(newAge int) {
	fmt.Println(ee.name, ee.position, ee.age, ee.sex)
	ee.age = newAge
	fmt.Println(ee.name, ee.position, ee.age, ee.sex)
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s", a.city, a.state)
}

func fullAddressfun(a address) {
	fmt.Printf("Full address: %s, %s", a.city, a.state)
}

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}
