---
layout: base
title: Über blog
summary: A tech blog where I usually write about problems I run into and solutions I stumble over.
---
<ul class="tags">
    {% for category in site.categories | sort %}
        <li>
            <a class="tag" href="/{{ category | first | slugize }}/">
                {{ category | first }}
            </a>
        </li>
    {% endfor %}
</ul>

<ul class="posts">
    {% assign year = "" %}

    {% for post in site.posts %}
        {% capture newyear %}{{ post.date | date: "%Y" }}{% endcapture%}
        {% if newyear != year %}
            {% if year != "" %}
                </ul></li>
            {% endif %}

            <li><h2>{{ newyear }}</h2><ul>

            {% assign year = newyear %}
        {% endif %}

        <li>
            <time>{{ post.date | date: "%Y-%m-%d" }}</time>
            <a href="{{ post.url }}">{{ post.title }}</a>
        </li>
    {% endfor %}
    </ul></li>
</ul>
