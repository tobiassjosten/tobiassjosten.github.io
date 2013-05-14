---
layout: post
title: Jekyll tag cloud
category: jekyll
tags: [jekyll]
summary: Tag clouds was all the fuss some years ago and they can still be useful for content heavy sites. But how do you create a tag cloud in Jekyll?
---
Tag clouds was all the fuss some years ago and while the fad has now run its course, they can still be useful for content heavy sites. But how do you create a *tag cloud in Jekyll*?

There are of course plugins to build tag clouds but as I host my [Jekyll sites at GitHub Pages](/jekyll/reproduce-jekyll-on-github-pages/), I want to stay clear from plugins. For example, my [Jekyll sitemap is plugin free](/jekyll/jekyll-sitemap-without-plugins/) and you can do the same with tag clouds.

## Listing tags and categories

Tags and categories both exists within the `site` variable, under `site.tags` and `site.categories` respectively. Tag clouds obviously denotes using tags but the way I use the two myself, I go with categories.

    {{ "{% for category in site.categories " }}%}
        {{ "{{ category | first " }}}}
    {{ "{% endfor " }}%}

This will print all your categories. The `first` filter takes care of extracting the string representation of the category, as it is actually an array. But you also want to link them somewhere, right?

Personally I manually create landing pages for all my tags and categories, so that I can fill them up with some more descriptive content, helpful links and such. If my tag is `jekyll` then I create a page at `jekyll/index.md`, which will be then be available at [/jekyll/](/jekyll/).

With this setup I can easily link my tags and categories to their respective landing pages.

    {{ "{% for category in site.categories " }}%}
        <a href="/{{ "{{ category | first | slugize " }}}}/">
            {{ "{{ category | first " }}}}
        </a>
    {{ "{% endfor " }}%}

## Tag clouds in Jekyll

Now here comes the tricky part; a tag cloud weighs its tags by the number of uses. [My own blog](/blog/), for example, would probably have [PHP](/php/) and [Symfony](/symfony/) in really big letters because I wrote a lot on those topics.

The second (`last`) part of the category or tag contains its posts, so we can use that to count its usage. Divide that by the total number of posts and you will end up with a measure of relative usage. Because Liquid, Jekyll's rendering engine, does not support math expressions we will have to hack this by chaining filters together.

    {{ "{% for category in site.categories " }}%}
        <li style="font-size: {{ "{{ category | last | size | times: 100 | divided_by: site.categories.size " }}}}%">
            <a href="/{{ "{{ category | first | slugize " }}}}/">
                {{ "{{ category | first " }}}}
            </a>
        </li>
    {{ "{% endfor " }}%}

This snippet will get you a list of your categories, with their relative usage in percentage as their font size. Hacky, verbose but it works!

Of course, seldom used tags and categories could easily become nearly invisible at just a few percentages of your font size. One solution is to add an arbitrary number to it so it looks good.

    {{ "{{ category | last | size | times: 100 | divided_by: site.categories.size | plus: 70 " }}}}

And there you have it; *a tag cloud in Jekyll* with no plugins!
