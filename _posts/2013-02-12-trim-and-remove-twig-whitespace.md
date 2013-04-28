---
layout: post
title: Trim and remove Twig whitespace
category: symfony
tags: [symfony, twig]
summary: 
---
We seldom see the HTML produced by our backends, at least in their pristine state, still unprettified by our web inspectors. Yet I am not alone in taking pride in clean markup.

The template engine I use a lot nowadays is [Twig](/twig/) and it has two ways of controlling whitespace, to make our HTML the envy of our peers. One is *the `spaceless` tag* and the other is *the whitespace control modifier*.

## Spaceless tag

The spaceless tag removes all whitespace between its child tags, even Twig tags. I use it myself in the top level `base.html.twig` template to compress the entire HTML document some.

    {% raw  %}
    {% spaceless %}
        <html>
            {% if name %}
                <body>   Hello {{ name }}!   </body>
            {% endif %}
        </html>
    {% endspaceless %}
    {% endraw %}

… will produce …

    <html><body>Hello Tobias!</body></html>

[Its documentation](http://twig.sensiolabs.org/doc/tags/spaceless.html) says not to use this for size optimization but I am just a rebel like that.

## Whitespace control modifier

While the spaceless tag trims whitespace inside of it, the whitespace control modifier instead trims outside of itself. Just add a `-` to the delimiter you want to trim from.

    <div>
        {% raw %}{{- name }}{% endraw %}
    </div>

… will produce …

    <div>Tobias
    </div>

Use `{% raw %}{{- name -}}{% endraw%}` to also trim from the right side. This works for `{% raw %}{{ … }}{% endraw %}` prints, `{% raw %}{% … %}{% endraw %}` statements and `{# … #}` comments.

As per [the documentation](http://twig.sensiolabs.org/doc/templates.html#whitespace-control), this was introduced in Twig 1.1. So you might need to upgrade if you are running older code.
