---
layout: default
title: Go 语言中文翻译项目
---
# {{ page.title }}

## 文档列表
<ul>
{% for docs in site.docs %}
<li><a href="{{ site.baseurl }}{{ docs.url }}">{{ docs.title }}</a></li>
{% endfor %}
</ul>
