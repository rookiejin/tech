package seed

import (
	"math/rand"
)
var Titles = []string{
	"Laravel Vue 开发 SPA 应用" ,
	"Laravel 5.5 新特性",
	"Mac 从零开始配置开发环境",
	"PhpStorm 使用教程",
	"ES6 教程",
	"Laravel 实用 package 推荐",
	"从零部署一个网站",
	"PHP 单元测试",
	"Element UI 教程",
	"Laravel Vuejs 实战：开发知乎",
	"学习正则表达式",
	"Laravel 5.4 新特性",
	"Vue js 2.0 教程",
	"Mac开发大杂烩",
	"Webpack 教程",
	"用 Collection 重构代码",
	"Laravel和前端那些事",
	"Git 版本管理",
	"那些很cool的服务",
	"开发 Laravel 扩展",
	"Laravel 5.3 新特性",
}
var Descriptions = []string{
	"本系列视频会使用 Laravel 5.5 和 Vuejs 2.4 版本开发一个单页应用，其中会涉及前后端 API 交互，用户注册和登录，JWT 认证，vuex 状态管理等主要知识点。",
	"Laravel 5.5 正式版发布在即，又到了我们来学习新特性的时候了，本系列视频会详细介绍 Laravel 5.5 的各个新特性：比如验证规则，Package 自动发现，自定义 Blade 条件判断等。希望能够在你升级 Laravel 5.5 的时候可以帮上你的忙。",
	"本系列主要是从一个新的 Mac OS 一步一步配置一个符合 PHP 开发者的开发环境，包括一些个人推荐使用的软件和简单可用的配置。而且视频也会充分考虑 Laravel 开发者，将讲到包括命令行配置，phpstorm 的配置，Laravel Valet 安装和配置，Homestead 安装配置和使用流程等。",
	"PhpStorm 在我认为是最好的 PHP IDE，它拥有许多的特性比如：多点编辑，代码模板，xdebug 配合使用，测试集成，git 集成等。本系列视频将全面讲解 PhpStorm 的使用。",
	"es6 的系列教程会包括 es6 的诸多知识点：新的语法，箭头函数，类，作用域，this 关键字理解等。学完整个系列教程，必然可以掌握 es6 的核心知识点",
	"Laravel 社区其实非常活跃，在针对开发应用时需要使用的一些常见的 package，社区里基本都可以找到成熟的解决方案。这些 package 对于常见的应用场景来说，都是非常好的实践，本系列视频就说一些常用 package 的使用。",
	"Laravel 社区其实非常活跃，在针对开发应用时需要使用的一些常见的 package，社区里基本都可以找到成熟的解决方案。这些 package 对于常见的应用场景来说，都是非常好的实践，本系列视频就说一些常用 package 的使用。",
	"本系列视频意在完全从零开始教会大家怎么去部署一个别人也可以访问到的网站，其中的内容会包括购买域名，购买服务器，服务器设置，SSL 设置等各方面的知识",
	"ElementUI 是有饿了么前端团队开发的一套基于 Vue.js 2.0 的 UI 组件库，https://github.com/ElemeFE/element ，它兼顾了设计和交互的多个应用原则，可以帮助我们快速开发一个 Vue.js Web 应用。",
	"一个很好的学习思路就是：在学会了使用之后，再学习怎么扩展整个系统。所以在学习了如何使用 Laravel 之后，我们来学习怎么开发一个 Laravel 的扩展，包括写一下简单的单元测试和基础 CI 工具等",
}

var Thumb = []string {
	"",
}


func RandomString(n int) string{
	b := make([]byte ,n)
	for i:= 0 ; i < n ; i ++ {
		rnd := rand.Intn(122 - 97) + 97
		b[i] = byte(rnd)
	}
	return string(b)
}