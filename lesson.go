package main

import "fmt"

// メソッドとポインタレシーバーと値レシーバー & コンストラクタ & Embedded

type Vertex struct {
	x, y int
}

func (v Vertex) Area() int {
	return v.x * v.y
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

type Vertex3D struct {
	Vertex
	z int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

// non-structのメソッド

type MyInt int

func (i MyInt) Double() int {
	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", 1, 1)
	return int(i * 2)
}

// インターフェースとダックタイピング

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

// タイプアサーションとswitch type文

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

// Stringer

type PersonStringer struct {
	Name string
	Age  int
}

// ここでStringインターフェースに属する
func (p PersonStringer) String() string {
	return fmt.Sprintf("My name is %v.\n", p.Name)
}

// カスタムエラー

type UserNotFound struct {
	Username string
}

// ここでErrorインターフェースに属する
func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	// Something wrong
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}

func main() {
	// メソッドとポインタレシーバーと値レシーバー & コンストラクタ & Embedded

	v := New(3, 4, 5)
	v.Scale(10)
	fmt.Println(v.Area())
	fmt.Println(v.Area3D())

	// non-structのメソッド

	myInt := MyInt(10)
	fmt.Println(myInt.Double())

	// インターフェースとダックタイピング

	var mike Human = &Person{"Mike"}
	DriveCar(mike)

	var x Human = &Person{"X"}
	DriveCar(x)

	// var dog Dog = Dog{"dog"}
	// DriveCar(dog)

	// タイプアサーションとswitch type文

	do(10)
	do("Mike")
	do(true)

	// 型コンバージョンの実装令
	var i int = 10
	ii := float64(10)
	fmt.Println(i, ii)

	// Stringer

	mikeStringer := PersonStringer{"Mike", 22}
	fmt.Println(mikeStringer)

	// カスタムエラー

	e1 := &UserNotFound{"mike"}
	e2 := &UserNotFound{"mike"}
	fmt.Println(e1 == e2)
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}

}
