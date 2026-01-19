<h2>Указатели</h2>

<h3>Теория:</h3>

<ul>
    <li>
        Все передается по значению и это значение копируется.
        func changeName(name string) {} - передается и копируется значение.
        func changeName(name *string) {} - передается и копируется ссылка на значение.
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
        </code></pre>
    </li>
</ul>

<h3>Задачи:</h3>

<details>
    <summary>
       Лекция 2.2. Вариант №1. Что будет выведено?
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
       Лекция 2.2. Вариант №2. Что будет выведено?
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
       Лекция 2.3. Вариант №3. Что будет выведено?
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
            fmt.Println(addr) // {New York Piccadilly 10}
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
        </pre></code>
    </details>
</details>
