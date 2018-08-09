/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80011
 Source Host           : localhost:3306
 Source Schema         : vue_blog

 Target Server Type    : MySQL
 Target Server Version : 80011
 File Encoding         : 65001

 Date: 10/08/2018 04:16:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标题',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '文章内容',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `publish_time` datetime(0) NULL DEFAULT NULL COMMENT '发布时间',
  `status` enum('DRAFT','PUBLISHED','OFFLINE') CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文章状态(DRAFT: 草稿， PUBLISHED: 发布，OFFLINE: 下线)',
  `category` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '分类',
  `tag` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标签',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES (1, '基于Vue2开发的读书WebAPP', '\n## 前言\n\n初学vue.js，官网的文档写的很清楚，看着不难。俗话说：光说不练假把式。程序猿学个新东西还是要敲出来看看效果比较好。最开始是想搞个音乐类的，毕竟天天都会听歌，但发现搞音乐类的太多了，我再搞个多没意思。考虑了一下，还是搞个看书的吧，这个还是很少有人搞的。找了找发现只有追书神器的api暴露出来了，起点之类的api找不到。最终效果因为api数据的限制，参考了起点中文网app、起点中文网web端，以及追书神器web端，再加上自己的一些想法搞出来的。项目中的小图标使用的是[阿里巴巴的矢量图标库Iconfont](http://www.iconfont.cn/)。\n\n## 技术栈\nvue2 + vuex + vue-router + webpack + ES6 + axios + sass\n\n##  目前只做了第一个版本，还存在一些需要进行优化的细节问题，后续会继续进行维护更新。如果发现什么问题，也欢迎跟我联系反馈。\n> - 如果觉得做的还行，对您有所帮助，欢迎“start”一下。\n\n\n## 开发中遇到的一些问题\n\n\n- ### 多个子组件循环，父组件如何处理加载状态（精选页面）\n\n```\n每个子组件加载完后是同`this.$emit`通知父组件，父组件判断所有子组件加载完成后隐藏loading。\n```\n\n- ### 跳转页面后active标记\n\n```\n最开始使用url.indexOf来处理，后来直接发现vue-router的exact属性更好用。\n```\n    \n- ### 调用第三方api接口时跨域的问题\n\n```\n1. 本地使用proxyTbale\n\n	在config/index.js中添加配置：\n	\n	\'/api\': {\n		target: \'http://api.zhuishushenqi.com\',\n		changeOrigin: true,\n		pathRewrite: {                \n		  \'^/api\': \'\'\n		}   \n	}\n	\n\n2. 部署服务器后通过nginx代理\n\n	nginx中配置：\n	\n	location /api/ {\n		 proxy_pass http://api.zhuishushenqi.com/;\n	 }\n\n3. 调用接口时只需要以`/api`开头就可以\n```\n        \n- ### 服务器部署后vue-router的虚拟路由在刷新时出现404页面\n\n```\n修改nginx配置：\n\nlocation / {\n	 try_files $uri $uri/ @router;          //增加的内容\n	 root /home/don/book;\n	 index index.html;\n}\n \nlocation @router {                          ', '2018-08-06 10:22:55', '2018-08-06 21:22:52', 'PUBLISHED', '前端', 'Vue');
