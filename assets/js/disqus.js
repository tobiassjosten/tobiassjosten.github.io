function disqus_config() {
    var _gaq = _gaq || [];
    this.callbacks.onNewComment = [ function() {
        _gaq.push([ '_trackEvent', 'Comments', 'New comment' ]);
    } ];
}
