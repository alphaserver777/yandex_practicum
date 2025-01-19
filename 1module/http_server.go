package main

import "net/http"

func mainPage(res http.ResponseWriter, req *http.Request) {
	/*
	   Здесь у нас есть два параметра:

	   •   res http.ResponseWriter
	   •   req *http.Request

	   Зачем нужен тип http.ResponseWriter для параметра res?

	   1.  Определение функциональности:
	       *   http.ResponseWriter — это интерфейс, который определяет набор методов (действий), которые может выполнять любой тип данных, реализующий этот интерфейс.
	       *   http.ResponseWriter имеет методы, такие как Write(), WriteHeader(), Header() и другие, которые необходимы для отправки HTTP-ответа клиенту (например, браузеру).
	       *   Указав тип параметра res как http.ResponseWriter, мы гарантируем, что внутри функции mainPage мы можем использовать методы, необходимые для формирования ответа.

	   2.  Безопасность типов:
	       *   Тип параметра http.ResponseWriter говорит Go, что мы ожидаем получить именно этот тип, а не что-то другое.
	       *   Если бы мы попытались передать в параметр res что-то, что не реализует интерфейс http.ResponseWriter, то Go выдал бы ошибку компиляции. Это помогает нам избежать ошибок во время выполнения программы.

	   3.  Управление ответом:
	       *   http.ResponseWriter предоставляет нам доступ к "тарелке" (как мы это называли ранее), на которую мы можем "положить" данные для отправки клиенту.
	       *   Мы можем:
	           *   Записывать данные в тело ответа (с помощью res.Write()).
	           *   Устанавливать HTTP-заголовки (например, Content-Type: text/plain).
	           *   Устанавливать код ответа (например, 200 OK, 404 Not Found).
	       *   Без типа параметра http.ResponseWriter у нас не было бы возможности формировать ответ.

	   Зачем нужен тип *http.Request для параметра req?

	   1.  Доступ к информации о запросе:
	       *   *http.Request — это указатель на структуру http.Request, которая содержит всю информацию о запросе от клиента.
	       *   Через req мы можем получить доступ к:
	           *   URL-адресу запроса.
	           *   HTTP-методу (GET, POST, PUT, DELETE и т.д.).
	           *   Заголовкам запроса.
	           *   Параметрам запроса (например, ?name=John&age=30).
	           *   Телу запроса (например, данные, отправленные через форму).

	   2.  Обработка запроса:
	       *   Информация из *http.Request позволяет нам правильно обработать запрос.
	       *   Мы можем определить, какой ресурс запрашивает клиент, какие данные он передает, и в соответствии с этим сформировать ответ.
	       *   Например, мы можем извлечь имя пользователя из параметров запроса и использовать его в ответе.
	       *   Без типа параметра *http.Request мы не знали бы, что именно запросил клиент.

	   3.  Безопасность типов:
	       *   Тип *http.Request гарантирует, что мы можем безопасно работать со структурой запроса. Go проверяет, что req действительно является указателем на http.Request.

	   В контексте веб-сервера:

	   •   http.ResponseWriter и *http.Request — это ключевые типы, которые позволяют Go создавать веб-серверы.
	   •   Когда веб-сервер получает запрос, он создает объекты типа http.ResponseWriter и *http.Request и передает их в функцию-обработчик (например, mainPage).
	   •   Функция-обработчик использует эти объекты для формирования и отправки ответа.
	   •   Типы параметров обеспечивают безопасность, гибкость и управляемость работы с HTTP-запросами и ответами.

	   Простыми словами:

	   •   http.ResponseWriter — это как тарелка для ответа. Тип говорит нам, что это именно "тарелка" для HTTP-ответа, на которую можно положить данные, чтобы отправить их клиенту.
	   •   *http.Request — это как заказ, который пришел от клиента. Тип говорит нам, что это именно "заказ" типа http.Request , из которого мы можем получить информацию о том, что хочет клиент.

	   В итоге, типы параметров в func mainPage(res http.ResponseWriter, req *http.Request):

	   •   Указывают, какие типы данных ожидает эта функция (тарелка и заказ от клиента)
	   •  Позволяют нам использовать методы этих типов для отправки ответа и обработки запроса.
	   •  Обеспечивают безопасность кода за счет проверки типов\
	*/
	res.Write([]byte("Привет!"))
	/*
	   Здесь res.Write([]byte("Привет!")) — это единственная строка кода, которая выполняет основную задачу функции: отправить ответ клиенту.

	   Разберем res.Write([]byte("Привет!")) по частям:

	   1.  res:
	   		•   Как мы уже обсуждали, res — это параметр типа http.ResponseWriter, который передается в функцию mainPage.
	   		•   http.ResponseWriter — это интерфейс, который представляет собой "тарелку" для ответа. Именно через этот объект мы можем отправить данные клиенту.

	   2.  .Write(...):
	   		•   .Write — это метод, который вызывается на объекте res (http.ResponseWriter). (Он есть внутри интерфейса, если посмотреть внутрь)
	   		•   Метод .Write предназначен для записи данных в тело HTTP-ответа.
	   		•   Это как взять ложку и положить что-то на "тарелку".

	   3.  ([]byte("Привет!")):
	   		•   "Привет!" — это текстовая строка (string), которую мы хотим отправить в ответе.
	   		•   []byte(...) — это функция, которая преобразует строку "Привет!" в срез байтов ([]byte).
	   		•   Почему нужны именно байты? Потому что интернет (и HTTP) передают данные в виде байтов.
	   		•   Байт - это единица информации, состоящая из 8 бит.
	   		•   Текст "Привет!" - это набор символов, каждому символу соответствует определенный код, а код можно представить в виде байтов.

	   Зачем всё это нужно?

	   1.  Отправить ответ:
	   		•   res.Write([]byte("Привет!")) — это единственный способ отправить ответ клиенту (например, браузеру).
	   		•   Клиент отправил HTTP-запрос (например, открыл страницу в браузере), и мы должны отправить ему ответ.
	   		•   Без res.Write функция mainPage ничего бы не сделала, и браузер не получил бы никакого контента.

	   2.  Создать HTTP-ответ:
	   		•   HTTP-ответ состоит из двух частей:
	   				*   Заголовки: Метаданные о ответе (например, тип контента, код состояния).
	   				*   Тело: Содержание ответа (например, HTML-код, текст, изображение).
	   		•   res.Write() записывает данные в тело ответа.
	   		•   По умолчанию Go установит заголовок Content-Type в text/plain, поскольку мы отправляем простой текст.

	   3.  Показать что-то пользователю:
	   		•   Наше сообщение "Привет!" - это именно то, что мы хотим показать пользователю.
	   		•   Если бы мы отправили другой текст (например, res.Write([]byte("Hello, world!"))), то пользователь увидел бы "Hello, world!" в браузере.

	   Почему байты?

	   •   Как уже говорилось, интернет работает с байтами.
	   •   HTTP (протокол, по которому работают веб-серверы) передает данные в виде потока байтов.
	   •   Чтобы браузер мог понять, что именно мы ему отправляем (текст, HTML, изображение, и т.д.), мы должны отправить ему байты, а не просто строки.

	   Аналогия:

	   Представьте, что res — это тарелка, которую вы несете клиенту в кафе.

	   •   res.Write(...) — это как вы кладете на тарелку блюдо.
	   •   ([]byte("Привет!")) — это как само блюдо (текст "Привет!", преобразованный в вид, который понимает HTTP), которое вы кладете на тарелку.

	   В итоге:

	   res.Write([]byte("Привет!")) — это ключевая строка кода, которая отправляет клиенту сообщение "Привет!". Без этой строки функция mainPage была бы бесполезной. Она отвечает за то, чтобы веб-сервер правильно формировал и отправлял ответ клиенту. Она преобразует текст в массив байт и записывает в "тарелку" res, которую потом веб-сервер передает клиенту.
	*/
}
func apiPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Это страница /api."))
}

