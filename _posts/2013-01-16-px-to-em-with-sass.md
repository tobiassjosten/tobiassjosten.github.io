---
layout: post
title: PX to EM with Sass
category: css
tags: [css]
summary: I used to manually calculate px-to-em but with Sass and Compass, this is a thing of the past.
---
Whenever I used to write [CSS](/css/) I whipped out *a px-em calculator* to help my pixel-thinking brain write proper-em-values. After having been introduced to *Sass and Compass*, this is thankfully a thing of the past.

Sass lets you define variables to hold the values you use a lot.

    $base-font-size: 20px;

    html {
        font-size: $base-font-size;
    }

Sass also lets you define functions, which can take parameters and even default parameters to set values.

    @function em($px, $base: $base-font-size) {
        @return ($px / $base) * 1em;
    }

All you have to do is give this function whatever pixel value you would normally have converted manually.

    .title {
        font-size: em(37px);
    }

Since `em` values are relative its tag's parent, you need to keep changes to the current size when nesting.

    .fine-print {
        $font-size: 10px;
        font-size: em($font-size);
        .title {
            font-size: em(10px, $font-size);
        }
    }

The Sass documentation has some [more information on functions](http://sass-lang.com/docs/yardoc/Sass/Script/Functions.html) if you feel like diving in.
