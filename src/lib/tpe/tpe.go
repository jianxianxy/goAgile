package tpe

import(
    "config"
    "html/template"
)

//创建视图
func Assign(tpl string)*template.Template{
    //调用模版
	view, err := template.ParseFiles(
        config.Get("ROOT_VIEW") + tpl,
        config.Get("ROOT_VIEW") + "layout/head.html",
        config.Get("ROOT_VIEW") + "layout/top.html",
        config.Get("ROOT_VIEW") + "layout/menu.html",
        config.Get("ROOT_VIEW") + "layout/foot.html")
	if err != nil {
		panic("tpe->Assign Error.")
	}
	return view
}