func main() {
	mux := http.NewServeMux()
	/*
	   Создает новый мультиплексор (или "роутер") для обработки HTTP-запросов.
	   Мультиплексор (mux) — это объект, который решает, какую функцию вызвать в зависимости от пути запроса (например, /api/ или /).

	   Аналогия:
	   Представьте, что мультиплексор — это диспетчер в ресторане. Когда клиент (запрос) приходит, диспетчер решает, к какому официанту (функции) его направить.

	   На практике рекомендуется использовать свою переменную-маршрутизатор функцией NewServeMux() *ServeMux и вызвать для неё методы HandleFunc() с маршрутами и обработчиками.
	*/
	mux.HandleFunc(`/api/`, apiPage)
	/*
	Регистрирует функцию apiPage для обработки запросов, которые начинаются с /api/.
	Например, если кто-то отправит запрос на http://localhost:8080/api/, сервер вызовет функцию apiPage.
	Аналогия:
	Диспетчер говорит: "Если клиент хочет заказать десерт (путь /api/), направь его к официанту apiPage."
	*/
	mux.HandleFunc(`/`, mainPage)

	err := http.ListenAndServe(`:8080`, mux)
	/*
	   Запускает веб-сервер на порту 8080 и использует мультиплексор mux для обработки запросов.
	   Аналогия:
	   Диспетчер (mux) говорит: "Я буду обслуживать клиентов (запросы) на порту 8080. Если кто-то придет, я буду решать, куда его направить, используя свой внутренний справочник (mux.HandleFunc)."
	*/
	if err != nil {
		panic(err)
	}
}
