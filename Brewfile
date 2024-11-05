#!/usr/bin/env brew
# @see https://github.com/Homebrew/homebrew-bundle

###
# Application dev dependencies.
#
# Install with:
#   `brew bundle install`
##

# Only install these dependencies on non-ci environments
if ENV.fetch('CI', 'false') != 'true'
  brew 'direnv'
  brew 'pre-commit'
end

# go dependencies
# brew 'gcc'
if ENV.fetch('CI', 'false') != 'true'
  brew 'go@1.23'
  brew 'goreleaser'
end

puts 'Installing go tools'
%x( go version )
File.readlines('go.tools').each do |line|
  %x( go install #{line} )
end
