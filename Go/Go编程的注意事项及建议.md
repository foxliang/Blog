# Go编程的注意事项及建议

接下来我们会说几个PHP程序员在刚开始用Go写程序时几个需要改变的编码习惯和要注意的地方。

## 尽量使用结构体切片代替字典

我们有的新同学特别爱使用Go里面的Map，有的时候还是切片里边套Map，比如我看一开始有的同学把一些配置信息放在map[string]string类型的Map里，多个的话再把Map放进切片里，比如这样。
```
var configMap = []map[string]string{
    {
        "stockNum": "100",
        "name":     "芒果TV周卡",
        "type":     "virtual",
    },
}
```

后面程序使用的时候再去用键去取值，这么做程序当然能实现，但你会发现Go里面因为是强类型，你在用上面字典里面的数值时还得对他们做类型转换。很多同学马上会说，那我把Map的类型换成map[string]interface{}，我只能说你试试，看你用的时候Go让不让你做类型断言。
这其实是涉及一个思维的转变，那么在像Go这样的强类型语言里针对这种情况该怎么办呢？这就需要让我们养成先定义结构体类型后使用的习惯了，比如像上面的情况我就可以先定义一个类型。
```
type Product struct {
    StockNum  int64
    Name      string
    Type      string
}

var configs = []*Product {
    {
        StockNum: 100,
        Name: "芒果TV周卡",
        Type: "virtual",
    },
  ......
}
```

这么做就能避免像上面那样使用StockNum前还得把它转成整型的问题了，而且编辑器还能做类型提示，不需要你刻意记得Map里的键，还能避免你一时疏忽把键拼错导致BUG的尴尬。

除了上面说的还有人喜欢在返回值里返回Map，这种写法除了会导致上面说的那样问题，让别人使用起来也特别不方便。比如我要用你的方法我还得进去看看你的代码里这个Map到底有哪些键。

所以我们写Go代码时，其实Map的使用率要比在PHP里使用数组低很多，很多时候都是用结构体以及结构体切片的，对于那种key为数据ID，值为数据Map的这种映射，也是改成Key为数据ID，值为数据自己定义的类型才对。比如下面这个Map类型的变量，它的Key是产品的ID，值的类型是我们上面定义的Product结构体
```
var productMap = map[int64]*Product {
    123:    {
        StockNum: 100,
        Name: "芒果TV周卡",
        Type: "virtual",
    },
}
```

### 针对这部分说的这个问题我觉得记住："根据数据先定类型再使用"这个原则就行了。

说完这个在代码里出现率最高的问题后，下面我们再说几个写Go代码时的要注意的细节。

## 零值陷阱
未进行初始化的变量默认值为其类型的零值，需要注意的是slice，map，chan和*T类型对应的零值是nil。

这些类型的变量在未初始化前是无法在程序里直接使用的，有些情况下会导致运行时错误。

常见的两种运行时错误是：
```

panic: assignment to entry in nil map


panic: invalid memory address or nil pointer dereference
```

第一个错误是因为对一个未初始化的map进行赋值导致的，所以使用map类型的变量前要记得用make函数对变量进行初始化，与map类似的切片在使用append函数 向nil slice追加新元素就可以，原因是append函数会生成新的切片，在底层为切片分配了底层数组。

第二个错误是对nil指针进行了解引用导致的，指针的零值nil与*T{}并不相等。所以指针类型的变量在使用前要注意使用new函数进行初始化。

还有就是前端同学们非常不喜欢接口返回值的字段有数据的时候是个列表，没数据的时候是Null，这也是切片未初始化导致的，如果数据库里没查到数据，那么在代码逻辑里就执行不到给切片append数据的循环里，所以就会出现这个问题。这是一个保持接口字段类型一致性的一个很重要的细节。

使用error返回函数错误

在使用PHP时，函数的错误是通过抛出异常，甚至是通过返回0，false之类的值来表示函数遇到的错误（这种，即使写PHP也不推荐这种做法）

比如好的写法，可这样写：
```
public function updateUserFavorites(User $user, $favoriteData)
{
    try {
        // database execution
                ......
    } catch (QueryException $queryException) {
        throw new UserManageException(func_get_args(), 'Error Message', '501' , $queryException);
    }

    return true;
}
```

但很多的人会这么写：
```
public function updateUserFavorites(User $user, $favoriteData)
{
    // database execution
        if ($conn.AffectedRows <= 0) {
        return false
    }

    return true;
}
```
在Go语言里虽然没有异常机制，但是可以让函数返回error明确遇到的错误。所以除非确定函数不需要返回error，多数情况下我们的函数都是需要返回error的，所以在定义函数时要明确，返回的数据和error的区别，两种返回值的职责范围不一样。要通过函数返回的error是否为空，而不是返回数据是0或者false之类的值判断函数是否执行成功。

谨慎使用map[string]interface{}做参数

写过PHP的同学都知道，PHP里的数组近乎万能，可以用来当列表、字典，而且当字典用时还能保证字典key的遍历顺序，这点是很多语言的字典类型办不到的事情。

很多刚从PHP转到用Go开发的同学还是带着在PHP里使用数组参数的习惯，那么在Go语言里，最像PHP数组的可能就是map[string]interface{}了。

这种还是典型的动态语言编程的思维，在使用Go的时候，针对比较复杂的代表一类事物的参数，我们也是应该先定义结构体，然后使用结构体指针或者结构体指针切片作为参数。尽量不使用map[string]interface{}这种类型的参数，IDE也没法帮助提示这些参数的内部结构，这让其他人使用这个代码时就会很苦恼，还得先看看函数实现里具体用到了字典的哪些键。比如下面这两个函数的对比：
```
type UserInput struct{
        Name     string
        Age      int32
        Password string
}
func AuthenticateUser(input *UserInput) error {
    findUser(input.Name, input.Password)
    ...
}

func DummyAuthenticateUser(input map[string]interface{}) error {
    findUser(input["name"], input["password"])
    ...
}
```
一般在业务级别的程序开发里，我们要传递存储在数据表里的额外信息的时候才会使用到map[string]interface{}类型的参数。写表前把这部分数据编码成JSON格式再写入，当然这个主要看使用场景，凡事没有绝对，这里只是强调一些在编码习惯上的问题。
