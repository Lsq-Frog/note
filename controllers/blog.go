package controllers

import (
	"github.com/astaxie/beego"
)

type BlogController struct {
	beego.Controller
}

// @router /home [get]
func (b *BlogController) Home() {
	b.Layout = "layout.html"
	b.TplName = "index.html"
	b.Data["IsHome"] = true
	// b.LayoutSections = make(map[string]string)
	// b.LayoutSections["HtmlHead"] = "blogs/html_head.tpl"
	// b.LayoutSections["Scripts"] = "blogs/scripts.tpl"
	// b.LayoutSections["Sidebar"] = ""
}

// @router /support [get]
func (b *BlogController) Support() {
	b.Layout = "layout.html"
	b.TplName = "support.html"
	b.Data["IsSupport"] = true
}

// @router /about [get]
func (b *BlogController) About() {
	b.Layout = "layout.html"
	b.TplName = "about.html"
	b.Data["IsAbout"] = true
}

// @router /blog [get]
func (b *BlogController) Blog() {
	b.Layout = "layout.html"
	b.TplName = "blog.html"
	b.Data["IsBlog"] = true
}

// @router /contact [get]
func (b *BlogController) Contact() {
	b.Layout = "layout.html"
	b.TplName = "contact.html"
	b.Data["IsContact"] = true
}
