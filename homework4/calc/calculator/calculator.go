package calculator

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"math"
	"strconv"
	"strings"
)

func Calculate(expr string) (float64, error) {
	// Поиск управляющих операторов в строке и построение дерева
	// Вернет корневой элемент в переменную root все математические действия, построенные
	// в виде дерева
	// Node будет содержать в себе действие для вычисления
	// exprNode() даст доступ к ветвям дерева, исходящим из корня/ноды
	//
	// type Expr interface {
	//  Node
	//  exprNode()
	// }
	root, err := parser.ParseExpr(expr)

	if err != nil {
		return -1, err
	} else {
		return eval(root)
	}
}

type Func struct {
	Name string
	Args int
	Func func(args ...float64) float64
}

var funcMap map[string]Func

func Help() {
	fmt.Println("Справка по работе с калькулятором:")
	fmt.Printf("\tОперации:\n")
	fmt.Printf("\t\t сложение: x+y \n")
	fmt.Printf("\t\t вычитание: x-y \n")
	fmt.Printf("\t\t умножение: x*y \n")
	fmt.Printf("\t\t деление: x/y \n")
	fmt.Printf("\t\t возведение в степень: x^y \n")
	fmt.Printf("\t\t остаток от деления: x%sy \n", "%")
	fmt.Printf("\t\t побитовое И: x&y\n")
	fmt.Printf("\t\t побитовое ИЛИ: x|y\n")
	fmt.Printf("\tФункции:\n")
	fmt.Printf("\t\t квадратный корень: sqrt(x)\n")
	fmt.Printf("\t\t абсолютное значение: abs(x)\n")
	fmt.Printf("\t\t логарифм числа X по основанию Y: log(y,x)\n")
	fmt.Printf("\t\t натуральный логарифм: ln(x)\n")
	fmt.Printf("\t\t синус: sin(x)\n")
	fmt.Printf("\t\t косинус: cos(x)\n")
	fmt.Printf("\t\t тангенс: tan(x)\n")
	fmt.Printf("\t\t арксинус: arcsin(x)\n")
	fmt.Printf("\t\t арккосинус: arccos(x)\n")
	fmt.Printf("\t\t арктангенс: arctan(x)\n")
	fmt.Printf("\t\t максимальнок число: max(x, y)\n")
	fmt.Printf("\t\t минимальное число: min(x, y)\n")
	fmt.Printf("\t\t число ПИ: pi\n")
	fmt.Printf("\t\t золотое сечение: phi\n")
	fmt.Printf("\t\t основание натурального логарифма: е\n")
}

// Создание массива функций для обработки данных
func init() {
	funcMap = make(map[string]Func)
	funcMap["sqrt"] = Func{
		Name: "sqrt",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Sqrt(args[0])
		},
	}
	funcMap["abs"] = Func{
		Name: "abs",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Abs(args[0])
		},
	}
	funcMap["log"] = Func{
		Name: "log",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Log(args[0]) / math.Log(args[1])
		},
	}
	funcMap["ln"] = Func{
		Name: "ln",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Log(args[0])
		},
	}
	funcMap["sin"] = Func{
		Name: "sin",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Sin(args[0])
		},
	}
	funcMap["cos"] = Func{
		Name: "cos",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Cos(args[0])
		},
	}
	funcMap["tan"] = Func{
		Name: "tan",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Tan(args[0])
		},
	}
	funcMap["arcsin"] = Func{
		Name: "arcsin",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Asin(args[0])
		},
	}
	funcMap["arccos"] = Func{
		Name: "arccos",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Acos(args[0])
		},
	}
	funcMap["arctan"] = Func{
		Name: "arctan",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Atan(args[0])
		},
	}
	funcMap["max"] = Func{
		Name: "max",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Max(args[0], args[1])
		},
	}
	funcMap["min"] = Func{
		Name: "min",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Min(args[0], args[1])
		},
	}
}

// Разбор полученных данных
func eval(expr ast.Expr) (float64, error) {
	switch expr.(type) {
	case *ast.BasicLit:
		return basic(expr.(*ast.BasicLit))
	// Бинарные выражения
	case *ast.BinaryExpr:
		return binary(expr.(*ast.BinaryExpr))
	// Вложенные вычисления
	case *ast.ParenExpr:
		return eval(expr.(*ast.ParenExpr).X)
	// Обработка посредством математических функций
	case *ast.CallExpr:
		return call(expr.(*ast.CallExpr))
	// Случай для обработки констант
	case *ast.Ident:
		return ident(expr.(*ast.Ident))
	default:
		return -1, errors.New("Не удалось распознать оператор")
	}
}

// Разбор чисел
func basic(lit *ast.BasicLit) (float64, error) {
	switch lit.Kind {
	case token.INT:
		i, err := strconv.ParseInt(lit.Value, 10, 64)

		if err != nil {
			return -1, err
		} else {
			return float64(i), nil
		}
	case token.FLOAT:
		i, err := strconv.ParseFloat(lit.Value, 64)

		if err != nil {
			return -1, err
		} else {
			return i, nil
		}
	default:
		return -1, errors.New("Неизвестный аргумент")
	}
}

func binary(expr *ast.BinaryExpr) (ret float64, err error) {
	x, err1 := eval(expr.X)
	y, err2 := eval(expr.Y)
	ret = -1

	if (err1 == nil) && (err2 == nil) {

		switch expr.Op {
		case token.ADD:
			ret = x + y //СЛОЖЕНИЕ
		case token.SUB:
			ret = x - y // вычитание
		case token.MUL:
			ret = x * y // умножение
		case token.QUO:
			ret = x / y //деление
		case token.REM:
			ret = float64(int64(x) % int64(y)) //остаток от деления
		case token.AND:
			ret = float64(int64(x) & int64(y)) // И
		case token.OR:
			ret = float64(int64(x) | int64(y)) // ИЛИ
		case token.XOR:
			ret = math.Pow(x, y) //возведение в степень
		default:
			err = errors.New("Неизвестный бинарный оператор")
		}
	} else {
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
	}

	return
}

// Обработка констант
func ident(id *ast.Ident) (float64, error) {
	switch n := strings.ToLower(id.Name); n {
	case "pi":
		return math.Pi, nil
	case "e":
		return math.E, nil //основание натурального log
	case "phi":
		return math.Phi, nil //золотое сечение
	default:
		return -1, errors.New("Неизвестная константа " + n)
	}
}

// Обработка функциональных операторов с помощью созданного массива функций для обработки
func call(c *ast.CallExpr) (float64, error) {
	switch t := c.Fun.(type) {
	case *ast.Ident:
	default:
		_ = t
		return -1, errors.New("Неизвестный тип функции")
	}

	ident := c.Fun.(*ast.Ident)

	args := make([]float64, len(c.Args))
	for i, expr := range c.Args {
		var err error
		args[i], err = eval(expr)
		if err != nil {
			return -1, err
		}
	}

	name := strings.ToLower(ident.Name)

	if val, ok := funcMap[name]; ok {
		if len(args) == val.Args {
			return val.Func(args...), nil
		} else {
			return -1, errors.New("Слишком много аргументов для " + name)
		}
	} else {
		return -1, errors.New("Неизвестная функция " + name)
	}
}
