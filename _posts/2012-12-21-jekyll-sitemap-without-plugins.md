---
layout: post
title: Jekyll sitemap without plugins
category: jekyll
tags: [jekyll]
summary: There are plenty of Jekyll plugins available if you want to generate a sitemap for your blog. The only problem is that GitHub Pages, where a lot of us are hosting our Jekyll sites, does not allow plugins.
---
There are plenty of [Jekyll](/jekyll/) plugins available if you want to generate [a sitemap](http://www.sitemaps.org/) for your blog. The only problem is that [GitHub Pages](http://pages.github.com/), where a lot of us are hosting our Jekyll sites, does not allow plugins.

Thankfully, Jekyll is quite versatile and I was able to *create a sitemap without using Jekyll plugins*. In the undocumented `:pages` accessor you can find all non-post pages in your site and iterating over it gives you all the data you need to build a sitemap.

Just remember to exclude things like CSS files (detectable by missing layout) and your feeds (layout is `feed` in my case).

    ---
    ---
    <?xml version="1.0" encoding="UTF-8"?>
    <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
        {{ "{% for post in site.posts " }}%}
        <url>
            <loc>http://foo.com{{ "{{ post.url | remove: 'index.html' " }}}}</loc>
        </url>
        {{ "{% endfor " }}%}

        {{ "{% for page in site.pages " }}%}
        {{ "{% if page.layout != nil "}}%}
        {{ "{% if page.layout != 'feed' "}}%}
        <url>
            <loc>http://foo.com{{ "{{ page.url | remove: 'index.html' " }}}}</loc>
        </url>
        {{ "{% endif " }}%}
        {{ "{% endif " }}%}
        {{ "{% endfor " }}%}
    </urlset>

Obviously, for your own use you will need to update it with your own domain name. Save it as `sitemap.xml` and you are good to go!

**Update:** As ArmNo pointed out, you will need to include `site.posts` *and* `site.page` in order to cover all the pages of your website.
