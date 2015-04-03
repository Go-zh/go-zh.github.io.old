---
layout: default
title: 博客翻译状态
summary: 记录了Go官方博客的翻译状态及任务分配。
---

本文档用于记录[Go官方博客](http://blog.golang.org/)的翻译状态及任务分配。

**注意**：曾经有一部分文章在Go的主源码库中，后来官方将其转换为`article`格式并迁移到了博客中。
我们翻译过的文章还未完成转换，旧的文件在本站源码库的`zh-articles`目录内。
新旧文件间的对应关系已在下方的列表中用`=对应=`标记标出，转换方式及译法请阅读[翻译规范](.trans-spec.html)。

博客内容位于`Go-zh/blog`子代码库的`content`目录中。
关于`article`格式的详细信息，参见`Go-zh/tools/{blog, present}`内的源码和文档。

鉴于官方博客更新频繁且文件名死长，故单独分离出列表显示状态。格式如下：

<pre>
译者: 状态[(注释)]|译文标题: 文件名

其中：
译者 = 译者的ID[ & 同译者的ID]
状态 = *待译*|*翻译*|*校对*|*整理*|完成|…
注释 = 注释者ID: 具体说明

|  表示或者
[] 表示可选
</pre>

例如：

    OlingCat & minux: 完成 (OlingCat: *整理*)|竞态检测器: race-detector.article

## 以下为文章列表
<pre>
4years.article
5years.article
a-conversation-with-the-go-team.article
advanced-go-concurrency-patterns.article
appengine-dec2013.article
building-stathat-with-go.article
c-go-cgo.article =对应= (c_go_cgo.html)
concurrency-is-not-parallelism.article
constants.article
context.article
cover.article
debugging-go-code-status-report.article
debugging-go-programs-with-gnu-debugger.article
defer-panic-and-recover.article =对应= (defer_panic_recover.html)
docker.article
error-handling-and-go.article =对应= (error_handling.html)
first-class-functions-in-go-and-new-go.article
first-go-program.article
fosdem14.article
from-zero-to-go-launching-on-google.article
gccgo-in-gcc-471.article
generate.article
getthee-to-go-meetup.article
getting-to-know-go-community.article
gif-decoder-exercise-in-go-interfaces.article
go-11-is-released.article
go-and-google-app-engine.article
go-and-google-cloud-platform.article
go-app-engine-sdk-155-released.article
go-at-google-io-2011-videos.article
go-at-heroku.article
go-at-io-frequently-asked-questions.article
go-becomes-more-stable.article
go-concurrency-patterns-timing-out-and.article =对应= (concurrency_patterns.html)
go-fmt-your-code.article
go-for-app-engine-is-now-generally.article
go-image-package.article =对应= (image_package.html)
go-imagedraw-package.article =对应= (image_draw.html)
go-maps-in-action.article
go-one-year-ago-today.article
go-programming-language-turns-two.article
go-programming-session-video-from.article
go-slices-usage-and-internals.article =对应= (slices_usage_and_internals.html)
go-turns-three.article
go-updates-in-app-engine-171.article
go-version-1-is-released.article
go-videos-from-google-io-2012.article
go-whats-new-in-march-2010.article
go-wins-2010-bossie-award.article
go1.3.article
go1.4.article
go12.article
gobs-of-data.article =对应= (gobs_of_data.html)
godoc-documenting-go-code.article =对应= (godoc_documenting_go_code.html)
gopher.article
gophercon.article
gophergala.article
gos-declaration-syntax.article =对应= (gos_declaration_syntax.html)
gothamgo.article
introducing-go-playground.article
introducing-gofix.article
io2014.article
json-and-go.article =对应= (json_and_go.html)
json-rpc-tale-of-interfaces.article =对应= (json_rpc_tale_of_interfaces.html)
laws-of-reflection.article =对应= (laws_of_reflection.html)
learn-go-from-your-browser.article
new-talk-and-tutorials.article
normalization.article
organizing-go-code.article
oscon.article
osconreport.article
package-names.article
pipelines.article
playground.article
preview-of-go-version-1.article
profiling-go-programs.article
race-detector.article =对应= (race_detector.html)
real-go-projects-smarttwitter-and-webgo.article
share-memory-by-communicating.article
slices.article
spotlight-on-external-go-libraries.article
strings.article
the-app-engine-sdk-and-workspaces-gopath.article
the-path-to-go-1.article
third-party-libraries-goprotobuf-and.article
two-go-talks-lexical-scanning-in-go-and.article
two-recent-go-articles.article
two-recent-go-talks.article
upcoming-google-io-go-events.article
writing-scalable-app-engine.article
</pre>
