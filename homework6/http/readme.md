
# Вопросы:

При формировании строки в шаблоне вида
 <p>Hello <strong>{{ .Name }}</strong>!</p> 

 значение из кода не передается
 func Hello1(w http.ResponseWriter, r *http.Request) {

	Name := r.FormValue("name")
	tmpl, _ := template.ParseFiles("static/hello.html")
	tmpl.Execute(w, Name)
	
}
передается в шаблон вида 
 <p>Hello <strong>{{ . }}</strong>!</p>

В чем может быть ошибка?