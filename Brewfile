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
  brew 'go'
  brew 'goreleaser'
end
