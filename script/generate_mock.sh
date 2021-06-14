# ~/go/bin/mockgen -destination ./internal/mocks/mockcache/kodingcache.go -package mockcache github.com/koding/cache Cache

#bin/sh

# cache
mockgen -destination ./internal/mocks/mockcache/kodingcache.go -package mockcache github.com/koding/cache Cache

# module
mockgen -source ./internal/module/shorturl.go -destination internal/mocks/mockmodule/shorturl_mock.go -package mockmodule
