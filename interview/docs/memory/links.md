<h2>Указатели</h2>

<h3>Теория:</h3>

<ul>
    <li>
        Все передается по значению и это значение копируется. Как В функцию, так и ИЗ функции.
        func changeName(name string) {} - передается и копируется значение.
        func changeName(name *string) {} - передается и копируется ссылка на значение.
    </li>
    <li>
        Слайсы, маппы и каналы передаются в функцию по ссылке.
    </li>
    <li>
        defer выполняется после return, но до того как переменной, которая возвращается из функции, будет присвоено новое значение. 
    </li>
    <li>
        <pre><code>
            func changeSlice(s1 *[]int) {
                *s = append(*s1, 4) // так мы работаем с исходным slice
            }
        </code></pre>
    </li>
</ul>

<h3>Задачи:</h3>

<details>
    <summary>
       Лекция 2.2. Вариант №1
    </summary>
    <pre><code>
        func changeName(name *string) {
            newName := "Alice" 
            name = &newName
        }
        func main() {
            name := "Bob"
            changeName(&name)
            fmt.Println(name)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            func changeName(name *string) {
                // name - копия переданной ссылки
                newName := "Alice" 
                name = &newName // изменили адрес ссылки (переменная name), которая явлется копией
                // *name = "Alice" // так сработает, так как name имеет такую же ссылку на значение, как переменная name в функции main
            }
            func main() {
                name := "Bob"
                changeName(&name)
                fmt.Println(name) // Bob, так как в функции changeName работали с копией ссылки
            }
        </pre></code>
    </details>
</details>

<details>
    <summary>
       Лекция 2.2. Вариант №2
    </summary>
    <pre><code>
        type Person struct {
            name string
            age int
        }
        func (p *Person) changeName(name string) {
            p = &Person{
                name: name
            }
        }
        func main() {
            p := Person{
                name: "Bob",
                age: 25,
            }
            p.changeName("Alice")
            fmt.Println(p)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            type Person struct {
                name string
                age int
            }
            func (p *Person) changeName(name string) {
                // p - копия на ссылку
                // p присвается ссылка на новую структуру Person,
                // но это никак не меняет значение p, которая объявлена в main.
                p = &Person{
                    name: name
                }
                // Исправление
                // p.name = name // мутация значения переменной, переданной по ссылку можно
            }
            func main() {
                p := Person{
                    name: "Bob",
                    age: 25,
                }
                p.changeName("Alice")
                fmt.Println(p) // {Alice 25}
            }
        </pre></code>
    </details>
</details>

<details>
    <summary>
       Лекция 2.3. Вариант №3
    </summary>
    <pre><code>
        type Address struct {
            city   string
            street string
            house  int
        }
        func (a *Address) setCity(city string) {
            a = &Address{
                city: city,
            }
        }
        func (a *Address) setStreet(street string) {
            a.street = street
        }
        func setHouse(addr *Address, house int) {
            addr = &Adrress{
                house: house,
            }
        }
        //
        func main() {
            addr := Address{
                city: "New York",
                street: "Broadway",
                house: 10,
            }
            addr.setCity("London")
            addr.setStreet("Piccadilly")
            setHouse(&addr, 5)
            fmt.Println(addr)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            type Address struct {
                city   string
                street string
                house  int
            }
            func (a *Address) setCity(city string) {
                a = &Address{
                    city: city,
                }
            }
            func (a *Address) setStreet(street string) {
                a.street = street
            }
            func setHouse(addr *Address, house int) {
                addr = &Adrress{
                    house: house,
                }
            }
            //
            func main() {
                addr := Address{
                    city: "New York",
                    street: "Broadway",
                    house: 10,
                }
                addr.setCity("London") // Не изменяется, так как идет работа с копией ссылки
                addr.setStreet("Piccadilly") // Изменяется, так как идет работа с мутабильным значением копии ссылки
                setHouse(&addr, 5) // Не изменяется, так как идет работа с копией ссылки
                fmt.Println(addr) // {New York Piccadilly 10}
            }
        </pre></code>
    </details>
</details>

<details>
    <summary>
       Лекция 2.4. Вариант №4
    </summary>
    <pre><code>
        func changeSlice(s []int) {
            s[1] = 10
            s2 = append(s, 4)
            fmt.Println(s2)
        }
        func main() {
            s := []int{1, 2, 3}
            changeSlice(s)
            fmt.Println(s)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            func changeSlice(s []int) {  // значение копируется и передается по ссылке
                s[1] = 10 // l=3, c=3, based on s
                s2 = append(s, 4) // l=4, c=6, new [], s2=[1,2,3,4]
                fmt.Println(s2)
            }
            func main() {
                s := []int{1, 2, 3} // l=3, c=3, [1,2,3]
                changeSlice(s)
                fmt.Println(s)
            }
        </pre></code>
    </details>
</details>

<details>
    <summary>
       Лекция 2.4. Вариант №5
    </summary>
    <pre><code>
        func changeSlice(s []int) {
            s = append(s, 4)
        }
        func main() {
            s := make([]int, 3, 4)
            for i:= range 3 {
                s[i] = i + 1
            }
            changeSlice(s)
            s1 := s[0:4]
            fmt.Println(s1)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            func changeSlice(s1 []int) {
                s = append(s1, 4) // l=3, c=4, based on s1, s=[1,2,3,4], s1=[1,2,3],4
            }
            func main() {
                s := make([]int, 3, 4) // l=3, c=4, s=[0,0,0],0
                for i:= range 3 {
                    s[i] = i + 1 // l=3, c=4, s=[1,2,3],0
                }
                changeSlice(s)
                s1 := s[0:4] // l=4,c=4, based on s, [1,2,3,4]
                fmt.Println(s1)
            }
        </pre></code>
    </details>
</details>

<details>
    <summary>
       Лекция 2.5. Вариант №6
    </summary>
    <pre><code>
        changeMap(m map[string]int) {
            m["Bob"] = 52
        }
        func main() {
            s := make(map[string]int)
            s["Bob"] = 25
            changeMap(s)
            fmt.Println(s)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            changeMap(m map[string]int) {
                m["Bob"] = 52
            }
            func main() {
                s := make(map[string]int)
                s["Bob"] = 25
                changeMap(s)
                fmt.Println(s) // map[Bob:52]
            }
        </pre></code>
    </details>
</details>

<details>
    <summary>
       Лекция 2.6. Вариант №7
    </summary>
    <pre><code>
        type Person struct {
            name string
            age int
        }
        func changeName(p *Person, name string) {
            p.name = name
        }
        func getPerson() Person {
            var p Person
            defer changeName(&p, "Alice")
            p = Person{
                name: "Bob",
                age: 52,
            }
            return p
        }
        func main() {
            p := getPerson()
            fmt.Println(p)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            type Person struct {
                name string
                age int
            }
            func changeName(p *Person, name string) {
                p.name = name
            }
            func getPerson() Person {
                var p Person
                defer changeName(&p, "Alice")
                p = Person{
                    name: "Bob",
                    age: 52,
                }
                return p
            }
            func main() {
                p := getPerson()
                fmt.Println(p) { Bob 52 }
            }
        </pre></code>
    </details>
</details>

<details>
    <summary>
       Лекция 2.6. Вариант №8
    </summary>
    <pre><code>
        func changeSlice(s []int) {
            s[1] = 10
        }
        func getSlice() []int {
            s := []int{1, 2, 3}
            defer changeSlice(s)
            return s
        }
        func main() {
            p := getSlice()
            fmt.Println(p)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            func changeSlice(s []int) {
                s[1] = 10  
            }
            func getSlice() []int {
                s := []int{1, 2, 3} // [1,2,3]
                defer changeSlice(s)
                return s 
            }
            func main() {
                p := getSlice() // based on s, slice=[1,10,3]
                fmt.Println(p)
            }
        </pre></code>
    </details>
</details>

<details>
    <summary>
       Лекция 2.6. Вариант №9
    </summary>
    <pre><code>
        func named() (a, b int) {
            a, b = 1, 2
            defer func() {
                a = 10
                b = 20
            }()
            return a, b 
        }
        func unnamed() (int, int) {
            a, b := 1, 2
            defer func() {
                a = 10
                b = 20
            }()
            return a, b
        }
        func main() {
            a, b := named()
            fmt.Println(a, b)
            a, b = unnamed()
            fmt.Println(a, b)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            func named() (a, b int) {
                a, b = 1, 2
                defer func() {
                    a = 10
                    b = 20
                }()
                return a, b // 1,2
            }
            func unnamed() (int, int) {
                a, b := 1, 2
                defer func() {
                    a = 10
                    b = 20
                }()
                return a, b // 1,2
            }
            func main() {
                a, b := named() // 10,20 
                fmt.Println(a, b)
                a, b = unnamed() // 1,2
                fmt.Println(a, b)
            }
        </pre></code>
    </details>
</details>
