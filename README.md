# go-vue-blog

#前言
>前后端分离个人博客，技术栈：vue.js + nuxt.js+ SCSS + ES6/7 +mysql+beego

##beego api接口

参数配置：
>//app.conf

MYSQL地址localhost

>dbhost = 127.0.0.1

MYSQL端口

>dbport = 3306

MYSQL用户名

>dbuser = root

MYSQL密码

>dbpassword = root

MYSQL数据库名称

>dbname = vue_blog


>//cd go-vue-blog

>bee run -gendoc=true -downdoc=true

-gendoc=true  生成swagger接口文档，
-downdoc=true  下载swagger自动化接口文档器

>http://localhost:8081/swagger/


#前端运行环境：

>cd go-vue-blog-front

>cnpm install 

>cnpm run dev

#环境依赖

>cnpm i --save axios sass-loader node-sass moment marked highlight.js

>cnpm run dev

#后台管理

>cd go-vue-blog-admin

>cnpm install

>cnpm i --save vue-router axios sass-loader node-sass  popmotion file-loader github-markdown-css marked prismjs moment marked

>cnpm run serve

Nuxt.js算是第一次接触，跟Vue.js的单页面应用的方式还是有点区别的，不过多看看文档和多搜索下问题不大


