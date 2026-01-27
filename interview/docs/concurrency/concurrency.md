<h2>Конкурентность</h2>

<h3>Теория:</h3>

<ul>
    <li>
        Конкурентность объединяет в себя два понятия - асинхронный код и параллелизм. Планировщик сам решает как ему запускать код: в одном поток (асинхронный код) или выделить новый поток (параллелизм). Это называется fork/join моделью.
    </li>
    <li>
        В конкурентности есть три проблемы: синхронизация, гонка и утчека гоурутин.
        Синхронизация - ожидания выполнения всех гоурутин.
        Гонка - возникает когда несколько гоурутин работают с общими данными. Решение: использовать атомик, мютекс, канал
    </li>
    <li>
        Атомик - низкоуровневый примитив. На них построены мютексы и каналы.
        Мютекс - блокирует доступ к переменной, если с ней работает несколько горутин, что позволяет работать атомарно.
        Канал - используется как для синхронизации, так и для доступа к общим ресурсам.
    </li>
    <li>
        <pre><code>
            func foo(wg *sync.WaitGroup) {
                defer wg.Done()
                fmt.Println("Hello from foo")
            }
            func main() {
                wg := &sync.WaitGroup{} // нужно для синхронизации гоурутин
                wg.Add(1) // добавляем гоурутину в группу
                go func() {
                    defer wg.Done() // сигнализируем о том, что гоурутина закончилась
                    fmt.Println("Hello from gouroutine")
                }()
                wg.Add(1)
                go foo()
                wg.Wait() // Ждем выполнение всей группы
                fmt.Println("Hello from main")
            }
        </code></pre>
    </li>
    <li>
        <pre><code>
            func main() {
                // Гоурутины неправильно работают с общими данными из-за того, что money++ является неатомарной операцией
                money := 0
                wg := &sync.WaitGroup{}
                wg.Add(1000)
                for range 1000 {
                    go func() {
                        defer wg.Done()
                        money++
                    }()
                }
                wg.Wait()
                fmt.Println(money)
            }
        </code></pre>
    </li>
    <li>
        <pre><code>
            func main() {
                // Гонка. Исправленный вариант. Атомики
                var money atomic.Int32
                wg := &sync.WaitGroup{}
                wg.Add(1000)
                for range 1000 {
                    go func() {
                        defer wg.Done()
                        money.Add(1)
                    }()
                }
                wg.Wait()
                fmt.Println(money.Load())
            }
        </code></pre>
    </li>
    <li>
        <pre><code>
            func main() {
                // Гонка. Исправленный вариант. Мютекс
                mx := &sync.Mutex{}
                wg := &sync.WaitGroup{}
                wg.Add(1000)
                for range 1000 {
                    go func() {
                        defer wg.Done()
                        mx.Lock()
                        money++
                        mx.Unlock()
                    }()
                }
                wg.Wait()
                fmt.Println(money)
            }
        </code></pre>
    </li>
</ul>

<h3>Задачи:</h3>
