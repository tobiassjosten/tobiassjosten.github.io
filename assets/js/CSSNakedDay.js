/*jslint browser:true*/

/*
 * This file is part of the CSSNakedDay.js project.
 *
 * (c) Tobias Sjösten <tobias.sjosten@gmail.com>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

(function () {
    'use strict';

    if (!document.styleSheets) {
        return;
    }

    var stripInline, stripExternal,
        today = new Date(),
        oldOnload = window.onload;

    if (today.getDate() !== 9 || today.getMonth() !== 3) {
        return;
    }

    stripExternal = function () {
        var i, j = document.styleSheets.length;
        for (i = 0; i < j; i += 1) {
            document.styleSheets[i].disabled = true;
        }
    };

    stripInline = document.querySelectorAll
        ? function () {
            var i, elements = document.querySelectorAll('*[style]'),
                j = elements.length;

            for (i = 0; i < j; i += 1) {
                elements[i].style.cssText = '';
            }
        }
        : function (element) {
            var i, j = element.childNodes.length;

            for (i = 0; i < j; i += 1) {
                stripInline(element.childNodes[i]);
            }

            if (element.style) {
                element.style.cssText = '';
            }
        };

    stripExternal();

    window.onload = function () {
        if (oldOnload) { oldOnload(); }

        stripExternal();
        stripInline(document);
    };
}());
