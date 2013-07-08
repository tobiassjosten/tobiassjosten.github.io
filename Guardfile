guard 'jekyll' do
  watch %r{^(?!_site/|_includes/sass)}
end

guard 'sass', :input => '_includes/sass', :output => '_includes/css'

guard 'livereload' do
  watch '_site/index.html'
end
