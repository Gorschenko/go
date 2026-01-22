<h2>Nil значения</h2>

<h3>Теория:</h3>

<ul>
    <li>
        Мапы. Нельзя писать, можно читать.
        var m map[string]int - nil мапа
        m := make(map[string]int) - исправленный вариант
    </li>
    <li>
        Каналы. Нельзя писать, нельзя читать. Будет полная блокировка.
        var ch chan int - nil канал
        go func() {
            for i := range 5 {
                ch <- i
            }
        }()
        for val := range ch {
            fmt.Println(val)
        }
        ch := make(chan int)
    </li>
    <li>
        Функция. Нельзя выполнять. Будет паника.
        var fn func() - nil функция
        fn() - паника
    </li>
    <li>
        Слайсы. Нельзя читать по индексу. Можно писать.
    </li>
</ul>

<h3>Задачи:</h3>
