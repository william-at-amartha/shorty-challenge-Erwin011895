package module

import (
	"testing"

	"github.com/koding/cache"
	"github.com/golang/mock/gomock"

	pkgdto "github.com/Erwin011895/shorty-challenge/pkg/dto"
	"github.com/Erwin011895/shorty-challenge/internal/util/testutil"
	"github.com/Erwin011895/shorty-challenge/internal/component"
	"github.com/Erwin011895/shorty-challenge/internal/model"
	"github.com/Erwin011895/shorty-challenge/pkg/constant"
	"github.com/Erwin011895/shorty-challenge/internal/mocks"
)

func TestShortenURL(t *testing.T) {
	// setups
	mc := mocks.InitMockComponent(t)
	sc := component.InitMockSharedComponent(mc)

	module := NewShortURLModule(&ShortURLModuleParams{
		Cache: sc.Cache,
	})
    ctx := testutil.CreateContext()

    // sub tests
    t.Run("new shortcode", func(t *testing.T){
    	wantedShortcode := "exampl"

		mc.KodingCache.EXPECT().Get(wantedShortcode).Return(nil, cache.ErrNotFound)
		mc.KodingCache.EXPECT().Set(wantedShortcode, gomock.Any()).Return(nil)

		shortcode, err := module.ShortenURL(ctx, &pkgdto.BodyPostShortenURL{
			URL: "https://example.com",
			Shortcode: wantedShortcode,
		})

		if err != nil {
			t.Fatalf(`got %q, %v, want match for %q, %v`, shortcode, err, wantedShortcode, nil)
		}
    })

	t.Run("no shortcode on body", func(t *testing.T){
		mc.KodingCache.EXPECT().Get(gomock.Any()).Return(nil, cache.ErrNotFound)
		mc.KodingCache.EXPECT().Set(gomock.Any(), gomock.Any()).Return(nil)

		shortcode, err := module.ShortenURL(ctx, &pkgdto.BodyPostShortenURL{
			URL: "https://example.com",
		})

		if err != nil {
			t.Fatalf(`got %q, %v, want match for %q, %v`, shortcode, err, gomock.Any(), nil)
		}
    })

    t.Run("empty URL", func(t *testing.T){
    	wantedShortcode := "exampl"

		shortcode, err := module.ShortenURL(ctx, &pkgdto.BodyPostShortenURL{
			Shortcode: wantedShortcode,
		})

		if err != constant.ErrMissingBodyURL {
			t.Fatalf(`got %q, %v, want match for %q, %v`, shortcode, err, gomock.Any(), constant.ErrMissingBodyURL)
		}
    })

    t.Run("invalid shortcode", func(t *testing.T){
    	wantedShortcode := "example" // 7 letter is invalid, must be 6

		shortcode, err := module.ShortenURL(ctx, &pkgdto.BodyPostShortenURL{
			URL: "https://example.com",
			Shortcode: wantedShortcode,
		})

		if err != constant.ErrShortcodeNotAlphaNumeric {
			t.Fatalf(`got %q, %v, want match for %q, %v`, shortcode, err, gomock.Any(), constant.ErrShortcodeNotAlphaNumeric)
		}
    })

    t.Run("shortcode already in use", func(t *testing.T){
    	wantedShortcode := "exampl"

		mc.KodingCache.EXPECT().Get(wantedShortcode).Return(nil, nil)

		shortcode, err := module.ShortenURL(ctx, &pkgdto.BodyPostShortenURL{
			URL: "https://example.com",
			Shortcode: wantedShortcode,
		})

		if err != constant.ErrShortcodeAlreadyInUse {
			t.Fatalf(`got %q, %v, want match for %q, %v`, shortcode, err, gomock.Any(), constant.ErrShortcodeAlreadyInUse)
		}
    })
}

func TestGetURL(t *testing.T) {
	// setups
	mc := mocks.InitMockComponent(t)
	sc := component.InitMockSharedComponent(mc)

	module := NewShortURLModule(&ShortURLModuleParams{
		Cache: sc.Cache,
	})
    ctx := testutil.CreateContext()

    // sub tests
    t.Run("success get URL", func(t *testing.T){
    	wantedShortcode := "exampl"
    	wantedURL := "https://example.com"

		mc.KodingCache.EXPECT().Get(wantedShortcode).Return(model.ShortURL{
			RedirectCount: 0,
			URL: wantedURL,
		}, nil)
		mc.KodingCache.EXPECT().Set(wantedShortcode, gomock.Any()).Return(nil)

		url, err := module.GetURL(ctx, wantedShortcode)

		if err != nil {
			t.Fatalf(`got %q, %v, want match for %q, %v`, url, err, wantedURL, nil)
		}
    })

    t.Run("shortcode not found", func(t *testing.T){
    	wantedShortcode := "exampl"

		mc.KodingCache.EXPECT().Get(wantedShortcode).Return(nil, cache.ErrNotFound)

		url, err := module.GetURL(ctx, wantedShortcode)

		if err != constant.ErrShortcodeNotFound {
			t.Fatalf(`got %q, %v, want match for %q, %v`, url, err, "", constant.ErrShortcodeNotFound)
		}
    })
}

func TestGetStats(t *testing.T) {
	// setups
	mc := mocks.InitMockComponent(t)
	sc := component.InitMockSharedComponent(mc)

	module := NewShortURLModule(&ShortURLModuleParams{
		Cache: sc.Cache,
	})
    ctx := testutil.CreateContext()

    // sub tests
    t.Run("success get stats", func(t *testing.T){
    	wantedShortcode := "exampl"
    	wantedURL := "https://example.com"

		mc.KodingCache.EXPECT().Get(wantedShortcode).Return(model.ShortURL{
			RedirectCount: 0,
			URL: wantedURL,
		}, nil)

		stats, err := module.GetStats(ctx, wantedShortcode)

		if err != nil {
			t.Fatalf(`got %v, %v, want match for %v, %v`, stats, err, stats, nil)
		}
    })

    t.Run("shortcode not found", func(t *testing.T){
    	wantedShortcode := "exampl"

		mc.KodingCache.EXPECT().Get(wantedShortcode).Return(nil, cache.ErrNotFound)

		stats, err := module.GetStats(ctx, wantedShortcode)

		if err != constant.ErrShortcodeNotFound {
			t.Fatalf(`got %v, %v, want match for %v, %v`, stats, err, "", constant.ErrShortcodeNotFound)
		}
    })
}
