package frame

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
	"io"

	"github.com/rs/zerolog"
	"github.com/vmihailenco/msgpack/v5"
)

const (
	PageHashAttr     = "data-hlive-hash"
	PageHashAttrTmpl = "{data-hlive-hash}"
)

const (
	msgpackExtHTML int8 = iota
	msgpackExtTag
	msgpackExtAttr
	msgpackExtNodeGroup
)

func init() {
	msgpack.RegisterExt(msgpackExtHTML, (*engine.HTML)(nil))
	msgpack.RegisterExt(msgpackExtTag, (*engine.Tag)(nil))
	msgpack.RegisterExt(msgpackExtAttr, (*engine.Attribute)(nil))
	msgpack.RegisterExt(msgpackExtNodeGroup, (*engine.NodeGroup)(nil))
}

// Cache allow cache adapters to be used in HLive
type Cache interface {
	Get(key any) (value any, hit bool)
	Set(key any, value any)
}

// PipelineProcessorRenderHashAndCache that will cache the returned tree to support SSR
func PipelineProcessorRenderHashAndCache(logger zerolog.Logger, renderer *engine.Renderer, cache Cache) *engine.PipelineProcessor {
	pp := engine.NewPipelineProcessor(engine.PipelineProcessorKeyRenderer)

	pp.AfterWalk = func(ctx context.Context, w io.Writer, node *engine.NodeGroup) (*engine.NodeGroup, error) {
		byteBuf := bytes.NewBuffer(nil)
		hasher := sha256.New()
		multiW := io.MultiWriter(byteBuf, hasher)

		if err := renderer.HTML(multiW, node); err != nil {
			return node, fmt.Errorf("renderer.HTML: %w", err)
		}

		doc := byteBuf.Bytes()
		hhash := fmt.Sprintf("%x", hasher.Sum(nil))
		// Add hhash to the output
		doc = bytes.Replace(doc, []byte(PageHashAttrTmpl), []byte(hhash), 1)

		if nodeBytes, err := msgpack.Marshal(node); err != nil {
			logger.Err(err).Msg("PipelineProcessorRenderHashAndCache: msgpack.Marshal")
		} else {
			cache.Set(hhash, nodeBytes)
			logger.Debug().Str("hhash", hhash).Int("size", len(nodeBytes)/1024).Msg("cache set")
		}

		if _, err := w.Write(doc); err != nil {
			return node, fmt.Errorf("write doc: %w", err)
		}

		return node, nil
	}

	return pp
}
