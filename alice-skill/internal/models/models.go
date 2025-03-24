/*
Сейчас в навыке для Алисы мы никак не обрабатываем входящий запрос
от пользователя, а ответ передаём в виде предопределённых байт.
Чтобы это исправить, добавим модели запроса-ответа и обработаем данные
с помощью сериализаторов.
*/

package models

const (
	TypeSimpleUtterance = "SimpleUtterance"
)

type Request struct {
	Request SimpleUtterance `json:"request"`
	Version string          `json:"version"`
}

type SimpleUtterance struct {
	Type    string `json:"type"`
	Command string `json:"command"`
}

type Response struct {
	Response ResponsePayload `json:"response"`
	Version  string          `json:"version"`
}

type ResponsePayload struct {
	Text string `json:"text"`
}
