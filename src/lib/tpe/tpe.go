package tpe

import(
    "config"
    "html/template"
)

//创建视图
func Assign(tpl string)*template.Template{
    //调用模版
	view, err := template.ParseFiles(
        tpl,
        config.Get("ROOT_PATH") + "static/layout/head.html",
        config.Get("ROOT_PATH") + "static/layout/top.html",
        config.Get("ROOT_PATH") + "static/layout/menu.html",
        config.Get("ROOT_PATH") + "static/layout/foot.html")
	if err != nil {
		panic("tpe->Assign Error.")
	}
	return view
}