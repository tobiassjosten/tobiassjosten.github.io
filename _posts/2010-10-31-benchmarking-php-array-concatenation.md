---
layout: post
title: Benchmarking PHP array concatenation
category: php
tags: [php, benchmark]
summary: I recently began working on a contribution to a small programming contest, where the challenge is to optimize ones code to run as fast as possible. One of the parts in my code that caught my interest was some implode() calls. My gut feeling told me that it might be slower than a straight string concatenation.
---
I recently began working on a contribution to a small [programming contest](http://www.phpportalen.net/viewtopic.php?t=113904) (swedish), where the challenge is to optimize ones code to run as fast as possible. One of the parts in my code that caught my interest was some `implode()` calls. My gut feeling told me that it might be slower than a straight string concatenation.

The code I was looking at was basically something like this:

    implode('|', array(0, 1));

Of course to find out for sure if A is faster than B, you need to benchmark. So I wrote some code to benchmark `implode()` against string concatenation. Then I ran the tests a couple of times and picked their fastest execution. Here are my results.

    // Benchmarking implode().
    $x = array(0, 1);
    for ($i = 0; $i < 1000000; $i++)
    {
      $y = implode('|', $x);
    }

*user  0m2.670s*

    // Benchmarking string concatenation.
    $x = array(0, 1);
    for ($i = 0; $i < 1000000; $i++)
    {
      $y = $x[0].'|'.$x[1];
    }

*user  0m0.660s*

It turns out I was right. String concatenation is over four times as fast! I also ran two more tests to see how other concatenation methods performed.

    // Benchmarking variable interpolation.
    $x = array(0, 1);
    for ($i = 0; $i < 1000000; $i++)
    {
      $y = "{$x[0]}|{$x[1]}";
    }

*user  0m0.680s*

    // Benchmarking sprintf().
    $x = array(0, 1);
    for ($i = 0; $i < 1000000; $i++)
    {
      $y = sprintf('%d|%d',$x[0],$x[1]);
    }

*user  0m2.730s*

All in all it seems working with strings is much, much faster than the convenience functions `implode()` and `sprintf()`.

I was happy having learned this, but then I figured maybe there is some merit to using `implode()` anyway. Speed wise, of course, since it is convenient to use `implode()` over string concatenation at times, rather than having to implement the logic in `implode()` all over.

I decided to benchmark bigger arrays and filled them up with fifty items before running the same tests again.

    // Benchmarking implode() with a bigger array.
    $x = array(
      0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
      0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
      0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
      0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
      0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
    );
    for ($i = 0; $i < 1000000; $i++)
    {
      $y = implode('|', $x);
    }

*user  0m6.060s*

Here `implode()` performed pretty nicely, with only a 227% increase in time for a 2500% bigger array!

    // Benchmarking string concatenation with a bigger array.
    $x = array(
      0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
      0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
      0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
      0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
      0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
    );
    for ($i = 0; $i < 1000000; $i++)
    {
      $y = $x[0].'|'.$x[1].'|'.$x[2].'|'.$x[3].'|'.$x[4]
      .'|'.$x[5].'|'.$x[6].'|'.$x[7].'|'.$x[8].'|'.$x[9]
      .'|'.$x[10].'|'.$x[11].'|'.$x[12].'|'.$x[13].'|'.$x[14]
      .'|'.$x[15].'|'.$x[16].'|'.$x[17].'|'.$x[18].'|'.$x[19]
      .'|'.$x[20].'|'.$x[21].'|'.$x[22].'|'.$x[23].'|'.$x[24]
      .'|'.$x[25].'|'.$x[26].'|'.$x[27].'|'.$x[28].'|'.$x[29]
      .'|'.$x[30].'|'.$x[31].'|'.$x[32].'|'.$x[33].'|'.$x[34]
      .'|'.$x[35].'|'.$x[36].'|'.$x[37].'|'.$x[38].'|'.$x[39]
      .'|'.$x[40].'|'.$x[41].'|'.$x[42].'|'.$x[43].'|'.$x[44]
      .'|'.$x[45].'|'.$x[46].'|'.$x[47].'|'.$x[48].'|'.$x[49]
      ;
    }

*0m17.490s*

Now this was interesting. String concatenation was 2649% slower with a 2500% bigger array. It scales more or less 1:1 with the size of the array.

This must mean there is a breaking point at which time it does not matter, for speed's sake, which method you use. Further benchmarks showed me this lays somewhere around having ten items in the array, at which point both methods took 3.2-3.3 seconds to complete.

I finish with a tl;dr summary:

    $y = count($x) >= 10 ? implode('|', $x) : $x[0].'|'.$x[1]/*etc*/;
