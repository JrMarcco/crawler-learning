package collect

import (
	"bufio"
	"context"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"log"
	"net/http"
)

type Fetcher interface {
	Get(ctx context.Context, url string) ([]byte, error)
}

type BaseFetcher struct {
}

func (bf BaseFetcher) Get(ctx context.Context, url string) ([]byte, error) {

	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("get url failed: %v", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")

	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(rsp.Body)

	if http.StatusOK != rsp.StatusCode {
		return nil, fmt.Errorf("error status code %d", rsp.StatusCode)
	}

	br := bufio.NewReader(rsp.Body)
	e := determineEncoding(br)

	u8Reader := transform.NewReader(br, e.NewEncoder())

	return io.ReadAll(u8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {

	bytes, err := r.Peek(1024)
	if err != nil {
		log.Print(err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
