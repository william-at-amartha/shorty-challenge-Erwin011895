package module

import (
	"context"
	"time"

	"github.com/koding/cache"

	pkgdto "github.com/Erwin011895/shorty-challenge/pkg/dto"
	"github.com/Erwin011895/shorty-challenge/pkg/constant"
	"github.com/Erwin011895/shorty-challenge/internal/model"
	"github.com/Erwin011895/shorty-challenge/internal/component"
	"github.com/Erwin011895/shorty-challenge/internal/util/shortcodeutil"
)

type ShortURLModuleWrapper interface {
	ShortenURL(ctx context.Context, body *pkgdto.BodyPostShortenURL) (shortcode string, err error)
	GetURL(ctx context.Context, shortcode string) (url string, err error)
	GetStats(ctx context.Context, shortcode string) (pkgdto.ResponseGetStats, error)
}

type ShortURLModuleParams struct {
	Cache *component.Cache
}

type ShortURLModule struct {
	Cache *component.Cache
}

func NewShortURLModule(params *ShortURLModuleParams) *ShortURLModule {
	return &ShortURLModule{
		Cache: params.Cache,
	}
}

func (m *ShortURLModule) ShortenURL(ctx context.Context, body *pkgdto.BodyPostShortenURL) (shortcode string, err error) {
	if len(body.URL) == 0 {
		err = constant.ErrMissingBodyURL
		return
	}

	shortcode = body.Shortcode
	
	if shortcode == "" {
	    shortcode = shortcodeutil.GenerateShortcode(6)
	}

	isValid := shortcodeutil.ValidateShortcode(shortcode)
	if !isValid {
		err = constant.ErrShortcodeNotAlphaNumeric
		return
	}

	_, err = m.Cache.KodingCache.Get(shortcode)
	if err != cache.ErrNotFound {
		err = constant.ErrShortcodeAlreadyInUse
		return
	}

	now := time.Now()
	err = m.Cache.KodingCache.Set(shortcode, model.ShortURL{
		URL: body.URL,
		StartDate: now,
		LastSeenDate: now,
		RedirectCount: 0,
	})

	return
}

func (m *ShortURLModule) GetURL(ctx context.Context, shortcode string) (url string, err error) {
	shortURLi, err := m.Cache.KodingCache.Get(shortcode)
	if err != nil {
		err = constant.ErrShortcodeNotFound
		return
	}

	shortURL := shortURLi.(model.ShortURL)
	shortURL.LastSeenDate = time.Now()
	shortURL.RedirectCount++
	err = m.Cache.KodingCache.Set(shortcode, shortURL)

	url = shortURL.URL
	return
}

func (m *ShortURLModule) GetStats(ctx context.Context, shortcode string) (pkgdto.ResponseGetStats, error) {
	resp := pkgdto.ResponseGetStats{}

	shortURLi, err := m.Cache.KodingCache.Get(shortcode)
	if err != nil {
		err = constant.ErrShortcodeNotFound
		return resp, err
	}

	shortURL := shortURLi.(model.ShortURL)

	resp = pkgdto.ResponseGetStats{
		StartDate: shortURL.StartDate,
		LastSeenDate: shortURL.LastSeenDate,
		RedirectCount: shortURL.RedirectCount,
	}
	return resp, err
}
