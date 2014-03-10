/*jslint browser:true, nomen:true*/

var _gaq = _gaq || [];

/**
 * Track link clicks.
 */
(function () {
    'use strict';

    function trackLink(link) {
        var category = link.getAttribute('data-clicktrack-category'),
            label = link.getAttribute('data-clicktrack-label');

        if (!category || !label) {
            return;
        }

        link.setAttribute('target', '_blank');
        link.addEventListener('click', function () {
            _gaq.push(['_trackEvent', category, 'Click', label]);
            mixpanel.track('Click ' + category + ' ' + label);
        });
    }

    var links = document.getElementsByTagName('a'),
        i = 0,
        n = links.length;

    for (i = 0; i < n; i = i + 1) {
        trackLink(links[i]);
    }
}());
