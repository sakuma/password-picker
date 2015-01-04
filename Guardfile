guard 'go', :server => 'main.go', :args => ['a'] do
  watch(%r{\.go$})
  watch(%r{\.tpl$})
end

guard :shell do
  watch(%r{\.(css|js)}) { `grunt` }
end