INSERT INTO `article` VALUES (2, 'Go Web 开发之 Beego 框架初探', '又到周末啦。干完活之后，喷了一篇关于学习 Beego 的文章。这应该是 2016 年最后一篇技术文章啦，也是一气呵成，没有什么技术含量。\r\n\r\n选择 Go 语言\r\n断断续续看了 Go 几个星期了，讲真的真是喜欢的不得了。认真学过之后，你会觉得非常的优雅，写东西很舒服。学习 Go 我觉得很有必要的是，Go 中自带的数据结构很少，类似于 List 或者 Tree 之类的，最好尝试一下如何去设计一些常用的数据结构。话说回来，Go 的出身始终是一门后端语言。我非常后悔用 Flask 或者 Django 来作为我的后端入门框架或者选择。封装的太好了，往往对于一个入门新手来说学习不到什么。\r\n\r\n而 Go 就不一样了，它天生被设计是一门后端语言。也就是说，你将会学习到非常多的后端知识。看看下面这一张图，当时我看着就有一种很过瘾的感觉，因为这些知识你都知道，但是作为一个后端开发者你没有去了解过，这是非常大的失误。并不是想去用学习好 Go 去弥补没有学习好 C++ 的遗憾，只是新生事物，多尝试尝试总是极好的，哪怕只是浅尝辄止。Go 作为一门新的语言，其语言设计的特性，背后 Google 爸爸的撑腰以及现在 Docker 技术发展，前景应该还是不错的。所以如果你是编程新手或者是想入门后端的开发者，我强烈建议你选择 Go 语言。\r\n\r\n语言学到最后，框架总是少不了的。虽然不能完全依赖框架，但是还是可以学习一下框架的设计思想。对于 Beego 框架的评价总是各种各样，这也要看自己的选择了。之所以选择 Beego 框架来入门，主要是因为其详细的文档以及教程示例非常多。\r\n\r\n\r\n\r\nGo Web 初探\r\n先看一下最基本的 Go 中的 web 服务，只用到了 Go 中的 net/http 这个包：\r\n\r\npackage main\r\n\r\n    import (\r\n        \"fmt\"\r\n        \"net/http\"\r\n        \"strings\"\r\n        \"log\"\r\n    )\r\n\r\n    func sayhelloName(w http.ResponseWriter, r *http.Request) {\r\n        r.ParseForm()  //解析参数，默认是不会解析的\r\n        fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息\r\n        fmt.Println(\"path\", r.URL.Path)\r\n        fmt.Println(\"scheme\", r.URL.Scheme)\r\n        fmt.Println(r.Form[\"url_long\"])\r\n        for k, v := range r.Form {\r\n            fmt.Println(\"key:\", k)\r\n            fmt.Println(\"val:\", strings.Join(v, \"\"))\r\n        }\r\n        fmt.Fprintf(w, \"Hello astaxie!\") //这个写入到w的是输出到客户端的\r\n    }\r\n\r\n    func main() {\r\n        http.HandleFunc(\"/\", sayhelloName) //设置访问的路由\r\n        err := http.ListenAndServe(\":9090\", nil) //设置监听的端口\r\n        if err != nil {\r\n            log.Fatal(\"ListenAndServe: \", err)\r\n        }\r\n    }\r\n执行 go run web.go根据提示在浏览器地址栏打开 URL ，如下图所示：', '2018-08-08 17:35:49', '2018-08-08 17:35:53', 'PUBLISHED', '后端', 'Go');
INSERT INTO `article` VALUES (3, '微信小程序通讯录demo', '## 前言\r\n\r\n帮别人做的一个展示用的demo，最开始只需要一些假数据展示看看效果。不过到最后也没用上，所以就拿出来分享一下吧。我自己把后台接口部分给简单补齐了，做了一些假数据，样式也做了一些调整。因为接口需要https以及上线需要微信审核之类的，所以目前此demo只能在本地查看。同时功能也不是很完善，很多功能做的也比较简陋。如果是初学微信小程序的话，可以参考着看看。\r\n\r\n\r\n\r\n帮别人做的一个展示用的demo，最开始只需要一些假数据展示看看效果。不过到最后也没用上，所以就拿出来分享一下吧。我自己把后台接口部分给简单补齐了，做了一些假数据，样式也做了一些调整。因为接口需要https以及上线需要微信审核之类的，所以目前此demo只能在本地查看。同时功能也不是很完善，很多功能做的也比较简陋。如果是初学微信小程序的话，可以参考着看看。\r\n', '2018-06-06 18:28:39', '2018-08-06 21:22:58', 'PUBLISHED', '小程序', '微信');
INSERT INTO `article` VALUES (4, 'Test1', '考虑了一下，还是搞个看书的吧，这个还是很少有人搞的。找了找发现只有追书神器的api暴露出来了，起点之类的api找不到。最终效果因为api数据的限制，参考了起点中文网app、起点中文网web端，以及追书神器web端，再加上自己的一些想法搞出来的。项目中的小图标使用的是[阿里巴巴的矢量图标库Iconfont](http://www.iconfont.cn/)。\n\n## 技术栈\nvue2 + vuex + vue-router + webpack + ES6 + axios + sass\n\n##  目前只做了第一个版本，还存在一些需要进行优化的细节问题，后续会继续进行维护更新。如果发现什么问题，也欢迎跟我联系反馈。\n> - 如果觉得做的还行，对您有所帮助，欢迎“start”一下。\n\n\n## 开发中遇到的一些问题\n\n\n- ### 多个子组件循环，父组件如何处理加载状态（精选页面）\n\n```\n每个子组件加载完后是同`this.$emit`通知父组件，父组件判断所有子组件加载完成后隐藏loading。\n```\n\n- ### 跳转页面后active标记\n\n```\n最开始使用url.indexOf来处理，后来直接发现vue-router的', '2018-07-06 02:22:55', '2018-08-06 21:23:01', 'PUBLISHED', '前端', 'Tag1');
INSERT INTO `article` VALUES (5, '基于Vue2开发的WebAPP', '\n## 前言\n\n初学vue.js，官网的文档写的很清楚，看着不难。俗话说：光说不练假把式。程序猿学个新东西还是要敲出来看看效果比较好。最开始是想搞个音乐类的，毕竟天天都会听歌，但发现搞音乐类的太多了，我再搞个多没意思。考虑了一下，还是搞个看书的吧，这个还是很少有人搞的。找了找发现只有追书神器的api暴露出来了，起点之类的api找不到。最终效果因为api数据的限制，参考了起点中文网app、起点中文网web端，以及追书神器web端，再加上自己的一些想法搞出来的。项目中的小图标使用的是[阿里巴巴的矢量图标库Iconfont](http://www.iconfont.cn/)。\n\n## 技术栈\nvue2 + vuex + vue-router + webpack + ES6 + axios + sass\n\n##  目前只做了第一个版本，还存在一些需要进行优化的细节问题，后续会继续进行维护更新。如果发现什么问题，也欢迎跟我联系反馈。\n> - 如果觉得做的还行，对您有所帮助，欢迎“start”一下。\n\n\n## 开发中遇到的一些问题\n\n\n- ### 多个子组件循环，父组件如何处理加载状态（精选页面）\n\n```\n每个子组件加载完后是同`this.$emit`通知父组件，父组件判断所有子组件加载完成后隐藏loading。\n```\n\n- ### 跳转页面后active标记\n\n```\n最开始使用url.indexOf来处理，后来直接发现vue-router的exact属性更好用。\n```\n    \n- ### 调用第三方api接口时跨域的问题\n\n```\n1. 本地使用proxyTbale\n\n	在config/index.js中添加配置：\n	\n	\'/api\': {\n		target: \'http://api.zhuishushenqi.com\',\n		changeOrigin: true,\n		pathRewrite: {                \n		  \'^/api\': \'\'\n		}   \n	}\n	\n\n2. 部署服务器后通过nginx代理\n\n	nginx中配置：\n	\n	location /api/ {\n		 proxy_pass http://api.zhuishushenqi.com/;\n	 }\n\n3. 调用接口时只需要以`/api`开头就可以\n```\n        \n- ### 服务器部署后vue-router的虚拟路由在刷新时出现404页面\n\n```\n修改nginx配置：\n\nlocation / {\n	 try_files $uri $uri/ @router;          //增加的内容\n	 root /home/don/book;\n	 index index.html;\n}\n \nlocation @router {                          ', '2018-06-06 10:22:55', '2018-08-06 21:23:03', 'PUBLISHED', '前端', 'Vue');

-- ----------------------------
-- Table structure for article_tag
-- ----------------------------
DROP TABLE IF EXISTS `article_tag`;
CREATE TABLE `article_tag`  (
  `id` int(11) NOT NULL,
  `article_id` int(11) NOT NULL COMMENT '文章id',
  `tag_id` int(11) NOT NULL COMMENT '标签id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article_tag
-- ----------------------------
INSERT INTO `article_tag` VALUES (1, 1, 1);
INSERT INTO `article_tag` VALUES (2, 2, 2);

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` int(11) NOT NULL,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文章分类名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (1, '前端');
INSERT INTO `category` VALUES (2, '开发');
INSERT INTO `category` VALUES (3, '小程序');

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag`  (
  `id` int(11) NOT NULL,
  `tag` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标签名',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tag
-- ----------------------------
INSERT INTO `tag` VALUES (1, 'Vue');
INSERT INTO `tag` VALUES (2, 'JS');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `password` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'admin', '123456');
INSERT INTO `user` VALUES (3, 'ad', '123');
INSERT INTO `user` VALUES (17, 'string', 'string');
INSERT INTO `user` VALUES (18, '', '1234');
INSERT INTO `user` VALUES (19, '123', '1234');
INSERT INTO `user` VALUES (20, 'ok123', '1234');
INSERT INTO `user` VALUES (21, 'oaak123', '1234');
INSERT INTO `user` VALUES (22, 'root', '123');

SET FOREIGN_KEY_CHECKS = 1;
